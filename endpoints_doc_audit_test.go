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
			name:     "ListInstances decodes official pageData fields",
			wantPath: "/resources/instance/infos",
			dataJSON: `{"total":1,"totalPage":1,"page":1,"rows":10,"pageData":[{"instanceCode":"VM192168232001","instanceIp":"192.168.232.1","deviceIp":"192.168.1.100","deviceIpSegment":"192.168.1","deviceCode":"GZNX-IDC_DM011010099026","instanceType":"business","instanceServerType":"3588","instanceGrade":"1","romVersion":"android12.0","merchantPoolNo":1,"idcCode":"GZNX-IDC-02","idcName":"南翔机房-可用区1","availableStatus":"available","usableStatus":"usable","maintainStatus":0,"recycleStatus":"normal","taskStatus":"none","instanceStatus":1,"deviceStatus":1,"malfunctionStatus":0,"networkStatus":0,"bindStatus":"bindable","controlStatus":"online","imageVersionId":10001,"imageVersionName":"3588_aosp10_v2.25.1(2025-03-13)","snapshotMountStatus":"mounted","snapshotId":5001,"memoryLimit":1024,"upSpeed":100.0,"downSpeed":100.0,"intranetUpSpeed":10.0,"intranetDownSpeed":10.0,"allocateTime":1733297481137,"recycleTime":1733297481137,"tags":[{"tagId":"10001","tagName":"tag_name"}],"emulationAuthStatus":0}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.ListInstances(ctx, &ListInstancesRequest{PageRequest: PageRequest{Page: 1}, Rows: 10})
				if err != nil {
					return err
				}
				if len(got.PageData) != 1 {
					t.Fatalf("pageData length = %d", len(got.PageData))
				}
				inst := got.PageData[0]
				if inst.InstanceCode != "VM192168232001" ||
					inst.InstanceIP != "192.168.232.1" ||
					inst.DeviceIP != "192.168.1.100" ||
					inst.DeviceIPSegment != "192.168.1" ||
					inst.DeviceCode != "GZNX-IDC_DM011010099026" ||
					inst.InstanceType != "business" ||
					inst.InstanceServerType != "3588" ||
					inst.InstanceGrade != "1" ||
					inst.RomVersion != "android12.0" ||
					inst.MerchantPoolNo.String() != "1" ||
					inst.IDCCode != "GZNX-IDC-02" ||
					inst.IDCName != "南翔机房-可用区1" ||
					inst.AvailableStatus != "available" ||
					inst.UsableStatus != "usable" ||
					inst.MaintainStatus != 0 ||
					inst.RecycleStatus != "normal" ||
					inst.TaskStatus != "none" ||
					inst.InstanceStatus != 1 ||
					inst.DeviceStatus != 1 ||
					inst.MalfunctionStatus != 0 ||
					inst.NetworkStatus != 0 ||
					inst.BindStatus != "bindable" ||
					inst.ControlStatus != "online" ||
					inst.ImageVersionID.String() != "10001" ||
					inst.ImageVersionName != "3588_aosp10_v2.25.1(2025-03-13)" ||
					inst.SnapshotMountStatus != "mounted" ||
					inst.SnapshotID.String() != "5001" ||
					inst.MemoryLimit.String() != "1024" ||
					inst.UpSpeed.String() != "100.0" ||
					inst.DownSpeed.String() != "100.0" ||
					inst.IntranetUpSpeed.String() != "10.0" ||
					inst.IntranetDownSpeed.String() != "10.0" ||
					inst.AllocateTime.String() != "1733297481137" ||
					inst.RecycleTime.String() != "1733297481137" ||
					len(inst.Tags) != 1 ||
					inst.Tags[0].TagID.String() != "10001" ||
					inst.Tags[0].TagName != "tag_name" ||
					inst.EmulationAuthStatus != 0 {
					t.Fatalf("instance = %#v", inst)
				}
				if _, ok := inst.Raw["instanceIp"]; !ok {
					t.Fatalf("raw missing instanceIp: %#v", inst.Raw)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "page", "rows")
				forbidKeys(t, body, "size")
			},
		},
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
			dataJSON: `[{"taskId":10,"padCode":"VM1","taskStatus":1,"taskResult":"ok","taskType":"script","createTime":11,"executeTime":12,"deviceIp":"10.0.0.1"}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetTaskResult(ctx, &GetTaskResultRequest{TaskIDs: []FlexibleString{"10"}})
				if err != nil {
					return err
				}
				if len(got) != 1 || got[0].PadCode != "VM1" || got[0].TaskStatus.String() != "1" || got[0].TaskType != "script" {
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
			dataJSON: `{"total":1,"page":1,"rows":10,"pagedata":[{"taskId":10,"padCode":"VM1","taskStatus":1,"taskType":"script"}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.ListTaskResults(ctx, &ListTaskResultsRequest{Page: 1, Rows: 10, TaskIDs: []FlexibleString{"10"}, TaskType: "script"})
				if err != nil {
					return err
				}
				if len(got.PageData) != 1 || got.PageData[0].TaskStatus.String() != "1" {
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
			dataJSON: `{"peek":100,"peek95":90,"average":50,"bandwidthList":[{"recordTime":"2026-01-01","send":10,"receive":20}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.ListNetworkBandwidth(ctx, &NetworkBandwidthRequest{IDCCode: "idc"})
				if err != nil {
					return err
				}
				if got.Peek.String() != "100" || len(got.BandwidthList) != 1 || got.BandwidthList[0].Send.String() != "10" || got.BandwidthList[0].Receive.String() != "20" {
					t.Fatalf("network bandwidth = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "idcCode") },
		},
		{
			name:     "InstanceAppMonitorInfo returns object",
			wantPath: "/resources/instance-app-monitor-info",
			dataJSON: `{"cpuRateTopTen":[{"appName":"app","metricInfoList":[{"recordTime":1733297481137,"value":1}]}],"memRateTopTen":[{"appName":"app","metricInfoList":[{"recordTime":1733297481138,"value":2}]}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetInstanceAppMonitorInfo(ctx, &InstanceAppMonitorInfoRequest{InstanceCode: "VM1"})
				if err != nil {
					return err
				}
				if len(got.CPURateTopTen) != 1 || got.CPURateTopTen[0].AppName != "app" || got.MemRateTopTen[0].MetricInfoList[0].Value.String() != "2" {
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
		{
			name:     "SummaryData returns typed totals",
			wantPath: "/resources/instance/summary-data",
			dataJSON: `{"totalData":{"totalInstanceCount":10,"idcList":[{"idcName":"idc","idcCode":"IDC","instanceCount":5}]},"availableData":{"totalInstanceCount":6,"idcList":[]},"boundData":{"totalInstanceCount":4,"idcList":[]}}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetSummaryData(ctx, &SummaryDataRequest{MerchantPoolNo: "1", IncludeSubPool: true})
				if err != nil {
					return err
				}
				if got.TotalData.TotalInstanceCount.String() != "10" || got.TotalData.IDCList[0].IDCCode != "IDC" || got.AvailableData.TotalInstanceCount.String() != "6" {
					t.Fatalf("summary data = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "merchantPoolNo", "includeSubPool") },
		},
		{
			name:     "ListDevices returns documented device page fields",
			wantPath: "/resources/device/page",
			dataJSON: `{"total":1,"pageData":[{"deviceIp":"192.168.1.100","deviceIpSegment":"192.168.1","deviceCode":"D1","idcCode":"IDC","idcName":"机房","instanceServerType":"3588","romVersion":"android12.0","instanceGrade":1,"instanceGradeLimit":5,"allocateTime":1733297481137,"recycleTime":1733297481138,"memoryModel":"16GB","cpuModel":"8C","diskModel":"256GB","deviceStatus":1}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.ListDevices(ctx, &ListDevicesRequest{PageRequest: PageRequest{Page: 1}, Rows: 10})
				if err != nil {
					return err
				}
				if len(got.PageData) != 1 || got.PageData[0].DeviceCode != "D1" || got.PageData[0].InstanceGrade.String() != "1" || got.PageData[0].CPUModel != "8C" {
					t.Fatalf("devices = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "page", "rows") },
		},
		{
			name:     "Emulation auth code page returns typed auth codes",
			wantPath: "/resources/emulation/auth-code-page",
			dataJSON: `{"total":1,"pageData":[{"authCode":"A","expireTime":1733297481137,"authStatus":1,"activateTime":1733297481138,"packageName":"pkg","deviceIp":"1.1.1.1","instanceIp":"2.2.2.2","instanceCode":"VM1","createTime":1733297481139}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.ListEmulationAuthCodes(ctx, &ListEmulationAuthCodesRequest{Page: 1, Rows: 10})
				if err != nil {
					return err
				}
				if len(got.PageData) != 1 || got.PageData[0].AuthCode != "A" || got.PageData[0].AuthStatus != 1 || got.PageData[0].ExpireTime.String() != "1733297481137" {
					t.Fatalf("auth code page = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "page", "rows") },
		},
		{
			name:     "Pad available count returns typed idc infos",
			wantPath: "/monitor/pad/available-count.html",
			dataJSON: `{"totalAvailableCount":3,"idcInfos":[{"idcName":"idc","idcCode":"IDC","availableCount":3}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetAvailablePadCount(ctx, &PadCountRequest{IDCCodes: []string{"IDC"}})
				if err != nil {
					return err
				}
				if got.TotalAvailableCount.String() != "3" || got.IDCInfos[0].AvailableCount.String() != "3" {
					t.Fatalf("available count = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "idcCodes") },
		},
		{
			name:     "Pad enable bind count returns typed idc infos",
			wantPath: "/monitor/pad/enable-bind-count.html",
			dataJSON: `{"totalEnableBindCount":2,"idcInfos":[{"idcName":"idc","idcCode":"IDC","enableBindCount":2}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetEnableBindPadCount(ctx, &PadCountRequest{IDCCodes: []string{"IDC"}})
				if err != nil {
					return err
				}
				if got.TotalEnableBindCount.String() != "2" || got.IDCInfos[0].EnableBindCount.String() != "2" {
					t.Fatalf("enable bind count = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "idcCodes") },
		},
		{
			name:     "QueryFlow returns documented task id object",
			wantPath: "/distribute/task/query-flow",
			dataJSON: `{"taskId":89695}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.QueryFlow(ctx, &QueryFlowRequest{InstanceCodes: []string{"VM1"}, StartTime: "2026-01-01 00:00:00", EndTime: "2026-01-01 01:00:00", BillingType: "FLOW"})
				if err != nil {
					return err
				}
				if got.TaskID.String() != "89695" {
					t.Fatalf("query flow = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) {
				requireKeys(t, body, "instanceCodes", "startTime", "endTime", "billingType")
			},
		},
		{
			name:     "QueryFlowResult returns documented fields",
			wantPath: "/distribute/task/query-flow-result",
			dataJSON: `{"taskId":89695,"taskDesc":"desc","billingType":"FLOW","taskStatus":"SUCCESS","msg":"","billingValue":1024,"bandwidthList":[{"recordTime":1733297481137,"send":1.5,"receive":2.5}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.QueryFlowResult(ctx, &QueryFlowResultRequest{TaskID: "89695"})
				if err != nil {
					return err
				}
				if got.TaskID.String() != "89695" || got.BandwidthList[0].Send.String() != "1.5" {
					t.Fatalf("flow result = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "taskId") },
		},
		{
			name:     "Device monitor query returns typed metric lists",
			wantPath: "/resources/device-monitor-info-query",
			dataJSON: `[{"deviceIp":"1.1.1.1","ioLoadRate":[{"recordTime":1733297481137,"value":1.2}],"cpuLoadRate":[{"recordTime":1733297481138,"value":2.3}]}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.QueryDeviceMonitorInfo(ctx, &DeviceMonitorInfoQueryRequest{DeviceIPs: []string{"1.1.1.1"}})
				if err != nil {
					return err
				}
				if len(got) != 1 || got[0].IOLoadRate[0].Value.String() != "1.2" || got[0].CPULoadRate[0].RecordTime.String() != "1733297481138" {
					t.Fatalf("device monitor query = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "deviceIps") },
		},
		{
			name:     "Instance monitor info returns typed fields",
			wantPath: "/resources/instance-monitor-info",
			dataJSON: `[{"instanceCode":"VM1","instanceIp":"2.2.2.2","recordTime":1733297481137,"cpuRate":1.1,"gpuRate":2.2,"memRate":3.3,"storageRate":4.4}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetInstanceMonitorInfo(ctx, &InstanceMonitorInfoRequest{InstanceCodes: []string{"VM1"}})
				if err != nil {
					return err
				}
				if len(got) != 1 || got[0].InstanceIP != "2.2.2.2" || got[0].CPURate.String() != "1.1" {
					t.Fatalf("instance monitor = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "instanceCodes") },
		},
		{
			name:     "Instance app monitor info returns typed top lists",
			wantPath: "/resources/instance-app-monitor-info",
			dataJSON: `{"cpuRateTopTen":[{"appName":"app","metricInfoList":[{"recordTime":1733297481137,"value":10.08}]}],"memRateTopTen":[{"appName":"app","metricInfoList":[{"recordTime":1733297481138,"value":20.16}]}]}`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetInstanceAppMonitorInfo(ctx, &InstanceAppMonitorInfoRequest{InstanceCode: "VM1"})
				if err != nil {
					return err
				}
				if got.CPURateTopTen[0].AppName != "app" || got.MemRateTopTen[0].MetricInfoList[0].Value.String() != "20.16" {
					t.Fatalf("app monitor = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "instanceCode") },
		},
		{
			name:     "Instance metric detail returns typed metrics",
			wantPath: "/resources/instance/metric-detail",
			dataJSON: `[{"cpuUtilizationRate":1,"memoryUtilizationRate":2,"storageUtilizationRate":3,"ioReadSpeed":4,"ioWriteSpeed":5,"storageSum":6,"storageFree":7,"memorySum":8,"memoryFree":9,"cpuFreq":10,"memoryFreq":11,"temperature":12,"cpuLoad":13,"ioWait":14,"deviceUpTime":15,"gpuFreq":16,"gpuUtilizationRate":17}]`,
			call: func(ctx context.Context, c *Client) error {
				got, err := c.GetInstanceMetricDetail(ctx, &InstanceMetricDetailRequest{InstanceCodes: []string{"VM1"}, RecordTime: "2026-01-01 00:00:00"})
				if err != nil {
					return err
				}
				if len(got) != 1 || got[0].GPUUtilizationRate.String() != "17" || got[0].IODeviceWait.String() != "14" {
					t.Fatalf("metric detail = %#v", got)
				}
				return nil
			},
			assertBody: func(t *testing.T, body map[string]any) { requireKeys(t, body, "instanceCodes", "recordTime") },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runEndpointAuditCase(t, tt)
		})
	}
}

func TestOfficialPaginationRequestsUseRowsNotSize(t *testing.T) {
	tests := []endpointAuditCase{
		{
			name:     "ListInstances sends rows",
			wantPath: "/resources/instance/infos",
			dataJSON: `{"total":0,"pageData":[]}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListInstances(ctx, &ListInstancesRequest{PageRequest: PageRequest{Page: 1}, Rows: 10})
				return err
			},
			assertBody: assertRowsNotSize(10),
		},
		{
			name:     "ListInstanceTags sends rows",
			wantPath: "/resources/instance-tag/page",
			dataJSON: `{"total":0,"pageData":[]}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListInstanceTags(ctx, &InstanceTagRequest{PageRequest: PageRequest{Page: 1}, Rows: 10})
				return err
			},
			assertBody: assertRowsNotSize(10),
		},
		{
			name:     "ListDevices sends rows",
			wantPath: "/resources/device/page",
			dataJSON: `{"total":0,"pageData":[]}`,
			call: func(ctx context.Context, c *Client) error {
				_, err := c.ListDevices(ctx, &ListDevicesRequest{PageRequest: PageRequest{Page: 1}, Rows: 10})
				return err
			},
			assertBody: assertRowsNotSize(10),
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

func assertRowsNotSize(rows int) func(*testing.T, map[string]any) {
	return func(t *testing.T, body map[string]any) {
		t.Helper()
		requireKeys(t, body, "page", "rows")
		forbidKeys(t, body, "size")
		if body["rows"].(float64) != float64(rows) {
			t.Fatalf("rows = %#v, want %d", body["rows"], rows)
		}
	}
}
