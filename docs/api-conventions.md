# API Conventions

This document explains how the SDK maps the BAC Open API into Go types.

Official API documentation: https://cloud.baidu.com/doc/ARMCM/s/2kei7tyr3

## Request Flow

All typed methods call `Client.Do` internally. `Do` handles the BAC protocol details:

- Sends `POST` requests with `Content-Type: application/json`.
- Adds query parameters `appkey`, `auth_ver`, `nonce`, and `s`.
- Signs requests with `MD5("appkey" + appKey + "auth_ver" + authVer + "nonce" + nonce + appSecret)`.
- Marshals the typed request struct to JSON.
- Encrypts the marshaled request JSON into the `msg` field using 3DES-CBC.
- Wraps the encrypted payload as:

```json
{
  "msg": "...",
  "createTime": 1710000000123
}
```

## Request Structs

Typed request structs use JSON field names from the official BAC document.

Example:

```go
results, err := client.UpdateDeviceImage(ctx, &bac.DeviceImageUpdateRequest{
	ImageVersionID:    10001,
	DeviceIPs:         []string{"10.10.240.2"},
	ConfigID:          "config-id",
	ResourcePackageID: bac.FlexibleString("20002"),
	Reset:             false,
	AutoInstall:       true,
})
```

This sends the documented top-level fields:

```json
{
  "imageVersionId": 10001,
  "deviceIps": ["10.10.240.2"],
  "configId": "config-id",
  "resourcePackageId": "20002",
  "reset": false,
  "autoInstall": true
}
```

## Response Shapes

The common BAC response envelope is:

```json
{
  "code": 0,
  "msg": "OK",
  "data": {},
  "ts": 1710000000123
}
```

`Client.Do` unwraps `data` into the response value passed by the caller. If `code != 0`, it returns `*APIError` and preserves the raw response body.

## FlexibleString

Some BAC fields appear as either JSON numbers or JSON strings across endpoints and environments, especially identifiers and timestamps. The SDK uses `FlexibleString` for these fields so decoding is stable.

```go
taskID := result.TaskID.String()
```

## Pagination

Many list endpoints return a page object with `page`, `rows`, `total`, `totalPage`, and `pageData`. The SDK represents these with:

```go
type Page[T any] struct {
	Records   []T
	List      []T
	PageData  []T
	Total     FlexibleString
	TotalPage FlexibleString
	Page      FlexibleString
	Rows      FlexibleString
}
```

Use `PageData` first for endpoints documented with `pageData`.

## RawObject

Some BAC responses are broad, loosely documented, or likely to evolve. For those endpoints, the SDK intentionally returns `RawObject`, `[]RawObject`, or `Page[RawObject]`.

This keeps the wrapper useful while avoiding inaccurate field guesses. You can decode the raw fields you need:

```go
data, err := client.GetSummaryData(ctx, &bac.SummaryDataRequest{})
if err != nil {
	return err
}

raw := data["totalCount"]
```

## Low-Level Do

Every endpoint can still be called directly:

```go
var out bac.RawObject
err := client.Do(ctx, "/resources/instance/summary-data",
	&bac.SummaryDataRequest{},
	&out,
)
```

Prefer typed methods when available. Use `Do` when the official API changes before the SDK adds a dedicated wrapper.
