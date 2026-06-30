package bac

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://platform.armvm.com"
	defaultTimeout = 30 * time.Second
)

type Config struct {
	BaseURL     string
	AuthVersion string
	Timeout     time.Duration
	HTTPClient  *http.Client
	Logger      *slog.Logger
	Debug       bool
	NonceFunc   func() int64
}

type Option func(*Config)

type Client struct {
	appKey      string
	appSecret   string
	desKey      string
	baseURL     *url.URL
	authVersion string
	httpClient  *http.Client
	logger      *slog.Logger
	debug       bool
	nonceFunc   func() int64
}

type RequestOption func(*requestConfig)

type requestConfig struct {
	authVersion string
}

func NewClient(appKey, appSecret, desKey string, opts ...Option) (*Client, error) {
	if appKey == "" {
		return nil, ErrMissingAppKey
	}
	if appSecret == "" {
		return nil, ErrMissingAppSecret
	}
	if len(desKey) < 24 {
		return nil, fmt.Errorf("%w: must be at least 24 bytes", ErrInvalidDESKey)
	}

	cfg := Config{
		BaseURL:     defaultBaseURL,
		AuthVersion: defaultAuthVersion,
		Timeout:     defaultTimeout,
		NonceFunc:   defaultNonce,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.AuthVersion == "" {
		cfg.AuthVersion = defaultAuthVersion
	}
	if cfg.NonceFunc == nil {
		cfg.NonceFunc = defaultNonce
	}

	base, err := url.Parse(cfg.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("bac: parse base url: %w", err)
	}
	if base.Scheme == "" || base.Host == "" {
		return nil, fmt.Errorf("bac: base url must be absolute: %q", cfg.BaseURL)
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: cfg.Timeout}
	} else if cfg.Timeout > 0 && httpClient.Timeout == 0 {
		cp := *httpClient
		cp.Timeout = cfg.Timeout
		httpClient = &cp
	}

	return &Client{
		appKey:      appKey,
		appSecret:   appSecret,
		desKey:      desKey,
		baseURL:     base,
		authVersion: cfg.AuthVersion,
		httpClient:  httpClient,
		logger:      cfg.Logger,
		debug:       cfg.Debug,
		nonceFunc:   cfg.NonceFunc,
	}, nil
}

func WithBaseURL(baseURL string) Option {
	return func(c *Config) {
		c.BaseURL = baseURL
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Config) {
		c.HTTPClient = httpClient
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

func WithAuthVersion(authVersion string) Option {
	return func(c *Config) {
		c.AuthVersion = authVersion
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(c *Config) {
		c.Logger = logger
	}
}

func WithDebug(debug bool) Option {
	return func(c *Config) {
		c.Debug = debug
	}
}

func WithNonceFunc(fn func() int64) Option {
	return func(c *Config) {
		c.NonceFunc = fn
	}
}

func WithRequestAuthVersion(authVersion string) RequestOption {
	return func(c *requestConfig) {
		c.authVersion = authVersion
	}
}

func (c *Client) Do(ctx context.Context, path string, req any, resp any, opts ...RequestOption) error {
	rc := requestConfig{authVersion: c.authVersion}
	for _, opt := range opts {
		opt(&rc)
	}
	if rc.authVersion == "" {
		rc.authVersion = defaultAuthVersion
	}

	nonce := c.nonceFunc()
	plain, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("bac: marshal request: %w", err)
	}
	msg, err := EncryptMessage(plain, c.desKey, nonce)
	if err != nil {
		return err
	}

	body, err := json.Marshal(envelope{Msg: msg, CreateTime: nonce})
	if err != nil {
		return fmt.Errorf("bac: marshal envelope: %w", err)
	}

	endpoint := c.endpoint(path)
	endpoint.RawQuery = signedQueryForNonce(c.appKey, c.appSecret, rc.authVersion, nonce).Encode()

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("bac: create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	if c.debug && c.logger != nil {
		c.logger.DebugContext(ctx, "bac request", "path", path)
	}

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	raw, readErr := io.ReadAll(httpResp.Body)
	if readErr != nil {
		return fmt.Errorf("bac: read response: %w", readErr)
	}

	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		return &HTTPError{StatusCode: httpResp.StatusCode, Status: httpResp.Status, RawBody: raw}
	}

	var apiResp Response[json.RawMessage]
	if err := json.Unmarshal(raw, &apiResp); err != nil {
		return &DecodeError{Err: err, RawBody: raw}
	}
	if apiResp.Code != 0 {
		return &APIError{Code: apiResp.Code, Message: apiResp.Message, Timestamp: apiResp.Timestamp, RawBody: raw}
	}
	if resp == nil || len(apiResp.Data) == 0 || string(apiResp.Data) == "null" {
		return nil
	}
	if err := json.Unmarshal(apiResp.Data, resp); err != nil {
		return &DecodeError{Err: err, RawBody: apiResp.Data}
	}
	return nil
}

func (c *Client) endpoint(path string) url.URL {
	u := *c.baseURL
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		parsed, err := url.Parse(path)
		if err == nil {
			return *parsed
		}
	}
	u.Path = strings.TrimRight(u.Path, "/") + "/" + strings.TrimLeft(path, "/")
	u.RawQuery = ""
	return u
}

type envelope struct {
	Msg        string `json:"msg"`
	CreateTime int64  `json:"createTime"`
}
