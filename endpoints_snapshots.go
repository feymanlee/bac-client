package bac

import "context"

const (
	pathDUFSSnapshotInit       = "/resources/dufs-snapshot/init"
	pathDUFSSnapshotMount      = "/resources/dufs-snapshot/mount"
	pathDUFSSnapshotBatchMount = "/resources/dufs-snapshot/batch-mount"
	pathDUFSSnapshotUnmount    = "/resources/dufs-snapshot/unmount"
	pathDUFSSnapshotQuotaSet   = "/resources/dufs-snapshot/quota-set"
	pathDUFSSnapshotPage       = "/resources/dufs-snapshot/page"
	pathDUFSSnapshotRemove     = "/resources/dufs-snapshot/remove"
	pathInstanceSnapshotPage   = "/resources/instance/snapshot-page"
)

type SnapshotRequest struct {
	SnapshotID     FlexibleString   `json:"snapshotId,omitempty"`
	SnapshotIDs    []FlexibleString `json:"snapshotIds,omitempty"`
	SnapshotName   string           `json:"snapshotName,omitempty"`
	SnapshotStatus string           `json:"snapshotStatus,omitempty"`
	InstanceCode   string           `json:"instanceCode,omitempty"`
	InstanceCodes  []string         `json:"instanceCodes,omitempty"`
	Page           int              `json:"page,omitempty"`
	Rows           int              `json:"rows,omitempty"`
	Quota          FlexibleString   `json:"quota,omitempty"`
}

type InitDUFSSnapshotRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
	QuotaCapacity int      `json:"quotaCapacity,omitempty"`
	MemoryLimit   int      `json:"memoryLimit,omitempty"`
}

type MountDUFSSnapshotRequest struct {
	InstanceCodes []string       `json:"instanceCodes,omitempty"`
	SnapshotID    FlexibleString `json:"snapshotId,omitempty"`
	QuotaCapacity int            `json:"quotaCapacity,omitempty"`
	MemoryLimit   int            `json:"memoryLimit,omitempty"`
}

type SnapshotMountInfo struct {
	InstanceCode  string         `json:"instanceCode,omitempty"`
	SnapshotID    FlexibleString `json:"snapshotId,omitempty"`
	QuotaCapacity int            `json:"quotaCapacity,omitempty"`
	MemoryLimit   int            `json:"memoryLimit,omitempty"`
}

type BatchMountDUFSSnapshotRequest struct {
	SnapshotMountInfos []SnapshotMountInfo `json:"snapshotMountInfos,omitempty"`
}

type UnmountDUFSSnapshotRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
	SnapshotName  string   `json:"snapshotName,omitempty"`
}

type SetDUFSSnapshotQuotaRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
	QuotaCapacity int      `json:"quotaCapacity,omitempty"`
}

type ListDUFSSnapshotsRequest struct {
	SnapshotName   string           `json:"snapshotName,omitempty"`
	SnapshotStatus string           `json:"snapshotStatus,omitempty"`
	SnapshotIDs    []FlexibleString `json:"snapshotIds,omitempty"`
	Page           int              `json:"page,omitempty"`
	Rows           int              `json:"rows,omitempty"`
}

type RemoveDUFSSnapshotRequest struct {
	SnapshotIDs []FlexibleString `json:"snapshotIds,omitempty"`
}

type ListInstanceSnapshotsRequest struct {
	SnapshotIDs    []FlexibleString `json:"snapshotIds,omitempty"`
	SnapshotName   string           `json:"snapshotName,omitempty"`
	SnapshotStatus string           `json:"snapshotStatus,omitempty"`
	Page           int              `json:"page,omitempty"`
	Rows           int              `json:"rows,omitempty"`
}

type SnapshotInfo struct {
	SnapshotID     FlexibleString `json:"snapshotId,omitempty"`
	SnapshotName   string         `json:"snapshotName,omitempty"`
	SnapshotStatus string         `json:"snapshotStatus,omitempty"`
	InstanceCode   string         `json:"instanceCode,omitempty"`
	IDCCode        string         `json:"idcCode,omitempty"`
	Progress       string         `json:"progress,omitempty"`
	SnapshotSize   FlexibleString `json:"snapshotSize,omitempty"`
	FailReason     string         `json:"failReason,omitempty"`
	Raw            RawObject      `json:"-"`
}

func (s *SnapshotInfo) UnmarshalJSON(data []byte) error {
	type alias SnapshotInfo
	var v alias
	if err := unmarshalRaw(data, (*RawObject)(&v.Raw), &v); err != nil {
		return err
	}
	*s = SnapshotInfo(v)
	return nil
}

func (c *Client) InitDUFSSnapshot(ctx context.Context, req *InitDUFSSnapshotRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathDUFSSnapshotInit, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) MountDUFSSnapshot(ctx context.Context, req *MountDUFSSnapshotRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathDUFSSnapshotMount, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) BatchMountDUFSSnapshot(ctx context.Context, req *BatchMountDUFSSnapshotRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathDUFSSnapshotBatchMount, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) UnmountDUFSSnapshot(ctx context.Context, req *UnmountDUFSSnapshotRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathDUFSSnapshotUnmount, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) SetDUFSSnapshotQuota(ctx context.Context, req *SetDUFSSnapshotQuotaRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathDUFSSnapshotQuotaSet, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) ListDUFSSnapshots(ctx context.Context, req *ListDUFSSnapshotsRequest) (*Page[SnapshotInfo], error) {
	var resp Page[SnapshotInfo]
	if err := c.Do(ctx, pathDUFSSnapshotPage, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) RemoveDUFSSnapshot(ctx context.Context, req *RemoveDUFSSnapshotRequest) error {
	return c.Do(ctx, pathDUFSSnapshotRemove, req, nil)
}

func (c *Client) ListInstanceSnapshots(ctx context.Context, req *ListInstanceSnapshotsRequest) (*Page[SnapshotInfo], error) {
	var resp Page[SnapshotInfo]
	if err := c.Do(ctx, pathInstanceSnapshotPage, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
