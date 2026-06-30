package bac

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestDoPostsEncryptedEnvelopeWithSignedQuery(t *testing.T) {
	var captured struct {
		Query       map[string]string
		ContentType string
		Envelope    envelope
		Plain       map[string]any
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/resources/instance/infos" {
			t.Fatalf("path = %q", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Fatalf("method = %q", r.Method)
		}
		captured.ContentType = r.Header.Get("Content-Type")
		captured.Query = map[string]string{
			"appkey":   r.URL.Query().Get("appkey"),
			"auth_ver": r.URL.Query().Get("auth_ver"),
			"nonce":    r.URL.Query().Get("nonce"),
			"s":        r.URL.Query().Get("s"),
		}
		if err := json.NewDecoder(r.Body).Decode(&captured.Envelope); err != nil {
			t.Fatalf("Decode body: %v", err)
		}
		plain, err := decryptForTest(captured.Envelope.Msg, "123456789012345678901234", captured.Envelope.CreateTime)
		if err != nil {
			t.Fatalf("decrypt msg: %v", err)
		}
		if err := json.Unmarshal(plain, &captured.Plain); err != nil {
			t.Fatalf("unmarshal plain: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":0,"msg":"OK","data":{"total":1},"ts":1710000000124}`))
	}))
	defer ts.Close()

	c, err := NewClient("ak", "sk", "123456789012345678901234", WithBaseURL(ts.URL), WithNonceFunc(func() int64 {
		return 1710000000123
	}))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	var out struct {
		Total int `json:"total"`
	}
	err = c.Do(context.Background(), "/resources/instance/infos", map[string]any{"page": 1}, &out)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if captured.ContentType != "application/json" {
		t.Fatalf("Content-Type = %q", captured.ContentType)
	}
	if captured.Query["appkey"] != "ak" || captured.Query["auth_ver"] != "3" || captured.Query["nonce"] != "1710000000123" {
		t.Fatalf("query = %#v", captured.Query)
	}
	if captured.Query["s"] != Sign("ak", "sk", "3", "1710000000123") {
		t.Fatalf("signature = %q", captured.Query["s"])
	}
	if captured.Envelope.CreateTime != 1710000000123 {
		t.Fatalf("createTime = %d", captured.Envelope.CreateTime)
	}
	if captured.Plain["page"].(float64) != 1 {
		t.Fatalf("plain request = %#v", captured.Plain)
	}
	if out.Total != 1 {
		t.Fatalf("response total = %d", out.Total)
	}
}

func TestDoReturnsAPIErrorForNonZeroCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":4001,"msg":"bad auth","data":null,"ts":"1710000000124"}`))
	}))
	defer ts.Close()

	c, err := NewClient("ak", "sk", "123456789012345678901234", WithBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	err = c.Do(context.Background(), "/x", map[string]string{"a": "b"}, nil)
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("Do() error = %T, want *APIError", err)
	}
	if apiErr.Code != 4001 || apiErr.Message != "bad auth" || apiErr.Timestamp.String() != "1710000000124" {
		t.Fatalf("APIError = %#v", apiErr)
	}
	if !bytes.Contains(apiErr.RawBody, []byte(`"code":4001`)) {
		t.Fatalf("RawBody = %s", apiErr.RawBody)
	}
}

func TestDoReturnsHTTPErrorForNon2xx(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusBadGateway)
	}))
	defer ts.Close()

	c, err := NewClient("ak", "sk", "123456789012345678901234", WithBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	err = c.Do(context.Background(), "/x", nil, nil)
	var httpErr *HTTPError
	if !errors.As(err, &httpErr) {
		t.Fatalf("Do() error = %T, want *HTTPError", err)
	}
	if httpErr.StatusCode != http.StatusBadGateway {
		t.Fatalf("StatusCode = %d", httpErr.StatusCode)
	}
}

func TestDoReturnsDecodeErrorForInvalidJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`not json`))
	}))
	defer ts.Close()

	c, err := NewClient("ak", "sk", "123456789012345678901234", WithBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	err = c.Do(context.Background(), "/x", nil, nil)
	var decodeErr *DecodeError
	if !errors.As(err, &decodeErr) {
		t.Fatalf("Do() error = %T, want *DecodeError", err)
	}
}

func TestDoHonorsContextTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		_, _ = w.Write([]byte(`{"code":0,"msg":"OK"}`))
	}))
	defer ts.Close()

	c, err := NewClient("ak", "sk", "123456789012345678901234", WithBaseURL(ts.URL))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err = c.Do(ctx, "/x", nil, nil)
	if !errors.Is(err, context.DeadlineExceeded) && !strings.Contains(err.Error(), "context deadline exceeded") {
		t.Fatalf("Do() error = %v, want context deadline exceeded", err)
	}
}

func decryptForTest(ciphertext string, desKey string, nonce int64) ([]byte, error) {
	raw, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	return decrypt3DESCBCForTest(raw, []byte(desKey[:24]), nonceIV(nonce))
}
