package bac

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type endpointAuditCase struct {
	name       string
	wantPath   string
	dataJSON   string
	call       func(context.Context, *Client) error
	assertBody func(*testing.T, map[string]any)
}

func TestOfficialDocAuditRequestAndResponseShapes(t *testing.T) {
	tests := []endpointAuditCase{
		{
			name:     "AuthorizedConnect response fields",
			wantPath: "/auth/instance/authorized-connect",
			dataJSON: `{"serverToken":"st","instanceCode":"VM1","streamModeList":["webrtc"],"sessionId":"sid","connData":{"k":"v"}}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.AuthorizedConnect(ctx, &AuthorizedConnectRequest{UUID: "u", InstanceCode: "VM1", OnlineTime: 0, GrantControl: "CONTROL"})
				if err != nil {
					return err
				}
				if got.InstanceCode != "VM1" || got.SessionID != "sid" || len(got.StreamModeList) != 1 || len(got.ConnData) == 0 {
					t.Fatalf("authorized connect response = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "uuid", "instanceCode", "onlineTime", "grantControl")
				if _, ok := body["onlineTime"]; !ok {
					t.Fatalf("onlineTime zero was omitted: %#v", body)
				}
			},
		},
		{
			name:     "Task info uses official task result fields",
			wantPath: "/command/pad/execute-task-info.html",
			dataJSON: `[{"taskId":10,"padCode":"VM1","taskStatus":"success","taskResult":"ok","taskType":"script","createTime":11,"executeTime":12,"deviceIp":"10.0.0.1"}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetTaskResult(ctx, &GetTaskResultRequest{TaskIDs: []FlexibleString{"10"}})
				if err != nil {
					return err
				}
				if len(got) != 1 || got[0].PadCode != "VM1" || got[0].TaskStatus != "success" || got[0].TaskType != "script" {
					t.Fatalf("task result = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "taskIds")
				if _, ok := body["taskIds"].([]any)[0].(float64); !ok {
					t.Fatalf("taskIds must marshal as numbers: %#v", body["taskIds"])
				}
			},
		},
		{
			name:     "Task page request fields and pagedata alias",
			wantPath: "/command/pad/execute-task-page",
			dataJSON: `{"total":1,"page":1,"rows":10,"pagedata":[{"taskId":10,"padCode":"VM1","taskStatus":"success","taskType":"script"}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.ListTaskResults(ctx, &ListTaskResultsRequest{Page: 1, Rows: 10, TaskIDs: []FlexibleString{"10"}, TaskType: "script"})
				if err != nil {
					return err
				}
				if len(got.PageData) != 1 || got.PageData[0].TaskStatus != "success" {
					t.Fatalf("task page = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "page", "rows", "taskIds", "taskType")
				forbidKeys(t, body, "taskId", "type", "size")
			},
		},
		{
			name:     "ExportIP returns array",
			wantPath: "/resources/instance/export-ip",
			dataJSON: `[{"instanceCode":"VM1","publicIpAddress":"1.2.3.4"}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.ExportIP(ctx, &ExportIPRequest{InstanceCodes: []string{"VM1"}})
				if err != nil {
					return err
				}
				if len(got) != 1 || got[0].PublicIPAddress != "1.2.3.4" {
					t.Fatalf("export ip = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "instanceCodes") },
		},
		{
			name:     "DeviceInfo uses documented fields",
			wantPath: "/resources/instance/device-info",
			dataJSON: `[{"instanceCode":"VM1","storageSum":1024,"storageFree":512,"appInstallList":[{"appId":1,"appName":"App","packageName":"pkg","versionName":"1.0","appIconUrl":"http://icon"}]}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetInstanceDeviceInfo(ctx, &GetInstanceDeviceInfoRequest{InstanceCodes: []string{"VM1"}})
				if err != nil {
					return err
				}
				if len(got) != 1 || got[0].InstanceCode != "VM1" || got[0].StorageFree.String() != "512" || got[0].AppInstallList[0].PackageName != "pkg" {
					t.Fatalf("device info = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "instanceCodes") },
		},
		{
			name:     "SessionControlSwitch returns failure list",
			wantPath: "/resources/instance/session-control-switch",
			dataJSON: `[{"serverToken":"st","errorMsg":"bad"}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.SwitchSessionControl(ctx, &SessionControlSwitchRequest{ControlServerTokens: []string{"st"}})
				if err != nil {
					return err
				}
				if len(got) != 1 || got[0].ErrorMsg != "bad" {
					t.Fatalf("session switch = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "controlServerTokens") },
		},
		{
			name:     "DataCopy sends required false reset and returns source target",
			wantPath: "/resources/instance/data-copy",
			dataJSON: `[{"sourceInstanceCode":"VM1","targetInstanceCode":"VM2","taskId":10}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.CopyInstanceData(ctx, &DataCopyRequest{DataCopyList: []DataCopyItem{{SourceInstanceCode: "VM1", TargetInstanceCode: "VM2"}}, Reset: false})
				if err != nil {
					return err
				}
				if len(got) != 1 || got[0].SourceInstanceCode != "VM1" || got[0].TargetInstanceCode != "VM2" {
					t.Fatalf("data copy = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "dataCopyList", "reset")
				if body["reset"] != false {
					t.Fatalf("reset = %#v", body["reset"])
				}
			},
		},
		{
			name:     "MemoryLimit sends required false isLimit",
			wantPath: "/resources/instance/memory-limit",
			dataJSON: `{"successList":[],"failList":[]}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.SetMemoryLimit(ctx, &MemoryLimitRequest{InstanceCodes: []string{"VM1"}, IsLimit: false, MemoryLimit: 1024})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "instanceCodes", "isLimit", "memoryLimit")
				if body["isLimit"] != false {
					t.Fatalf("isLimit = %#v", body["isLimit"])
				}
			},
		},
		{
			name:     "NetworkBandwidthList returns object",
			wantPath: "/monitor/network-bandwidth/list",
			dataJSON: `{"peek":100,"peek95":90,"average":50,"bandwidthList":[{"recordTime":"2026-01-01","bandwidth":10}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.ListNetworkBandwidth(ctx, &NetworkBandwidthRequest{IDCCode: "idc"})
				if err != nil {
					return err
				}
				if got.Peek.String() != "100" || len(got.BandwidthList) != 1 {
					t.Fatalf("network bandwidth = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "idcCode") },
		},
		{
			name:     "InstanceAppMonitorInfo returns object",
			wantPath: "/resources/instance-app-monitor-info",
			dataJSON: `{"cpuRateTopTen":[{"packageName":"pkg","cpuRate":1}],"memRateTopTen":[{"packageName":"pkg","memRate":2}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetInstanceAppMonitorInfo(ctx, &InstanceAppMonitorInfoRequest{InstanceCode: "VM1"})
				if err != nil {
					return err
				}
				if len(got.CPURateTopTen) != 1 || len(got.MemRateTopTen) != 1 {
					t.Fatalf("app monitor = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "instanceCode") },
		},
		{
			name:     "Emulation auth code page official fields",
			wantPath: "/resources/emulation/auth-code-page",
			dataJSON: `{"total":0,"pageData":[]}`,
			call: func(ctx context.Context, c *Client) error {
				status := 0
				_, err := c.ListEmulationAuthCodes(ctx, &ListEmulationAuthCodesRequest{AuthCodes: []string{"a"}, AuthStatus: &status, InstanceCodes: []string{"VM1"}, PackageName: "pkg", Page: 1, Rows: 10})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "authCodes", "authStatus", "instanceCodes", "packageName", "page", "rows")
				forbidKeys(t, body, "authCode", "instanceCode", "status", "size")
				if body["authStatus"].(float64) != 0 {
					t.Fatalf("authStatus = %#v", body["authStatus"])
				}
			},
		},
		{
			name:     "RebootInstance returns documented array",
			wantPath: "/resources/instance/reboot-remote-play",
			dataJSON: `[{"instanceCode":"VM1","taskId":10}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.RebootInstance(ctx, &InstanceActionRequest{InstanceCodes: []string{"VM1"}})
				if err != nil {
					return err
				}
				if len(got) != 1 || got[0].InstanceCode != "VM1" {
					t.Fatalf("reboot instance = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "instanceCodes") },
		},
		{
			name:     "CommandReset returns documented array",
			wantPath: "/command/pad/reset.html",
			dataJSON: `[{"padCode":"VM1","taskId":10}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.CommandReset(ctx, &InstanceActionRequest{PadCodes: []string{"VM1"}, ResetType: "DATA_ONLY"})
				if err != nil {
					return err
				}
				if len(got) != 1 || got[0].PadCode != "VM1" {
					t.Fatalf("command reset = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "padCodes", "resetType") },
		},
		{
			name:     "InstanceMetricDetail returns documented array",
			wantPath: "/resources/instance/metric-detail",
			dataJSON: `[{"cpuUtilizationRate":1,"memoryUtilizationRate":2,"storageUtilizationRate":3}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetInstanceMetricDetail(ctx, &InstanceMetricDetailRequest{InstanceCodes: []string{"VM1"}, RecordTime: "2026-01-01 00:00:00"})
				if err != nil {
					return err
				}
				if len(got) != 1 {
					t.Fatalf("metric detail = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "instanceCodes", "recordTime") },
		},
		{
			name:     "ListBaseImages official request and response",
			wantPath: "/resources/instance-base-image/list",
			dataJSON: `{"total":1,"pageData":[{"imageVersionId":5,"imageVersionName":"base","instanceServerType":"3588","romVersion":"android12.0","createTime":1710000000000}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.ListBaseImages(ctx, &ListBaseImagesRequest{Page: 1, Rows: 10, InstanceServerType: "3588", RomVersion: "android12.0", ImageVersionIDs: []FlexibleString{"5"}})
				if err != nil {
					return err
				}
				if len(got.PageData) != 1 || got.PageData[0].ImageVersionName != "base" {
					t.Fatalf("base images = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "page", "rows", "instanceServerType", "romVersion", "imageVersionIds")
				forbidKeys(t, body, "size", "imageId", "name")
			},
		},
		{
			name:     "ListInstanceImages official request and response",
			wantPath: "/resources/instance-image/list",
			dataJSON: `{"total":1,"pageData":[{"imageVersionId":6,"imageUploadStatus":"success","instanceServerType":"3588","romVersion":"android12.0","imageFiles":[{"imageFileUrl":"u"}],"baseImageVersionId":5,"imageVersionName":"custom","describe":"d","createTime":1710000000000}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.ListInstanceImages(ctx, &ListInstanceImagesRequest{Page: 1, Rows: 10, ImageVersionID: "6", ImageVersionIDs: []FlexibleString{"6"}, InstanceServerType: "3588", RomVersion: "android12.0", ImageVersionName: "custom"})
				if err != nil {
					return err
				}
				if len(got.PageData) != 1 || got.PageData[0].ImageUploadStatus != "success" || got.PageData[0].ImageFiles[0].ImageFileURL != "u" {
					t.Fatalf("instance images = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "page", "rows", "imageVersionId", "imageVersionIds", "instanceServerType", "romVersion", "imageVersionName")
				forbidKeys(t, body, "size", "imageId", "name")
			},
		},
		{
			name:     "UpdateImage omits default autoInstall unless explicit",
			wantPath: "/resources/instance-image/update",
			dataJSON: `[{"instanceCode":"VM1","taskId":1}]`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.UpdateImage(ctx, &UpdateImageRequest{ImageVersionID: "5", InstanceCodes: []string{"VM1"}})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "imageVersionId", "instanceCodes")
				forbidKeys(t, body, "autoInstall")
			},
		},
		{
			name:     "App version returns documented ids",
			wantPath: "/resources/app/new-version",
			dataJSON: `{"applicationVersionId":100,"appId":1}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.NewAppVersion(ctx, &AppVersionRequest{AppID: "1", AppName: "app", AppPackageName: "pkg", AppURL: "http://app"})
				if err != nil {
					return err
				}
				if got.ApplicationVersionID.String() != "100" || got.AppID.String() != "1" {
					t.Fatalf("app version = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "appId", "appName", "appPackageName", "appUrl")
			},
		},
		{
			name:     "RemoveAlarmStrategy uses alarmStrategyIds",
			wantPath: "/merchant/alarm-strategy/remove",
			dataJSON: `null`,
			call: func(ctx context.Context, c *Client) error {
				return c.RemoveAlarmStrategy(ctx, &RemoveAlarmStrategyRequest{AlarmStrategyIDs: []FlexibleString{"1"}})
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "alarmStrategyIds")
				forbidKeys(t, body, "alarmStrategyId")
			},
		},
		{
			name:     "UpdateAlarmStrategyEnableStatus sends zero",
			wantPath: "/merchant/alarm-strategy/update-enable-status",
			dataJSON: `null`,
			call: func(ctx context.Context, c *Client) error {
				return c.UpdateAlarmStrategyEnableStatus(ctx, &UpdateAlarmStrategyEnableStatusRequest{AlarmStrategyIDs: []FlexibleString{"1"}, EnableStatus: 0})
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "alarmStrategyIds", "enableStatus")
				if body["enableStatus"].(float64) != 0 {
					t.Fatalf("enableStatus = %#v", body["enableStatus"])
				}
			},
		},
		{
			name:     "SaveInstanceTag returns tagId",
			wantPath: "/resources/instance-tag/save",
			dataJSON: `{"tagId":10}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.SaveInstanceTag(ctx, &InstanceTagRequest{TagName: "tag"})
				if err != nil {
					return err
				}
				if got.TagID.String() != "10" {
					t.Fatalf("save tag = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "tagName") },
		},
		{
			name:     "ListInstanceSnapshots uses documented filters",
			wantPath: "/resources/instance/snapshot-page",
			dataJSON: `{"total":0,"pageData":[]}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListInstanceSnapshots(ctx, &ListInstanceSnapshotsRequest{SnapshotIDs: []FlexibleString{"1"}, SnapshotName: "snap", SnapshotStatus: "create_success", Page: 1, Rows: 10})
				return err
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "snapshotIds", "snapshotName", "snapshotStatus", "page", "rows")
				forbidKeys(t, body, "instanceCode")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runEndpointAuditCase(t, tt)
		})
	}
}

func runEndpointAuditCase(t *testing.T, tt endpointAuditCase) {
	t.Helper()
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
		t.Fatalf("%s error = %v", tt.name, err)
	}
	tt.assertBody(t, captured)
}
