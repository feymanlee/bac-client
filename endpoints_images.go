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

type ListImagesRequest struct {
	PageRequest
	Rows           int            `json:"rows,omitempty"`
	ImageVersionID FlexibleString `json:"imageVersionId,omitempty"`
	ImageID        string         `json:"imageId,omitempty"`
	Name           string         `json:"name,omitempty"`
}

type ImageInfo struct {
	ImageID string         `json:"imageId,omitempty"`
	Name    string         `json:"name,omitempty"`
	Version string         `json:"version,omitempty"`
	Status  FlexibleString `json:"status,omitempty"`
	Raw     RawObject      `json:"-"`
}

func (i *ImageInfo) UnmarshalJSON(data []byte) error {
	type alias ImageInfo
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*i = ImageInfo(a)
	return nil
}

func (c *Client) ListBaseImages(ctx context.Context, req *ListImagesRequest) (*Page[ImageInfo], error) {
	var resp Page[ImageInfo]
	if err := c.Do(ctx, pathBaseImages, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) ListInstanceImages(ctx context.Context, req *ListImagesRequest) (*Page[ImageInfo], error) {
	var resp Page[ImageInfo]
	if err := c.Do(ctx, pathInstanceImages, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type ImageActionRequest struct {
	ImageVersionID FlexibleString `json:"imageVersionId,omitempty"`
	ImageName      string         `json:"imageName,omitempty"`
	ImageURL       string         `json:"imageUrl,omitempty"`
	ImageMD5       string         `json:"imageMd5,omitempty"`
	InstanceCodes  []string       `json:"instanceCodes,omitempty"`
}

type ImageUpdateInfoRequest struct {
	TaskIDs []FlexibleString `json:"taskIds,omitempty"`
}

func (c *Client) UploadImage(ctx context.Context, req *ImageActionRequest) (*TaskResponse, error) {
	var resp TaskResponse
	if err := c.Do(ctx, pathImageUpload, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetImageUploadInfo(ctx context.Context, req *ImageActionRequest) (*TaskResult, error) {
	var resp TaskResult
	if err := c.Do(ctx, pathImageUploadInfo, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateImage(ctx context.Context, req *ImageActionRequest) (*TaskResponse, error) {
	var resp TaskResponse
	if err := c.Do(ctx, pathImageUpdate, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) GetImageUpdateInfo(ctx context.Context, req *ImageUpdateInfoRequest) (*TaskResult, error) {
	var resp TaskResult
	if err := c.Do(ctx, pathImageUpdateInfo, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RemoveImage(ctx context.Context, req *ImageActionRequest) error {
	return c.Do(ctx, pathImageRemove, req, nil)
}
