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
	PoolName       string         `json:"poolName,omitempty"`
	AutoReset      bool           `json:"autoReset,omitempty"`
	InstanceCodes  []string       `json:"instanceCodes,omitempty"`
	Extra          RawObject      `json:"extra,omitempty"`
}

func (c *Client) AllocateInstancePool(ctx context.Context, req *InstancePoolRequest) error {
	return c.Do(ctx, pathAllocateInstancePool, req, nil)
}

func (c *Client) SaveInstancePool(ctx context.Context, req *InstancePoolRequest) error {
	return c.Do(ctx, pathSaveInstancePool, req, nil)
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
	TagID   FlexibleString `json:"tagId,omitempty"`
	TagName string         `json:"tagName,omitempty"`
	Raw     RawObject      `json:"-"`
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

func (c *Client) SaveInstanceTag(ctx context.Context, req *InstanceTagRequest) error {
	return c.Do(ctx, pathSaveInstanceTag, req, nil)
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
