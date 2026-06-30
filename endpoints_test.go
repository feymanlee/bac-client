package bac

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTypedWrapperUsesDocumentPath(t *testing.T) {
	var gotPath string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":0,"msg":"OK","data":{"records":[],"total":"0"},"ts":1710000000124}`))
	}))
	defer ts.Close()

	c, err := NewClient("ak", "sk", "123456789012345678901234", WithBaseURL(ts.URL), WithNonceFunc(func() int64 {
		return 1710000000123
	}))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	_, err = c.ListInstances(context.Background(), &ListInstancesRequest{PageRequest: PageRequest{Page: 1, Size: 10}})
	if err != nil {
		t.Fatalf("ListInstances() error = %v", err)
	}
	if gotPath != "/resources/instance/infos" {
		t.Fatalf("path = %q", gotPath)
	}
}

func TestAdditionalTypedWrappersUseDocumentPaths(t *testing.T) {
	tests := []struct {
		name     string
		wantPath string
		dataJSON string
		call     func(context.Context, *Client) error
	}{
		{
			name:     "UploadApps",
			wantPath: "/distribute/apps/uploads.html",
			dataJSON: `{}`,
			call: func(ctx context.Context, c *Client) error {
				return c.UploadApps(ctx, &UploadAppsRequest{})
			},
		},
		{
			name:     "ListApps",
			wantPath: "/resources/app/page",
			dataJSON: `{"records":[],"pageData":[],"total":"0"}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListApps(ctx, &ListAppsRequest{})
				return err
			},
		},
		{
			name:     "GetDeviceMonitorInfo",
			wantPath: "/resources/device-monitor-info",
			dataJSON: `[]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetDeviceMonitorInfo(ctx, &DeviceMonitorInfoRequest{})
				return err
			},
		},
		{
			name:     "AddAlarmStrategy",
			wantPath: "/merchant/alarm-strategy/add",
			dataJSON: `{}`,
			call: func(ctx context.Context, c *Client) error {
				return c.AddAlarmStrategy(ctx, &AlarmStrategyRequest{})
			},
		},
		{
			name:     "DownloadFile",
			wantPath: "/resources/instance/file-download",
			dataJSON: `{"successList":[],"failList":[]}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DownloadFile(ctx, &DownloadFileRequest{})
				return err
			},
		},
		{
			name:     "ListDUFSSnapshots",
			wantPath: "/resources/dufs-snapshot/page",
			dataJSON: `{"records":[],"pageData":[],"total":"0"}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListDUFSSnapshots(ctx, &ListDUFSSnapshotsRequest{})
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotPath string
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotPath = r.URL.Path
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"code":0,"msg":"OK","data":` + tt.dataJSON + `,"ts":1710000000124}`))
			}))
			defer ts.Close()

			c, err := NewClient("ak", "sk", "123456789012345678901234", WithBaseURL(ts.URL), WithNonceFunc(func() int64 {
				return 1710000000123
			}))
			if err != nil {
				t.Fatalf("NewClient() error = %v", err)
			}

			if err := tt.call(context.Background(), c); err != nil {
				t.Fatalf("%s() error = %v", tt.name, err)
			}
			if gotPath != tt.wantPath {
				t.Fatalf("path = %q, want %q", gotPath, tt.wantPath)
			}
		})
	}
}

func TestFlexibleStringAcceptsNumbersAndStrings(t *testing.T) {
	var got struct {
		A FlexibleString `json:"a"`
		B FlexibleString `json:"b"`
	}
	if err := json.Unmarshal([]byte(`{"a":123,"b":"456"}`), &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if got.A.String() != "123" || got.B.String() != "456" {
		t.Fatalf("got A=%q B=%q", got.A, got.B)
	}
}

func TestUpdateDeviceImageUsesOfficialRequestAndResponseShape(t *testing.T) {
	var captured map[string]any
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/resources/device/image-update" {
			t.Fatalf("path = %q", r.URL.Path)
		}

		var env envelope
		if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
			t.Fatalf("decode envelope: %v", err)
		}
		plain, err := decryptForTest(env.Msg, "123456789012345678901234", env.CreateTime)
		if err != nil {
			t.Fatalf("decrypt request: %v", err)
		}
		if err := json.Unmarshal(plain, &captured); err != nil {
			t.Fatalf("decode plain request: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"code":0,"msg":"","data":[{"deviceIp":"10.10.240.2","taskId":9001}],"ts":1601021167163}`))
	}))
	defer ts.Close()

	c, err := NewClient("ak", "sk", "123456789012345678901234", WithBaseURL(ts.URL), WithNonceFunc(func() int64 {
		return 1710000000123
	}))
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	got, err := c.UpdateDeviceImage(context.Background(), &DeviceImageUpdateRequest{
		ImageVersionID:    10001,
		DeviceIPs:         []string{"10.10.240.2"},
		ConfigID:          "cfg-1",
		ResourcePackageID: FlexibleString("20002"),
		Reset:             true,
		AutoInstall:       false,
	})
	if err != nil {
		t.Fatalf("UpdateDeviceImage() error = %v", err)
	}

	if captured["imageVersionId"].(float64) != 10001 {
		t.Fatalf("imageVersionId = %#v", captured["imageVersionId"])
	}
	if captured["deviceIps"].([]any)[0].(string) != "10.10.240.2" {
		t.Fatalf("deviceIps = %#v", captured["deviceIps"])
	}
	if _, ok := captured["deviceCodes"]; ok {
		t.Fatalf("unexpected deviceCodes in request: %#v", captured)
	}
	if _, ok := captured["imageId"]; ok {
		t.Fatalf("unexpected imageId in request: %#v", captured)
	}
	if captured["configId"].(string) != "cfg-1" || captured["resourcePackageId"].(string) != "20002" {
		t.Fatalf("optional fields = %#v", captured)
	}
	if captured["reset"].(bool) != true || captured["autoInstall"].(bool) != false {
		t.Fatalf("bool fields = %#v", captured)
	}

	if len(got) != 1 || got[0].DeviceIP != "10.10.240.2" || got[0].TaskID.String() != "9001" {
		t.Fatalf("response = %#v", got)
	}
}

func TestOfficialRequestShapesForHighRiskEndpoints(t *testing.T) {
	tests := []struct {
		name       string
		wantPath   string
		dataJSON   string
		call       func(context.Context, *Client) error
		assertBody func(*testing.T, map[string]any)
	}{
		{
			name:     "GetServerToken",
			wantPath: "/auth/instance/cloud-phone-server-token",
			dataJSON: `{"successList":[],"errorList":[]}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.GetServerToken(ctx, &GetServerTokenRequest{
					UUID: "u1", InstanceCodes: []string{"VM1"}, OnlineTime: 300, GrantControl: "WATCH",
				})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "uuid", "instanceCodes", "onlineTime", "grantControl")
				forbidKeys(t, body, "instanceId", "userId")
			},
		},
		{
			name:     "BatchDisconnectInstances",
			wantPath: "/auth/instance/batch-disconnect",
			dataJSON: `{"successList":["st1"],"failList":[]}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BatchDisconnectInstances(ctx, &BatchDisconnectInstancesRequest{
					DisconnectList: []DisconnectCredential{{UUID: "u1", ServerToken: "st1"}},
				})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "disconnectList")
				forbidKeys(t, body, "instanceIds", "userId")
			},
		},
		{
			name:     "BackupInstance",
			wantPath: "/resources/instance/backup",
			dataJSON: `{"taskId":100,"instanceCode":"VM1","snapshotId":200,"snapshotName":"snap"}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BackupInstance(ctx, &BackupInstanceRequest{
					InstanceCode: "VM1",
					SnapshotName: "snap",
					OSSConfig:    OSSConfig{Endpoint: "s3.example.com", Bucket: "b", AccessKey: "ak", SecretKey: "sk", Protocol: "https"},
				})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "instanceCode", "snapshotName", "ossConfig")
				forbidKeys(t, body, "instanceId", "instanceIds")
			},
		},
		{
			name:     "Screenshot",
			wantPath: "/command/pad/screenshot.html",
			dataJSON: `[{"taskId":1,"padCode":"VM1"}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.Screenshot(ctx, &ScreenshotRequest{PadCodes: []string{"VM1"}, Quality: 80, PictureType: "png"})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "padCodes", "quality", "pictureType")
				forbidKeys(t, body, "instanceId")
			},
		},
		{
			name:     "OperateApp",
			wantPath: "/command/apps/app-operate.html",
			dataJSON: `[{"taskId":1,"padCode":"VM1"}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.OperateApp(ctx, &AppOperateRequest{PadCodes: []string{"VM1"}, PackageName: "pkg", OperateType: "start"})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "padCodes", "packageName", "operateType")
				forbidKeys(t, body, "operation", "instanceIds")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var captured map[string]any
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != tt.wantPath {
					t.Fatalf("path = %q, want %q", r.URL.Path, tt.wantPath)
				}
				var env envelope
				if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
					t.Fatalf("decode envelope: %v", err)
				}
				plain, err := decryptForTest(env.Msg, "123456789012345678901234", env.CreateTime)
				if err != nil {
					t.Fatalf("decrypt request: %v", err)
				}
				if err := json.Unmarshal(plain, &captured); err != nil {
					t.Fatalf("decode plain request: %v", err)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"code":0,"msg":"OK","data":` + tt.dataJSON + `,"ts":1710000000124}`))
			}))
			defer ts.Close()

			c, err := NewClient("ak", "sk", "123456789012345678901234", WithBaseURL(ts.URL), WithNonceFunc(func() int64 {
				return 1710000000123
			}))
			if err != nil {
				t.Fatalf("NewClient() error = %v", err)
			}
			if err := tt.call(context.Background(), c); err != nil {
				t.Fatalf("%s() error = %v", tt.name, err)
			}
			tt.assertBody(t, captured)
		})
	}
}

func requireKeys(t *testing.T, body map[string]any, keys ...string) {
	t.Helper()
	for _, key := range keys {
		if _, ok := body[key]; !ok {
			t.Fatalf("missing key %q in body %#v", key, body)
		}
	}
}

func forbidKeys(t *testing.T, body map[string]any, keys ...string) {
	t.Helper()
	for _, key := range keys {
		if _, ok := body[key]; ok {
			t.Fatalf("unexpected key %q in body %#v", key, body)
		}
	}
}
