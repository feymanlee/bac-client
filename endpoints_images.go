package bac

import "context"

const (
	pathBaseImages      = "/resources/instance-base-image/list"
	pathInstanceImages  = "/resources/instance-image/list"
	pathImageUpload     = "/resources/instance-image/upload"
	pathImageUploadInfo = "/resources/instance-image/upload-info"
	pathImageUpdate     = "/resources/instance-image/update"
	pathImageUpdateInfo = "/resources/instance-image/update-info"
	pathImageRemove     = "/resources/instance-image/remove"
)

type ListBaseImagesRequest struct {
	Page               int              `json:"page,omitempty"`
	Rows               int              `json:"rows,omitempty"`
	InstanceServerType string           `json:"instanceServerType,omitempty"`
	RomVersion         string           `json:"romVersion,omitempty"`
	ImageVersionIDs    []FlexibleString `json:"imageVersionIds,omitempty"`
}

type ListInstanceImagesRequest struct {
	Page               int              `json:"page,omitempty"`
	Rows               int              `json:"rows,omitempty"`
	ImageVersionID     FlexibleString   `json:"imageVersionId,omitempty"`
	ImageVersionIDs    []FlexibleString `json:"imageVersionIds,omitempty"`
	InstanceServerType string           `json:"instanceServerType,omitempty"`
	RomVersion         string           `json:"romVersion,omitempty"`
	ImageVersionName   string           `json:"imageVersionName,omitempty"`
}

type ListImagesRequest = ListInstanceImagesRequest

type BaseImageInfo struct {
	ImageVersionID     FlexibleString `json:"imageVersionId,omitempty"`
	ImageVersionName   string         `json:"imageVersionName,omitempty"`
	InstanceServerType string         `json:"instanceServerType,omitempty"`
	RomVersion         string         `json:"romVersion,omitempty"`
	CreateTime         FlexibleString `json:"createTime,omitempty"`
	Raw                RawObject      `json:"-"`
}

func (i *BaseImageInfo) UnmarshalJSON(data []byte) error {
	type alias BaseImageInfo
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*i = BaseImageInfo(a)
	return nil
}

type InstanceImageInfo struct {
	ImageVersionID     FlexibleString `json:"imageVersionId,omitempty"`
	ImageUploadStatus  string         `json:"imageUploadStatus,omitempty"`
	InstanceServerType string         `json:"instanceServerType,omitempty"`
	RomVersion         string         `json:"romVersion,omitempty"`
	ImageFiles         []ImageFile    `json:"imageFiles,omitempty"`
	BaseImageVersionID FlexibleString `json:"baseImageVersionId,omitempty"`
	ImageVersionName   string         `json:"imageVersionName,omitempty"`
	Describe           string         `json:"describe,omitempty"`
	CreateTime         FlexibleString `json:"createTime,omitempty"`
	Raw                RawObject      `json:"-"`
}

func (i *InstanceImageInfo) UnmarshalJSON(data []byte) error {
	type alias InstanceImageInfo
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*i = InstanceImageInfo(a)
	return nil
}

type ImageInfo = InstanceImageInfo

func (c *Client) ListBaseImages(ctx context.Context, req *ListBaseImagesRequest) (*Page[BaseImageInfo], error) {
	var resp Page[BaseImageInfo]
	if err := c.Do(ctx, pathBaseImages, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListInstanceImages(ctx context.Context, req *ListInstanceImagesRequest) (*Page[InstanceImageInfo], error) {
	var resp Page[InstanceImageInfo]
	if err := c.Do(ctx, pathInstanceImages, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UploadImageRequest struct {
	ImageFiles         []ImageFile    `json:"imageFiles,omitempty"`
	InstanceServerType string         `json:"instanceServerType,omitempty"`
	RomVersion         string         `json:"romVersion,omitempty"`
	BaseImageVersionID FlexibleString `json:"baseImageVersionId,omitempty"`
	ImageVersionName   string         `json:"imageVersionName,omitempty"`
	Describe           string         `json:"describe,omitempty"`
}

type ImageFile struct {
	ImageFileURL  string `json:"imageFileUrl,omitempty"`
	ImageFileName string `json:"imageFileName,omitempty"`
	ImageFileType string `json:"imageFileType,omitempty"`
	ImageFileMD5  string `json:"imageFileMd5,omitempty"`
}

type UploadImageResponse struct {
	ImageVersionID FlexibleString `json:"imageVersionId,omitempty"`
	Raw            RawObject      `json:"-"`
}

func (r *UploadImageResponse) UnmarshalJSON(data []byte) error {
	type alias UploadImageResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = UploadImageResponse(a)
	return nil
}

type ImageUploadInfoRequest struct {
	ImageVersionID FlexibleString `json:"imageVersionId,omitempty"`
}

type ImageUploadInfo struct {
	ImageUploadStatus string    `json:"imageUploadStatus,omitempty"`
	Raw               RawObject `json:"-"`
}

func (r *ImageUploadInfo) UnmarshalJSON(data []byte) error {
	type alias ImageUploadInfo
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = ImageUploadInfo(a)
	return nil
}

type RemoveImageRequest struct {
	ImageVersionIDs []FlexibleString `json:"imageVersionIds,omitempty"`
}

type UpdateImageRequest struct {
	ImageVersionID    FlexibleString `json:"imageVersionId,omitempty"`
	InstanceCodes     []string       `json:"instanceCodes,omitempty"`
	ConfigID          string         `json:"configId,omitempty"`
	ResourcePackageID FlexibleString `json:"resourcePackageId,omitempty"`
	Reset             bool           `json:"reset,omitempty"`
	AutoInstall       *bool          `json:"autoInstall,omitempty"`
}

type ImageUpdateInfoRequest struct {
	TaskIDs []FlexibleString `json:"taskIds,omitempty"`
}

type ImageUpdateInfo struct {
	TaskID            FlexibleString `json:"taskId,omitempty"`
	InstanceCode      string         `json:"instanceCode,omitempty"`
	ImageUpdateStatus string         `json:"imageUpdateStatus,omitempty"`
	Raw               RawObject      `json:"-"`
}

func (r *ImageUpdateInfo) UnmarshalJSON(data []byte) error {
	type alias ImageUpdateInfo
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = ImageUpdateInfo(a)
	return nil
}

func (c *Client) UploadImage(ctx context.Context, req *UploadImageRequest) (*UploadImageResponse, error) {
	var resp UploadImageResponse
	if err := c.Do(ctx, pathImageUpload, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetImageUploadInfo(ctx context.Context, req *ImageUploadInfoRequest) (*ImageUploadInfo, error) {
	var resp ImageUploadInfo
	if err := c.Do(ctx, pathImageUploadInfo, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateImage(ctx context.Context, req *UpdateImageRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathImageUpdate, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetImageUpdateInfo(ctx context.Context, req *ImageUpdateInfoRequest) ([]ImageUpdateInfo, error) {
	var resp []ImageUpdateInfo
	if err := c.Do(ctx, pathImageUpdateInfo, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) RemoveImage(ctx context.Context, req *RemoveImageRequest) error {
	return c.Do(ctx, pathImageRemove, req, nil)
}
