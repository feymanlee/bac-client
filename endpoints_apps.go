package bac

import "context"

const (
	pathUploadApps              = "/distribute/apps/uploads.html"
	pathDeleteApp               = "/resources/app/delete"
	pathListApps                = "/resources/app/page"
	pathNewAppVersion           = "/resources/app/new-version"
	pathUpgradeApp              = "/resources/app/upgrade"
	pathBuiltinInstallApp       = "/resources/app/app-builtin-install"
	pathBuiltinUninstallApp     = "/resources/app/app-builtin-uninstall"
	pathGetDesktopIconConfig    = "/resources/app/get-desktop-icon-config"
	pathSaveDesktopIconConfig   = "/resources/app/save-desktop-icon-config"
	pathRemoveDesktopIconConfig = "/resources/app/remove-desktop-icon-config"
	pathRecommendAppIconRefresh = "/resources/instance/recommend-app-icon-refresh"
	pathAppInstallList          = "/resources/instance/app-install-list"
)

type UploadAppsRequest struct {
	Apps   []UploadApp    `json:"apps,omitempty"`
	TaskID FlexibleString `json:"taskId,omitempty"`
}

type UploadApp struct {
	AppID   FlexibleString `json:"appId,omitempty"`
	AppName string         `json:"appName,omitempty"`
	PkgName string         `json:"pkgName,omitempty"`
	URL     string         `json:"url,omitempty"`
	MD5Sum  string         `json:"md5sum,omitempty"`
}

func (c *Client) UploadApps(ctx context.Context, req *UploadAppsRequest) error {
	return c.Do(ctx, pathUploadApps, req, nil)
}

type DeleteAppRequest struct {
	AppIDs []FlexibleString `json:"appIds,omitempty"`
}

func (c *Client) DeleteApp(ctx context.Context, req *DeleteAppRequest) error {
	return c.Do(ctx, pathDeleteApp, req, nil)
}

type ListAppsRequest struct {
	PageRequest
	Rows       int            `json:"rows,omitempty"`
	AppID      FlexibleString `json:"appId,omitempty"`
	AppPackage string         `json:"appPackage,omitempty"`
}

type AppInfo struct {
	AppID               FlexibleString `json:"appId,omitempty"`
	AppName             string         `json:"appName,omitempty"`
	AppPackage          string         `json:"appPackage,omitempty"`
	VersionName         string         `json:"versionName,omitempty"`
	AppTaskStatus       FlexibleString `json:"appTaskStatus,omitempty"`
	SyncStatus          FlexibleString `json:"syncStatus,omitempty"`
	ReleaseMarketStatus FlexibleString `json:"releaseMarketStatus,omitempty"`
	MD5Sum              string         `json:"md5sum,omitempty"`
	Raw                 RawObject      `json:"-"`
}

func (a *AppInfo) UnmarshalJSON(data []byte) error {
	type alias AppInfo
	var v alias
	if err := unmarshalRaw(data, (*RawObject)(&v.Raw), &v); err != nil {
		return err
	}
	*a = AppInfo(v)
	return nil
}

func (c *Client) ListApps(ctx context.Context, req *ListAppsRequest) (*Page[AppInfo], error) {
	var resp Page[AppInfo]
	if err := c.Do(ctx, pathListApps, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type AppVersionRequest struct {
	AppID          FlexibleString `json:"appId,omitempty"`
	AppName        string         `json:"appName,omitempty"`
	AppPackageName string         `json:"appPackageName,omitempty"`
	AppURL         string         `json:"appUrl,omitempty"`
}

func (c *Client) NewAppVersion(ctx context.Context, req *AppVersionRequest) error {
	return c.Do(ctx, pathNewAppVersion, req, nil)
}

func (c *Client) UpgradeApp(ctx context.Context, req *AppVersionRequest) error {
	return c.Do(ctx, pathUpgradeApp, req, nil)
}

type BuiltinAppRequest struct {
	InstanceCodes []string       `json:"instanceCodes,omitempty"`
	AppID         FlexibleString `json:"appId,omitempty"`
}

func (c *Client) BuiltinInstallApp(ctx context.Context, req *BuiltinAppRequest) (*BatchTaskResponse, error) {
	var resp BatchTaskResponse
	if err := c.Do(ctx, pathBuiltinInstallApp, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) BuiltinUninstallApp(ctx context.Context, req *BuiltinAppRequest) (*BatchTaskResponse, error) {
	var resp BatchTaskResponse
	if err := c.Do(ctx, pathBuiltinUninstallApp, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDesktopIconConfigRequest struct {
	InstanceCodes []string         `json:"instanceCodes,omitempty"`
	AppIDs        []FlexibleString `json:"appIds,omitempty"`
	Container     int              `json:"container,omitempty"`
	Screen        int              `json:"screen,omitempty"`
	X             int              `json:"x,omitempty"`
	Y             int              `json:"y,omitempty"`
}

type SaveDesktopIconConfigRequest struct {
	InstanceCodes       []string       `json:"instanceCodes,omitempty"`
	AppID               FlexibleString `json:"appId,omitempty"`
	Container           int            `json:"container"`
	Screen              int            `json:"screen"`
	X                   int            `json:"x"`
	Y                   int            `json:"y"`
	OverwriteCoordinate bool           `json:"overwriteCoordinate,omitempty"`
}

type RemoveDesktopIconConfigRequest struct {
	InstanceCodes []string         `json:"instanceCodes,omitempty"`
	AppIDs        []FlexibleString `json:"appIds,omitempty"`
}

type DesktopIconConfig struct {
	InstanceCode string         `json:"instanceCode,omitempty"`
	AppID        FlexibleString `json:"appId,omitempty"`
	Container    int            `json:"container,omitempty"`
	Screen       int            `json:"screen,omitempty"`
	X            int            `json:"x,omitempty"`
	Y            int            `json:"y,omitempty"`
	Raw          RawObject      `json:"-"`
}

func (r *DesktopIconConfig) UnmarshalJSON(data []byte) error {
	type alias DesktopIconConfig
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = DesktopIconConfig(a)
	return nil
}

func (c *Client) GetDesktopIconConfig(ctx context.Context, req *GetDesktopIconConfigRequest) ([]DesktopIconConfig, error) {
	var resp []DesktopIconConfig
	if err := c.Do(ctx, pathGetDesktopIconConfig, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) SaveDesktopIconConfig(ctx context.Context, req *SaveDesktopIconConfigRequest) error {
	return c.Do(ctx, pathSaveDesktopIconConfig, req, nil)
}

func (c *Client) RemoveDesktopIconConfig(ctx context.Context, req *RemoveDesktopIconConfigRequest) error {
	return c.Do(ctx, pathRemoveDesktopIconConfig, req, nil)
}

type RecommendAppIconRefreshRequest struct {
	InstanceCodes []string         `json:"instanceCodes,omitempty"`
	AppIDs        []FlexibleString `json:"appIds,omitempty"`
}

func (c *Client) RefreshRecommendAppIcons(ctx context.Context, req *RecommendAppIconRefreshRequest) (*BatchTaskResponse, error) {
	var resp BatchTaskResponse
	if err := c.Do(ctx, pathRecommendAppIconRefresh, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type AppInstallListRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
}

type InstalledAppList struct {
	InstanceCode         string               `json:"instanceCode,omitempty"`
	AppInstallRecordList []InstalledAppRecord `json:"appInstallRecordList,omitempty"`
	Raw                  RawObject            `json:"-"`
}

func (r *InstalledAppList) UnmarshalJSON(data []byte) error {
	type alias InstalledAppList
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = InstalledAppList(a)
	return nil
}

type InstalledAppRecord struct {
	AppName     string         `json:"appName,omitempty"`
	PackageName string         `json:"packageName,omitempty"`
	VersionName string         `json:"versionName,omitempty"`
	Size        FlexibleString `json:"size,omitempty"`
	Raw         RawObject      `json:"-"`
}

func (r *InstalledAppRecord) UnmarshalJSON(data []byte) error {
	type alias InstalledAppRecord
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = InstalledAppRecord(a)
	return nil
}

func (c *Client) ListInstalledApps(ctx context.Context, req *AppInstallListRequest) ([]InstalledAppList, error) {
	var resp []InstalledAppList
	if err := c.Do(ctx, pathAppInstallList, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
