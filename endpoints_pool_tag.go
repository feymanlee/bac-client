package bac

import "context"

const (
	pathAllocateInstancePool    = "/resources/instance-pool/allocate"
	pathSaveInstancePool        = "/resources/instance-pool/save"
	pathListInstanceTags        = "/resources/instance-tag/page"
	pathSaveInstanceTag         = "/resources/instance-tag/save"
	pathUpdateInstanceTag       = "/resources/instance-tag/update"
	pathRemoveInstanceTag       = "/resources/instance-tag/remove"
	pathAddInstanceTagRelate    = "/resources/instance-tag/relate-add"
	pathRemoveInstanceTagRelate = "/resources/instance-tag/relate-remove"
)

type InstancePoolRequest struct {
	MerchantPoolNo FlexibleString `json:"merchantPoolNo,omitempty"`
	AutoReset      bool           `json:"autoReset,omitempty"`
	InstanceCodes  []string       `json:"instanceCodes,omitempty"`
	Extra          RawObject      `json:"extra,omitempty"`
}

func (c *Client) AllocateInstancePool(ctx context.Context, req *InstancePoolRequest) error {
	return c.Do(ctx, pathAllocateInstancePool, req, nil)
}

type SaveInstancePoolRequest struct {
	ParentMerchantPoolNo FlexibleString `json:"parentMerchantPoolNo,omitempty"`
	InstancePoolName     string         `json:"instancePoolName,omitempty"`
	InstancePoolType     string         `json:"instancePoolType,omitempty"`
}

type SaveInstancePoolResponse struct {
	MerchantPoolNo FlexibleString `json:"merchantPoolNo,omitempty"`
	Raw            RawObject      `json:"-"`
}

func (r *SaveInstancePoolResponse) UnmarshalJSON(data []byte) error {
	type alias SaveInstancePoolResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = SaveInstancePoolResponse(a)
	return nil
}

func (c *Client) SaveInstancePool(ctx context.Context, req *SaveInstancePoolRequest) (*SaveInstancePoolResponse, error) {
	var resp SaveInstancePoolResponse
	if err := c.Do(ctx, pathSaveInstancePool, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstanceTagRequest struct {
	PageRequest
	Rows          int              `json:"rows,omitempty"`
	TagID         FlexibleString   `json:"tagId,omitempty"`
	TagIDs        []FlexibleString `json:"tagIds,omitempty"`
	TagName       string           `json:"tagName,omitempty"`
	Remark        string           `json:"remark,omitempty"`
	InstanceCodes []string         `json:"instanceCodes,omitempty"`
}

type InstanceTag struct {
	TagID                FlexibleString `json:"tagId,omitempty"`
	TagName              string         `json:"tagName,omitempty"`
	RelatedInstanceCount FlexibleString `json:"relatedInstanceCount,omitempty"`
	BelongToSubMerchant  bool           `json:"belongToSubMerchant,omitempty"`
	SubMerchantName      string         `json:"subMerchantName,omitempty"`
	Remark               string         `json:"remark,omitempty"`
	Raw                  RawObject      `json:"-"`
}

func (t *InstanceTag) UnmarshalJSON(data []byte) error {
	type alias InstanceTag
	var v alias
	if err := unmarshalRaw(data, (*RawObject)(&v.Raw), &v); err != nil {
		return err
	}
	*t = InstanceTag(v)
	return nil
}

func (c *Client) ListInstanceTags(ctx context.Context, req *InstanceTagRequest) (*Page[InstanceTag], error) {
	var resp Page[InstanceTag]
	if err := c.Do(ctx, pathListInstanceTags, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type SaveInstanceTagResponse struct {
	TagID FlexibleString `json:"tagId,omitempty"`
	Raw   RawObject      `json:"-"`
}

func (r *SaveInstanceTagResponse) UnmarshalJSON(data []byte) error {
	type alias SaveInstanceTagResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = SaveInstanceTagResponse(a)
	return nil
}

func (c *Client) SaveInstanceTag(ctx context.Context, req *InstanceTagRequest) (*SaveInstanceTagResponse, error) {
	var resp SaveInstanceTagResponse
	if err := c.Do(ctx, pathSaveInstanceTag, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateInstanceTag(ctx context.Context, req *InstanceTagRequest) error {
	return c.Do(ctx, pathUpdateInstanceTag, req, nil)
}

func (c *Client) RemoveInstanceTag(ctx context.Context, req *InstanceTagRequest) error {
	return c.Do(ctx, pathRemoveInstanceTag, req, nil)
}

func (c *Client) AddInstanceTagRelation(ctx context.Context, req *InstanceTagRequest) error {
	return c.Do(ctx, pathAddInstanceTagRelate, req, nil)
}

func (c *Client) RemoveInstanceTagRelation(ctx context.Context, req *InstanceTagRequest) error {
	return c.Do(ctx, pathRemoveInstanceTagRelate, req, nil)
}
