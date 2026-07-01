package bac

import (
	"context"
	"encoding/json"
)

const (
	pathListInstances        = "/resources/instance/infos"
	pathInstanceDeviceInfo   = "/resources/instance/device-info"
	pathUploadFile           = "/resources/instance/file-upload"
	pathDownloadFile         = "/resources/instance/file-download"
	pathBackupInstance       = "/resources/instance/backup"
	pathRestoreInstance      = "/resources/instance/restore"
	pathSessionControlSwitch = "/resources/instance/session-control-switch"
	pathSetResolution        = "/resources/instance/resolution"
	pathInstallApp           = "/resources/instance/install-app"
	pathUninstallApp         = "/resources/instance/uninstall-app"
	pathRebootRemotePlay     = "/resources/instance/reboot-remote-play"
	pathRebootDevice         = "/resources/device/reboot"
	pathNewPad               = "/distribute/pad/new-pad.html"
	pathCleanAppCache        = "/resources/device/clean-app-cache"
	pathBatchExecuteScript   = "/resources/instance/batch-execute-script"
	pathBindInfos            = "/resources/instance/bind-infos"
	pathCustomCodeUpdate     = "/resources/instance/custom-code-update"
	pathDataCopy             = "/resources/instance/data-copy"
	pathDeployMarketingSuite = "/resources/instance/deploy-marketing-suite"
	pathDestroyScreenshotURL = "/resources/instance/destroy-screenshot-url"
	pathEventTaskStop        = "/resources/instance/event-task-stop"
	pathExpireTimeIncrease   = "/resources/instance/expire-time-increase"
	pathExpireTimeUpdate     = "/resources/instance/expire-time-update"
	pathExportIP             = "/resources/instance/export-ip"
	pathGetScreenshotURL     = "/resources/instance/get-screenshot-url"
	pathMemoryLimit          = "/resources/instance/memory-limit"
	pathNetworkProxyWorkflow = "/resources/instance/network-proxy-workflow-create"
	pathSetSpeed             = "/resources/instance/set-speed"
	pathSSHInfo              = "/resources/instance/ssh-info"
	pathSummaryData          = "/resources/instance/summary-data"
	pathUpdateMaintainStatus = "/resources/instance/update-maintain-status"
	pathDeviceImageUpdate    = "/resources/device/image-update"
	pathDevicePage           = "/resources/device/page"
)

type ListInstancesRequest struct {
	PageRequest
	Rows                int              `json:"rows,omitempty"`
	InstanceID          string           `json:"instanceId,omitempty"`
	InstanceIDs         []string         `json:"instanceIds,omitempty"`
	InstanceCode        string           `json:"instanceCode,omitempty"`
	InstanceCodes       []string         `json:"instanceCodes,omitempty"`
	Status              string           `json:"status,omitempty"`
	PoolID              string           `json:"poolId,omitempty"`
	MerchantPoolNos     []FlexibleString `json:"merchantPoolNos,omitempty"`
	DeviceIPs           []string         `json:"deviceIps,omitempty"`
	DeviceIPSegment     string           `json:"deviceIpSegment,omitempty"`
	IDCCode             string           `json:"idcCode,omitempty"`
	InstanceServerType  string           `json:"instanceServerType,omitempty"`
	InstanceType        string           `json:"instanceType,omitempty"`
	RomVersion          string           `json:"romVersion,omitempty"`
	InstanceGradeName   string           `json:"instanceGradeName,omitempty"`
	UsableStatus        string           `json:"usableStatus,omitempty"`
	MaintainStatus      int              `json:"maintainStatus,omitempty"`
	RecycleStatus       string           `json:"recycleStatus,omitempty"`
	TaskStatus          string           `json:"taskStatus,omitempty"`
	InstanceStatus      int              `json:"instanceStatus,omitempty"`
	DeviceStatus        int              `json:"deviceStatus,omitempty"`
	MalfunctionStatus   int              `json:"malfunctionStatus,omitempty"`
	NetworkStatus       int              `json:"networkStatus,omitempty"`
	BindStatus          string           `json:"bindStatus,omitempty"`
	ControlStatus       string           `json:"controlStatus,omitempty"`
	ImageVersionID      FlexibleString   `json:"imageVersionId,omitempty"`
	SnapshotMountStatus string           `json:"snapshotMountStatus,omitempty"`
	SnapshotID          FlexibleString   `json:"snapshotId,omitempty"`
	TagID               FlexibleString   `json:"tagId,omitempty"`
	EmulationAuthStatus int              `json:"emulationAuthStatus,omitempty"`
}

type Instance struct {
	InstanceID          string         `json:"instanceId,omitempty"`
	InstanceCode        string         `json:"instanceCode,omitempty"`
	InstanceIP          string         `json:"instanceIp,omitempty"`
	DeviceIP            string         `json:"deviceIp,omitempty"`
	DeviceIPSegment     string         `json:"deviceIpSegment,omitempty"`
	DeviceCode          string         `json:"deviceCode,omitempty"`
	InstanceType        string         `json:"instanceType,omitempty"`
	InstanceServerType  string         `json:"instanceServerType,omitempty"`
	InstanceGrade       string         `json:"instanceGrade,omitempty"`
	RomVersion          string         `json:"romVersion,omitempty"`
	MerchantPoolNo      FlexibleString `json:"merchantPoolNo,omitempty"`
	IDCCode             string         `json:"idcCode,omitempty"`
	IDCName             string         `json:"idcName,omitempty"`
	AvailableStatus     string         `json:"availableStatus,omitempty"`
	UsableStatus        string         `json:"usableStatus,omitempty"`
	MaintainStatus      int            `json:"maintainStatus,omitempty"`
	RecycleStatus       string         `json:"recycleStatus,omitempty"`
	TaskStatus          string         `json:"taskStatus,omitempty"`
	InstanceStatus      int            `json:"instanceStatus,omitempty"`
	DeviceStatus        int            `json:"deviceStatus,omitempty"`
	MalfunctionStatus   int            `json:"malfunctionStatus,omitempty"`
	NetworkStatus       int            `json:"networkStatus,omitempty"`
	BindStatus          string         `json:"bindStatus,omitempty"`
	ControlStatus       string         `json:"controlStatus,omitempty"`
	ImageVersionID      FlexibleString `json:"imageVersionId,omitempty"`
	ImageVersionName    string         `json:"imageVersionName,omitempty"`
	SnapshotMountStatus string         `json:"snapshotMountStatus,omitempty"`
	SnapshotID          FlexibleString `json:"snapshotId,omitempty"`
	MemoryLimit         FlexibleString `json:"memoryLimit,omitempty"`
	UpSpeed             FlexibleString `json:"upSpeed,omitempty"`
	DownSpeed           FlexibleString `json:"downSpeed,omitempty"`
	IntranetUpSpeed     FlexibleString `json:"intranetUpSpeed,omitempty"`
	IntranetDownSpeed   FlexibleString `json:"intranetDownSpeed,omitempty"`
	AllocateTime        FlexibleString `json:"allocateTime,omitempty"`
	RecycleTime         FlexibleString `json:"recycleTime,omitempty"`
	Tags                []InstanceTag  `json:"tags,omitempty"`
	EmulationAuthStatus int            `json:"emulationAuthStatus,omitempty"`
	Name                string         `json:"name,omitempty"`
	Status              FlexibleString `json:"status,omitempty"`
	IP                  string         `json:"ip,omitempty"`
	Raw                 RawObject      `json:"-"`
}

func (i *Instance) UnmarshalJSON(data []byte) error {
	type alias Instance
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*i = Instance(a)
	return nil
}

func (c *Client) ListInstances(ctx context.Context, req *ListInstancesRequest) (*Page[Instance], error) {
	var resp Page[Instance]
	if err := c.Do(ctx, pathListInstances, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetInstanceDeviceInfoRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
}

type DeviceInfo struct {
	InstanceCode   string             `json:"instanceCode,omitempty"`
	PadCode        string             `json:"padCode,omitempty"`
	StorageSum     FlexibleString     `json:"storageSum,omitempty"`
	StorageFree    FlexibleString     `json:"storageFree,omitempty"`
	AppInstallList []DeviceAppInstall `json:"appInstallList,omitempty"`
	InstanceID     string             `json:"instanceId,omitempty"`
	AndroidID      string             `json:"androidId,omitempty"`
	IMEI           string             `json:"imei,omitempty"`
	IMSI           string             `json:"imsi,omitempty"`
	Mac            string             `json:"mac,omitempty"`
	Status         FlexibleString     `json:"status,omitempty"`
	Raw            RawObject          `json:"-"`
}

func (d *DeviceInfo) UnmarshalJSON(data []byte) error {
	type alias DeviceInfo
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*d = DeviceInfo(a)
	return nil
}

type DeviceAppInstall struct {
	AppID       FlexibleString `json:"appId,omitempty"`
	AppName     string         `json:"appName,omitempty"`
	PackageName string         `json:"packageName,omitempty"`
	VersionName string         `json:"versionName,omitempty"`
	AppIconURL  string         `json:"appIconUrl,omitempty"`
	Raw         RawObject      `json:"-"`
}

func (a *DeviceAppInstall) UnmarshalJSON(data []byte) error {
	type alias DeviceAppInstall
	var v alias
	if err := unmarshalRaw(data, (*RawObject)(&v.Raw), &v); err != nil {
		return err
	}
	*a = DeviceAppInstall(v)
	return nil
}

func (c *Client) GetInstanceDeviceInfo(ctx context.Context, req *GetInstanceDeviceInfoRequest) ([]DeviceInfo, error) {
	var resp []DeviceInfo
	if err := c.Do(ctx, pathInstanceDeviceInfo, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type UploadFileRequest struct {
	InstanceCodes     []string `json:"instanceCodes,omitempty"`
	FileURL           string   `json:"fileUrl,omitempty"`
	FileName          string   `json:"fileName,omitempty"`
	FileMD5           string   `json:"fileMd5,omitempty"`
	CustomizeFilePath string   `json:"customizeFilePath,omitempty"`
	AutoInstall       int      `json:"autoInstall,omitempty"`
}

type InstanceFileUpload struct {
	InstanceCode string `json:"instanceCode,omitempty"`
	FilePath     string `json:"filePath,omitempty"`
	FileURL      string `json:"fileUrl,omitempty"`
	TargetPath   string `json:"targetPath,omitempty"`
}

type UploadFileResponse struct {
	TaskID FlexibleString `json:"taskId,omitempty"`
	FileID FlexibleString `json:"fileId,omitempty"`
	Raw    RawObject      `json:"-"`
}

func (r *UploadFileResponse) UnmarshalJSON(data []byte) error {
	type alias UploadFileResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = UploadFileResponse(a)
	return nil
}

func (c *Client) UploadFile(ctx context.Context, req *UploadFileRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathUploadFile, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type DownloadFileRequest struct {
	InstanceFiles []InstanceFileDownload `json:"instanceFiles,omitempty"`
}

type InstanceFileDownload struct {
	InstanceCode string `json:"instanceCode,omitempty"`
	FilePath     string `json:"filePath,omitempty"`
}

func (c *Client) DownloadFile(ctx context.Context, req *DownloadFileRequest) (*BatchTaskResponse, error) {
	var resp taskResponseList
	if err := c.Do(ctx, pathDownloadFile, req, &resp); err != nil {
		return nil, err
	}
	return &BatchTaskResponse{SuccessList: []TaskResponse(resp)}, nil
}

type InstanceActionRequest struct {
	InstanceCode    string           `json:"instanceCode,omitempty"`
	InstanceCodes   []string         `json:"instanceCodes,omitempty"`
	PadCodes        []string         `json:"padCodes,omitempty"`
	MerchantPoolNos []FlexibleString `json:"merchantPoolNos,omitempty"`
	DeviceCodes     []string         `json:"deviceCodes,omitempty"`
	RebootType      int              `json:"rebootType,omitempty"`
	ResetType       string           `json:"resetType,omitempty"`
}

type TaskResponse struct {
	TaskID       FlexibleString `json:"taskId,omitempty"`
	InstanceCode string         `json:"instanceCode,omitempty"`
	PadCode      string         `json:"padCode,omitempty"`
	DeviceCode   string         `json:"deviceCode,omitempty"`
	DeviceIP     string         `json:"deviceIp,omitempty"`
	AuthCode     string         `json:"authCode,omitempty"`
	SnapshotID   FlexibleString `json:"snapshotId,omitempty"`
	ErrorMsg     string         `json:"errorMsg,omitempty"`
	Raw          RawObject      `json:"-"`
}

type taskResponseList []TaskResponse

func (l *taskResponseList) UnmarshalJSON(data []byte) error {
	var list []TaskResponse
	if err := json.Unmarshal(data, &list); err == nil {
		*l = list
		return nil
	}
	var batch BatchTaskResponse
	if err := json.Unmarshal(data, &batch); err != nil {
		return err
	}
	*l = batch.SuccessList
	return nil
}

func (r *TaskResponse) UnmarshalJSON(data []byte) error {
	type alias TaskResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = TaskResponse(a)
	return nil
}

type BatchTaskResponse struct {
	SuccessList []TaskResponse `json:"successList,omitempty"`
	FailList    []TaskFailure  `json:"failList,omitempty"`
	Raw         RawObject      `json:"-"`
}

func (r *BatchTaskResponse) UnmarshalJSON(data []byte) error {
	var list []TaskResponse
	if err := json.Unmarshal(data, &list); err == nil {
		r.SuccessList = list
		r.FailList = nil
		r.Raw = nil
		return nil
	}

	type alias BatchTaskResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = BatchTaskResponse(a)
	return nil
}

type TaskFailure struct {
	InstanceCode string    `json:"instanceCode,omitempty"`
	ServerToken  string    `json:"serverToken,omitempty"`
	Msg          string    `json:"msg,omitempty"`
	Raw          RawObject `json:"-"`
}

func (f *TaskFailure) UnmarshalJSON(data []byte) error {
	type alias TaskFailure
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*f = TaskFailure(a)
	return nil
}

type BackupInstanceRequest struct {
	InstanceCode string    `json:"instanceCode,omitempty"`
	SnapshotName string    `json:"snapshotName,omitempty"`
	OSSConfig    OSSConfig `json:"ossConfig,omitempty"`
	SnapshotPath string    `json:"snapshotPath,omitempty"`
	Excludes     []string  `json:"excludes,omitempty"`
	Includes     []string  `json:"includes,omitempty"`
}

type OSSConfig struct {
	Endpoint  string `json:"endpoint,omitempty"`
	Bucket    string `json:"bucket,omitempty"`
	AccessKey string `json:"accessKey,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`
	Protocol  string `json:"protocol,omitempty"`
}

type BackupInstanceResponse struct {
	TaskID       FlexibleString `json:"taskId,omitempty"`
	InstanceCode string         `json:"instanceCode,omitempty"`
	SnapshotID   FlexibleString `json:"snapshotId,omitempty"`
	SnapshotName string         `json:"snapshotName,omitempty"`
	Raw          RawObject      `json:"-"`
}

func (r *BackupInstanceResponse) UnmarshalJSON(data []byte) error {
	type alias BackupInstanceResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = BackupInstanceResponse(a)
	return nil
}

func (c *Client) BackupInstance(ctx context.Context, req *BackupInstanceRequest) (*BackupInstanceResponse, error) {
	var resp BackupInstanceResponse
	if err := c.Do(ctx, pathBackupInstance, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type RestoreInstanceRequest struct {
	SnapshotID    FlexibleString `json:"snapshotId,omitempty"`
	InstanceCodes []string       `json:"instanceCodes,omitempty"`
}

func (c *Client) RestoreInstance(ctx context.Context, req *RestoreInstanceRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathRestoreInstance, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type SessionControlSwitchRequest struct {
	ControlServerTokens []string `json:"controlServerTokens,omitempty"`
	WatchServerTokens   []string `json:"watchServerTokens,omitempty"`
}

type SessionControlSwitchFailure struct {
	ServerToken string    `json:"serverToken,omitempty"`
	ErrorMsg    string    `json:"errorMsg,omitempty"`
	Raw         RawObject `json:"-"`
}

func (f *SessionControlSwitchFailure) UnmarshalJSON(data []byte) error {
	type alias SessionControlSwitchFailure
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*f = SessionControlSwitchFailure(a)
	return nil
}

func (c *Client) SwitchSessionControl(ctx context.Context, req *SessionControlSwitchRequest) ([]SessionControlSwitchFailure, error) {
	var resp []SessionControlSwitchFailure
	if err := c.Do(ctx, pathSessionControlSwitch, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type SetResolutionRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
	Width         int      `json:"width,omitempty"`
	Height        int      `json:"height,omitempty"`
	DPI           int      `json:"dpi,omitempty"`
	FPS           int      `json:"fps,omitempty"`
}

func (c *Client) SetResolution(ctx context.Context, req *SetResolutionRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathSetResolution, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type InstallAppRequest struct {
	AppID         FlexibleString `json:"appId,omitempty"`
	InstanceCodes []string       `json:"instanceCodes,omitempty"`
}

func (c *Client) InstallApp(ctx context.Context, req *InstallAppRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathInstallApp, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type UninstallAppRequest struct {
	AppID         FlexibleString `json:"appId,omitempty"`
	InstanceCodes []string       `json:"instanceCodes,omitempty"`
}

func (c *Client) UninstallApp(ctx context.Context, req *UninstallAppRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathUninstallApp, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) RebootInstance(ctx context.Context, req *InstanceActionRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathRebootRemotePlay, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type RebootDeviceRequest struct {
	DeviceCodes []string `json:"deviceCodes,omitempty"`
	RebootType  int      `json:"rebootType,omitempty"`
}

func (c *Client) RebootDevice(ctx context.Context, req *RebootDeviceRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathRebootDevice, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) ResetDevice(ctx context.Context, req *InstanceActionRequest) ([]TaskResponse, error) {
	return c.CommandReset(ctx, req)
}

type NewPadRequest struct {
	PadModels []PadModel `json:"padModels,omitempty"`
}

type PadModel struct {
	PadCode      string `json:"padCode,omitempty"`
	IMEI         string `json:"imei,omitempty"`
	SerialNo     string `json:"serialno,omitempty"`
	WiFiMac      string `json:"wifimac,omitempty"`
	AndroidID    string `json:"androidid,omitempty"`
	Model        string `json:"model,omitempty"`
	Brand        string `json:"brand,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
}

func (c *Client) NewPad(ctx context.Context, req *NewPadRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathNewPad, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type CleanAppCacheRequest struct {
	DeviceIPs []string `json:"DeviceIps,omitempty"`
}

func (c *Client) CleanAppCache(ctx context.Context, req *CleanAppCacheRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathCleanAppCache, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type BatchExecuteScriptRequest struct {
	Scripts []InstanceScript `json:"scripts,omitempty"`
}

type InstanceScript struct {
	InstanceCode  string `json:"instanceCode,omitempty"`
	ScriptContent string `json:"scriptContent,omitempty"`
}

func (c *Client) BatchExecuteScript(ctx context.Context, req *BatchExecuteScriptRequest) (*BatchTaskResponse, error) {
	var resp BatchTaskResponse
	if err := c.Do(ctx, pathBatchExecuteScript, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type BindInfosRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
}

type BindInfo struct {
	InstanceCode        string         `json:"instanceCode,omitempty"`
	ServerToken         string         `json:"serverToken,omitempty"`
	UUID                string         `json:"uuid,omitempty"`
	BusinessType        string         `json:"businessType,omitempty"`
	BindTime            FlexibleString `json:"bindTime,omitempty"`
	ExpireTime          FlexibleString `json:"expireTime,omitempty"`
	ControlStatus       string         `json:"controlStatus,omitempty"`
	GrantControl        string         `json:"grantControl,omitempty"`
	InstanceControlTime FlexibleString `json:"instanceControlTime,omitempty"`
	Raw                 RawObject      `json:"-"`
}

func (b *BindInfo) UnmarshalJSON(data []byte) error {
	type alias BindInfo
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*b = BindInfo(a)
	return nil
}

func (c *Client) GetBindInfos(ctx context.Context, req *BindInfosRequest) ([]BindInfo, error) {
	var resp []BindInfo
	if err := c.Do(ctx, pathBindInfos, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type CustomCodeUpdateRequest struct {
	InstanceCustomList []InstanceCustom `json:"instanceCustomList,omitempty"`
}

type InstanceCustom struct {
	InstanceCode string `json:"instanceCode,omitempty"`
	CustomCode   string `json:"customCode,omitempty"`
}

func (c *Client) UpdateCustomCode(ctx context.Context, req *CustomCodeUpdateRequest) (*BatchTaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathCustomCodeUpdate, req, &resp); err != nil {
		return nil, err
	}
	return &BatchTaskResponse{SuccessList: resp}, nil
}

type DataCopyRequest struct {
	DataCopyList []DataCopyItem `json:"dataCopyList,omitempty"`
	Includes     []string       `json:"includes,omitempty"`
	Excludes     []string       `json:"excludes,omitempty"`
	Reset        bool           `json:"reset"`
}

type DataCopyItem struct {
	SourceInstanceCode string `json:"sourceInstanceCode,omitempty"`
	TargetInstanceCode string `json:"targetInstanceCode,omitempty"`
	IMEI               string `json:"imei,omitempty"`
	SerialNo           string `json:"serialno,omitempty"`
	WiFiMac            string `json:"wifimac,omitempty"`
	AndroidID          string `json:"androidid,omitempty"`
	Model              string `json:"model,omitempty"`
	Brand              string `json:"brand,omitempty"`
	Manufacturer       string `json:"manufacturer,omitempty"`
}

type DataCopyResponse struct {
	SourceInstanceCode string         `json:"sourceInstanceCode,omitempty"`
	TargetInstanceCode string         `json:"targetInstanceCode,omitempty"`
	TaskID             FlexibleString `json:"taskId,omitempty"`
	Raw                RawObject      `json:"-"`
}

func (r *DataCopyResponse) UnmarshalJSON(data []byte) error {
	type alias DataCopyResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = DataCopyResponse(a)
	return nil
}

func (c *Client) CopyInstanceData(ctx context.Context, req *DataCopyRequest) ([]DataCopyResponse, error) {
	var resp []DataCopyResponse
	if err := c.Do(ctx, pathDataCopy, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type DeployMarketingSuiteRequest struct {
	InstanceCodes       []string `json:"instanceCodes,omitempty"`
	AuthCodes           []string `json:"authCodes,omitempty"`
	AppPackage          string   `json:"appPackage,omitempty"`
	UseMerchantAuthCode bool     `json:"useMerchantAuthCode,omitempty"`
}

func (c *Client) DeployMarketingSuite(ctx context.Context, req *DeployMarketingSuiteRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathDeployMarketingSuite, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type ScreenshotURLRequest struct {
	InstanceCode  string   `json:"instanceCode,omitempty"`
	InstanceCodes []string `json:"instanceCodes,omitempty"`
	FullQuality   int      `json:"fullQuality,omitempty"`
	Scale         int      `json:"scale,omitempty"`
	Rotate        int      `json:"rotate,omitempty"`
	PictureType   string   `json:"pictureType,omitempty"`
}

type ScreenshotURLResponse struct {
	SuccessList []ScreenshotURL `json:"successList,omitempty"`
	FailList    []string        `json:"failList,omitempty"`
	ExpireTime  FlexibleString  `json:"expireTime,omitempty"`
	Raw         RawObject       `json:"-"`
}

func (r *ScreenshotURLResponse) UnmarshalJSON(data []byte) error {
	type alias ScreenshotURLResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = ScreenshotURLResponse(a)
	return nil
}

type ScreenshotURL struct {
	InstanceCode  string    `json:"instanceCode,omitempty"`
	ScreenshotURL string    `json:"screenshotUrl,omitempty"`
	Raw           RawObject `json:"-"`
}

func (s *ScreenshotURL) UnmarshalJSON(data []byte) error {
	type alias ScreenshotURL
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*s = ScreenshotURL(a)
	return nil
}

func (c *Client) GetScreenshotURL(ctx context.Context, req *ScreenshotURLRequest) (*ScreenshotURLResponse, error) {
	var resp ScreenshotURLResponse
	if err := c.Do(ctx, pathGetScreenshotURL, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DestroyScreenshotURLRequest struct {
	InstanceCode string `json:"instanceCode,omitempty"`
}

func (c *Client) DestroyScreenshotURL(ctx context.Context, req *DestroyScreenshotURLRequest) error {
	return c.Do(ctx, pathDestroyScreenshotURL, req, nil)
}

type StopEventTasksRequest struct {
	TaskIDs []FlexibleString `json:"taskIds,omitempty"`
}

func (c *Client) StopEventTasks(ctx context.Context, req *StopEventTasksRequest) error {
	return c.Do(ctx, pathEventTaskStop, req, nil)
}

type IncreaseExpireTimeRequest struct {
	ServerTokens []string `json:"serverTokens,omitempty"`
	Time         int      `json:"time"`
}

type UpdateExpireTimeRequest struct {
	ServerTokens []string       `json:"serverTokens,omitempty"`
	ExpireTime   FlexibleString `json:"expireTime,omitempty"`
}

type ExpireTimeResponse struct {
	FailList []string  `json:"failList,omitempty"`
	Raw      RawObject `json:"-"`
}

func (r *ExpireTimeResponse) UnmarshalJSON(data []byte) error {
	type alias ExpireTimeResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = ExpireTimeResponse(a)
	return nil
}

func (c *Client) IncreaseExpireTime(ctx context.Context, req *IncreaseExpireTimeRequest) (*ExpireTimeResponse, error) {
	var resp ExpireTimeResponse
	if err := c.Do(ctx, pathExpireTimeIncrease, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateExpireTime(ctx context.Context, req *UpdateExpireTimeRequest) (*ExpireTimeResponse, error) {
	var resp ExpireTimeResponse
	if err := c.Do(ctx, pathExpireTimeUpdate, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ExportIPRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
}

type ExportIPResult struct {
	InstanceCode    string    `json:"instanceCode,omitempty"`
	PublicIPAddress string    `json:"publicIpAddress,omitempty"`
	Raw             RawObject `json:"-"`
}

func (r *ExportIPResult) UnmarshalJSON(data []byte) error {
	type alias ExportIPResult
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = ExportIPResult(a)
	return nil
}

func (c *Client) ExportIP(ctx context.Context, req *ExportIPRequest) ([]ExportIPResult, error) {
	var resp []ExportIPResult
	if err := c.Do(ctx, pathExportIP, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type MemoryLimitRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
	IsLimit       bool     `json:"isLimit"`
	MemoryLimit   int      `json:"memoryLimit,omitempty"`
}

func (c *Client) SetMemoryLimit(ctx context.Context, req *MemoryLimitRequest) (*BatchTaskResponse, error) {
	var resp taskResponseList
	if err := c.Do(ctx, pathMemoryLimit, req, &resp); err != nil {
		return nil, err
	}
	return &BatchTaskResponse{SuccessList: []TaskResponse(resp)}, nil
}

type NetworkProxyWorkflowRequest struct {
	NetworkProxyConfigs []NetworkProxyConfig `json:"networkProxyConfigs,omitempty"`
}

type NetworkProxyConfig struct {
	InstanceCode  string `json:"instanceCode,omitempty"`
	ProxyHost     string `json:"proxyHost,omitempty"`
	ProxyPort     int    `json:"proxyPort,omitempty"`
	ProxyUser     string `json:"proxyUser,omitempty"`
	ProxyPassword string `json:"proxyPassword,omitempty"`
	ProxyWhite    string `json:"proxyWhite,omitempty"`
}

func (c *Client) CreateNetworkProxyWorkflow(ctx context.Context, req *NetworkProxyWorkflowRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathNetworkProxyWorkflow, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type SetSpeedRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
	Direction     string   `json:"direction,omitempty"`
	Speed         float64  `json:"speed,omitempty"`
	IntranetSpeed float64  `json:"intranetSpeed,omitempty"`
}

func (c *Client) SetSpeed(ctx context.Context, req *SetSpeedRequest) (*BatchTaskResponse, error) {
	var resp taskResponseList
	if err := c.Do(ctx, pathSetSpeed, req, &resp); err != nil {
		return nil, err
	}
	return &BatchTaskResponse{SuccessList: []TaskResponse(resp)}, nil
}

type SSHInfoRequest struct {
	InstanceCode string `json:"instanceCode,omitempty"`
	ConnectType  int    `json:"connectType,omitempty"`
	LiveTime     int    `json:"liveTime,omitempty"`
}

type SSHInfoResponse struct {
	SSHInfo    SSHInfo        `json:"sshInfo,omitempty"`
	ExpireTime FlexibleString `json:"expireTime,omitempty"`
	Raw        RawObject      `json:"-"`
}

func (r *SSHInfoResponse) UnmarshalJSON(data []byte) error {
	type alias SSHInfoResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = SSHInfoResponse(a)
	return nil
}

type SSHInfo struct {
	InstanceCode string `json:"instanceCode,omitempty"`
	SSHCommand   string `json:"sshCommand,omitempty"`
	SSHPwd       string `json:"sshPwd,omitempty"`
}

func (c *Client) GetSSHInfo(ctx context.Context, req *SSHInfoRequest) (*SSHInfoResponse, error) {
	var resp SSHInfoResponse
	if err := c.Do(ctx, pathSSHInfo, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SummaryDataRequest struct {
	MerchantPoolNo FlexibleString `json:"merchantPoolNo,omitempty"`
	IncludeSubPool bool           `json:"includeSubPool"`
}

type SummaryData struct {
	TotalData     InstanceCountSummary `json:"totalData,omitempty"`
	AvailableData InstanceCountSummary `json:"availableData,omitempty"`
	BoundData     InstanceCountSummary `json:"boundData,omitempty"`
	Raw           RawObject            `json:"-"`
}

func (s *SummaryData) UnmarshalJSON(data []byte) error {
	type alias SummaryData
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*s = SummaryData(a)
	return nil
}

type InstanceCountSummary struct {
	TotalInstanceCount FlexibleString `json:"totalInstanceCount,omitempty"`
	IDCList            []IDCSummary   `json:"idcList,omitempty"`
	Raw                RawObject      `json:"-"`
}

func (s *InstanceCountSummary) UnmarshalJSON(data []byte) error {
	type alias InstanceCountSummary
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*s = InstanceCountSummary(a)
	return nil
}

type IDCSummary struct {
	IDCName       string         `json:"idcName,omitempty"`
	IDCCode       string         `json:"idcCode,omitempty"`
	InstanceCount FlexibleString `json:"instanceCount,omitempty"`
	Raw           RawObject      `json:"-"`
}

func (s *IDCSummary) UnmarshalJSON(data []byte) error {
	type alias IDCSummary
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*s = IDCSummary(a)
	return nil
}

func (c *Client) GetSummaryData(ctx context.Context, req *SummaryDataRequest) (*SummaryData, error) {
	var resp SummaryData
	if err := c.Do(ctx, pathSummaryData, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type MaintainStatusRequest struct {
	InstanceCodes  []string `json:"instanceCodes,omitempty"`
	MaintainStatus int      `json:"maintainStatus"`
}

func (c *Client) UpdateMaintainStatus(ctx context.Context, req *MaintainStatusRequest) error {
	return c.Do(ctx, pathUpdateMaintainStatus, req, nil)
}

type DeviceImageUpdateRequest struct {
	ImageVersionID    int            `json:"imageVersionId,omitempty"`
	DeviceIPs         []string       `json:"deviceIps,omitempty"`
	ConfigID          string         `json:"configId,omitempty"`
	ResourcePackageID FlexibleString `json:"resourcePackageId,omitempty"`
	Reset             bool           `json:"reset,omitempty"`
	AutoInstall       bool           `json:"autoInstall"`
}

type DeviceImageUpdateResult struct {
	DeviceIP string         `json:"deviceIp,omitempty"`
	TaskID   FlexibleString `json:"taskId,omitempty"`
	Raw      RawObject      `json:"-"`
}

func (r *DeviceImageUpdateResult) UnmarshalJSON(data []byte) error {
	type alias DeviceImageUpdateResult
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = DeviceImageUpdateResult(a)
	return nil
}

func (c *Client) UpdateDeviceImage(ctx context.Context, req *DeviceImageUpdateRequest) ([]DeviceImageUpdateResult, error) {
	var resp []DeviceImageUpdateResult
	if err := c.Do(ctx, pathDeviceImageUpdate, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type ListDevicesRequest struct {
	PageRequest
	Rows               int      `json:"rows,omitempty"`
	DeviceCodes        []string `json:"deviceCodes,omitempty"`
	DeviceIPs          []string `json:"deviceIps,omitempty"`
	DeviceIPSegment    string   `json:"deviceIpSegment,omitempty"`
	IDCCode            string   `json:"idcCode,omitempty"`
	InstanceServerType string   `json:"instanceServerType,omitempty"`
	RomVersion         string   `json:"romVersion,omitempty"`
	DeviceStatus       int      `json:"deviceStatus,omitempty"`
}

type Device struct {
	DeviceIP           string         `json:"deviceIp,omitempty"`
	DeviceIPSegment    string         `json:"deviceIpSegment,omitempty"`
	DeviceCode         string         `json:"deviceCode,omitempty"`
	IDCCode            string         `json:"idcCode,omitempty"`
	IDCName            string         `json:"idcName,omitempty"`
	InstanceServerType string         `json:"instanceServerType,omitempty"`
	RomVersion         string         `json:"romVersion,omitempty"`
	InstanceGrade      FlexibleString `json:"instanceGrade,omitempty"`
	InstanceGradeLimit FlexibleString `json:"instanceGradeLimit,omitempty"`
	AllocateTime       FlexibleString `json:"allocateTime,omitempty"`
	RecycleTime        FlexibleString `json:"recycleTime,omitempty"`
	MemoryModel        string         `json:"memoryModel,omitempty"`
	CPUModel           string         `json:"cpuModel,omitempty"`
	DiskModel          string         `json:"diskModel,omitempty"`
	DeviceStatus       int            `json:"deviceStatus,omitempty"`
	Raw                RawObject      `json:"-"`
}

func (d *Device) UnmarshalJSON(data []byte) error {
	type alias Device
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*d = Device(a)
	return nil
}

func (c *Client) ListDevices(ctx context.Context, req *ListDevicesRequest) (*Page[Device], error) {
	var resp Page[Device]
	if err := c.Do(ctx, pathDevicePage, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
