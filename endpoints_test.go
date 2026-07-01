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

	_, err = c.ListInstances(context.Background(), &ListInstancesRequest{PageRequest: PageRequest{Page: 1}, Rows: 10})
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
	if captured["configId"].(string) != "cfg-1" || captured["resourcePackageId"].(float64) != 20002 {
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
		{
			name:     "UploadImage",
			wantPath: "/resources/instance-image/upload",
			dataJSON: `{"imageVersionId":10001}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UploadImage(ctx, &UploadImageRequest{
					ImageFiles: []ImageFile{{
						ImageFileURL:  "http://example.com/root_aosp.img",
						ImageFileName: "root_aosp.img",
						ImageFileType: "root_aosp",
						ImageFileMD5:  "md5",
					}},
					InstanceServerType: "3588",
					RomVersion:         "android8.1",
					BaseImageVersionID: FlexibleString("5"),
					ImageVersionName:   "custom",
					Describe:           "desc",
				})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "imageFiles", "instanceServerType", "romVersion", "baseImageVersionId", "imageVersionName", "describe")
				forbidKeys(t, body, "imageUrl", "imageMd5", "imageName")
			},
		},
		{
			name:     "UpdateImage",
			wantPath: "/resources/instance-image/update",
			dataJSON: `[{"instanceCode":"VM1","taskId":100}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateImage(ctx, &UpdateImageRequest{
					ImageVersionID:    FlexibleString("5"),
					InstanceCodes:     []string{"VM1"},
					ConfigID:          "cfg",
					ResourcePackageID: FlexibleString("7"),
					Reset:             true,
					AutoInstall:       boolPtr(true),
				})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "imageVersionId", "instanceCodes", "configId", "resourcePackageId", "reset", "autoInstall")
			},
		},
		{
			name:     "RemoveImage",
			wantPath: "/resources/instance-image/remove",
			dataJSON: `{}`,
			call: func(ctx context.Context, c *Client) error {
				return c.RemoveImage(ctx, &RemoveImageRequest{ImageVersionIDs: []FlexibleString{"10001"}})
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "imageVersionIds")
				forbidKeys(t, body, "imageVersionId")
			},
		},
		{
			name:     "SaveDesktopIconConfig",
			wantPath: "/resources/app/save-desktop-icon-config",
			dataJSON: `null`,
			call: func(ctx context.Context, c *Client) error {
				return c.SaveDesktopIconConfig(ctx, &SaveDesktopIconConfigRequest{
					InstanceCodes:       []string{"VM1"},
					AppID:               FlexibleString("1"),
					Container:           -100,
					Screen:              0,
					X:                   1,
					Y:                   2,
					OverwriteCoordinate: true,
				})
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "instanceCodes", "appId", "container", "screen", "x", "y", "overwriteCoordinate")
				forbidKeys(t, body, "config")
			},
		},
		{
			name:     "ListInstalledApps",
			wantPath: "/resources/instance/app-install-list",
			dataJSON: `[{"instanceCode":"VM1","appInstallRecordList":[{"appName":"App","packageName":"pkg","versionName":"1.0","size":1024}]}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListInstalledApps(ctx, &AppInstallListRequest{InstanceCodes: []string{"VM1"}})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "instanceCodes")
				forbidKeys(t, body, "instanceCode")
			},
		},
		{
			name:     "DownloadFile",
			wantPath: "/resources/instance/file-download",
			dataJSON: `[{"taskId":1000,"instanceCode":"VM1"}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DownloadFile(ctx, &DownloadFileRequest{InstanceFiles: []InstanceFileDownload{{InstanceCode: "VM1", FilePath: "/sdcard/a.txt"}}})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "instanceFiles")
			},
		},
		{
			name:     "NewPad",
			wantPath: "/distribute/pad/new-pad.html",
			dataJSON: `[{"padCode":"VM1","taskId":100}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.NewPad(ctx, &NewPadRequest{PadModels: []PadModel{{PadCode: "VM1", IMEI: "imei", SerialNo: "sn"}}})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "padModels")
				forbidKeys(t, body, "padCodes")
			},
		},
		{
			name:     "CleanAppCache",
			wantPath: "/resources/device/clean-app-cache",
			dataJSON: `[{"taskId":"2813218757970419736","deviceIp":"11.101.1.49"}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.CleanAppCache(ctx, &CleanAppCacheRequest{DeviceIPs: []string{"11.101.1.49"}})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "DeviceIps")
				forbidKeys(t, body, "deviceIps", "instanceCodes")
			},
		},
		{
			name:     "UpdateCustomCode",
			wantPath: "/resources/instance/custom-code-update",
			dataJSON: `[{"instanceCode":"VM1","errorMsg":""}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateCustomCode(ctx, &CustomCodeUpdateRequest{InstanceCustomList: []InstanceCustom{{InstanceCode: "VM1", CustomCode: "DB-01"}}})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "instanceCustomList")
				forbidKeys(t, body, "instanceCodes", "customCode")
			},
		},
		{
			name:     "DeployMarketingSuite",
			wantPath: "/resources/instance/deploy-marketing-suite",
			dataJSON: `[{"instanceCode":"VM1","authCode":"auth","taskId":1001}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.DeployMarketingSuite(ctx, &DeployMarketingSuiteRequest{
					InstanceCodes:       []string{"VM1"},
					AuthCodes:           []string{"auth"},
					AppPackage:          "com.demo",
					UseMerchantAuthCode: true,
				})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "instanceCodes", "authCodes", "appPackage", "useMerchantAuthCode")
			},
		},
		{
			name:     "SaveInstancePool",
			wantPath: "/resources/instance-pool/save",
			dataJSON: `{"merchantPoolNo":"10001"}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SaveInstancePool(ctx, &SaveInstancePoolRequest{
					ParentMerchantPoolNo: FlexibleString("10000"),
					InstancePoolName:     "pool",
					InstancePoolType:     "cloud_phone",
				})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "parentMerchantPoolNo", "instancePoolName", "instancePoolType")
				forbidKeys(t, body, "poolName")
			},
		},
		{
			name:     "OpenAccount",
			wantPath: "/merchant/open-account/save",
			dataJSON: `{"password":"pwd"}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.OpenAccount(ctx, &OpenAccountRequest{
					UserName: "admin", Phone: "13800000000", Nickname: "Admin",
					RoleNames: []string{"role"}, MerchantPoolNos: []FlexibleString{"10000"},
				})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "userName", "phone", "nickname", "roleNames", "merchantPoolNos")
				forbidKeys(t, body, "merchantName", "contactName")
			},
		},
		{
			name:     "AddSubMerchant",
			wantPath: "/merchant/sub-merchant/add",
			dataJSON: `{"adminPassword":"pwd"}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.AddSubMerchant(ctx, &AddSubMerchantRequest{
					ParentMerchantCode: "10000", MerchantCode: "10001", MerchantName: "merchant",
					MerchantType: "cloud_phone", MerchantPhone: "13800000000", AdminUserName: "admin", AdminPhone: "13900000000",
				})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "parentMerchantCode", "merchantCode", "merchantName", "merchantType", "merchantPhone", "adminUserName", "adminPhone")
				forbidKeys(t, body, "merchantNo", "contactName")
			},
		},
		{
			name:     "InitDUFSSnapshot",
			wantPath: "/resources/dufs-snapshot/init",
			dataJSON: `[{"instanceCode":"VM1","taskId":100}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.InitDUFSSnapshot(ctx, &InitDUFSSnapshotRequest{InstanceCodes: []string{"VM1"}, QuotaCapacity: 256, MemoryLimit: 1024})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "instanceCodes", "quotaCapacity", "memoryLimit")
				forbidKeys(t, body, "quota")
			},
		},
		{
			name:     "BatchMountDUFSSnapshot",
			wantPath: "/resources/dufs-snapshot/batch-mount",
			dataJSON: `[{"instanceCode":"VM1","taskId":100}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.BatchMountDUFSSnapshot(ctx, &BatchMountDUFSSnapshotRequest{
					SnapshotMountInfos: []SnapshotMountInfo{{InstanceCode: "VM1", SnapshotID: FlexibleString("10"), QuotaCapacity: 256}},
				})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "snapshotMountInfos")
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

func boolPtr(v bool) *bool {
	return &v
}
