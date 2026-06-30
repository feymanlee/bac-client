# Baidu BAC Go SDK

Idiomatic Go SDK for 百度智能云「云手机 BAC」Open API.

Official API documentation: https://cloud.baidu.com/doc/ARMCM/s/2kei7tyr3

Default API base URL: `https://platform.armvm.com`.

## Features

- Standard-library-only runtime implementation.
- Automatic BAC query signing and encrypted `msg` request envelope.
- Context-aware HTTP calls with configurable base URL, timeout, HTTP client, auth version, and logger.
- Typed wrappers for the official Open API paths plus low-level `Do` for forward compatibility.
- Structured `APIError`, `HTTPError`, and `DecodeError` values.
- Unit tests and `httptest` coverage with no dependency on real Baidu credentials.

## Install

```bash
go get github.com/feymanlee/bac-client
```

## Quick Start

```go
package main

import (
	"context"
	"log"
	"os"
	"time"

	bac "github.com/feymanlee/bac-client"
)

func main() {
	client, err := bac.NewClient(
		os.Getenv("BAC_APP_KEY"),
		os.Getenv("BAC_APP_SECRET"),
		os.Getenv("BAC_DES_KEY"),
		bac.WithTimeout(15*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	page, err := client.ListInstances(ctx, &bac.ListInstancesRequest{
		PageRequest: bac.PageRequest{Page: 1},
		Rows:        20,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("instances: %+v", page)
}
```

Never commit real `appKey`, `appSecret`, or `desKey`. Use environment variables or a secret manager.

## Base URLs

The default endpoint is:

```text
https://platform.armvm.com
```

Use `WithBaseURL` for other BAC environments:

```go
client, err := bac.NewClient(appKey, appSecret, desKey,
	bac.WithBaseURL("https://oversea-platform.armvm.com"),
)
```

## Client Options

- `WithBaseURL(string)` sets a custom endpoint, useful for tests or proxying.
- `WithHTTPClient(*http.Client)` injects a caller-owned HTTP client.
- `WithTimeout(time.Duration)` sets the default client timeout when no timeout is already configured.
- `WithAuthVersion(string)` overrides the default `auth_ver`, which is `3`.
- `WithLogger(*slog.Logger)` and `WithDebug(bool)` enable opt-in structured debug logging. The SDK never prints by itself.

Every public API method accepts `context.Context`.

## Low-Level Do

Use `Do` for endpoints that do not yet have a typed wrapper:

```go
var out struct {
	TaskID bac.FlexibleString `json:"taskId"`
}

err := client.Do(ctx, "/resources/instance/backup",
	map[string]string{"instanceId": "your-instance-id"},
	&out,
)
```

`Do` automatically:

- Adds query parameters `appkey`, `auth_ver`, `nonce`, and `s`.
- Signs with `MD5("appkey" + appKey + "auth_ver" + authVer + "nonce" + nonce + appSecret)`.
- Marshals the real request to JSON, encrypts it with 3DES-CBC using `desKey[:24]`, `nonce` as a big-endian 8-byte IV, and PKCS5/PKCS7 padding.
- Sends `POST` with `Content-Type: application/json`.
- Decodes `{ "code": 0, "msg": "OK", "data": ..., "ts": ... }`.

## Errors

```go
var apiErr *bac.APIError
var httpErr *bac.HTTPError
var decodeErr *bac.DecodeError

switch {
case errors.As(err, &apiErr):
	log.Printf("BAC error: code=%d message=%s ts=%s", apiErr.Code, apiErr.Message, apiErr.Timestamp)
case errors.As(err, &httpErr):
	log.Printf("HTTP error: status=%s", httpErr.Status)
case errors.As(err, &decodeErr):
	log.Printf("decode error: %v", decodeErr.Err)
}
```

`APIError.RawBody`, `HTTPError.RawBody`, and `DecodeError.RawBody` keep the full response body for diagnostics. Treat these bytes as sensitive: responses may contain server tokens, SSH credentials, generated passwords, or business data. Redact or truncate before writing them to logs.

## Implemented Typed APIs

- Auth/session: `GetServerToken`, `GetAuthorizedServerToken`, `AuthorizedConnect`, `DisconnectInstance`, `BatchDisconnectInstances`, `DisconnectAllInstances`.
- Instance query and actions: `ListInstances`, `GetInstanceDeviceInfo`, `UploadFile`, `DownloadFile`, `BackupInstance`, `RestoreInstance`, `SwitchSessionControl`, `SetResolution`, `RebootInstance`, `RebootDevice`, `ResetDevice`, `NewPad`, `CleanAppCache`, `BatchExecuteScript`, `GetBindInfos`, `GetSSHInfo`, `GetScreenshotURL`, `DestroyScreenshotURL`, `StopEventTasks`, `SetSpeed`, `SetMemoryLimit`, `CopyInstanceData`.
- Apps and commands: `InstallApp`, `UninstallApp`, `OperateApp`, `StartApp`, `StopApp`, `Screenshot`, `GetScreenshotResult`, `ExecuteScript`, `UploadApps`, `ListApps`, `DeleteApp`, `NewAppVersion`, `UpgradeApp`, `ListInstalledApps`.
- Tasks: `GetTaskResult`, `ListTaskResults`, `ListTaskTypes`.
- Images: `ListBaseImages`, `ListInstanceImages`, `UploadImage`, `GetImageUploadInfo`, `UpdateImage`, `GetImageUpdateInfo`, `RemoveImage`.
- Monitoring, merchant, pool, tag, and snapshot APIs: see `docs/endpoints.md` for the full coverage table.

Additional documentation:

- [API conventions](docs/api-conventions.md)
- [Endpoint coverage](docs/endpoints.md)

## Extending New Endpoints

1. Add request and response structs with JSON tags matching the official document.
2. Add the path constant in the matching `endpoints_*.go` file.
3. Add a one-layer method wrapper:

```go
func (c *Client) SomeAction(ctx context.Context, req *SomeActionRequest) (*SomeActionResponse, error) {
	var resp SomeActionResponse
	if err := c.Do(ctx, "/official/path", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
```

Use `FlexibleString` for fields that the document or service may return as either JSON number or string.

Most official list APIs use `page` and `rows`. Prefer request-specific `Rows` fields, for example `ListInstancesRequest{PageRequest: bac.PageRequest{Page: 1}, Rows: 20}`.

## Tests

Tests use `httptest` and do not call the real Baidu service:

```bash
go test ./...
go vet ./...
```

The GitHub Actions workflow in `.github/workflows/ci.yml` runs formatting checks, `go test ./...`, and `go vet ./...` on pushes and pull requests.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE).
