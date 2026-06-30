package bac

import "context"

const (
	pathMalfunctionStatistics  = "/monitor/instance/malfunction-statistics"
	pathNetworkBandwidthList   = "/monitor/network-bandwidth/list"
	pathAvailablePadCount      = "/monitor/pad/available-count.html"
	pathEnableBindPadCount     = "/monitor/pad/enable-bind-count.html"
	pathDeviceMonitorInfo      = "/resources/device-monitor-info"
	pathDeviceMonitorInfoQuery = "/resources/device-monitor-info-query"
	pathInstanceMonitorInfo    = "/resources/instance-monitor-info"
	pathInstanceAppMonitorInfo = "/resources/instance-app-monitor-info"
	pathInstanceMetricDetail   = "/resources/instance/metric-detail"
)

type MalfunctionStatisticsRequest struct {
	StartTime          string `json:"startTime,omitempty"`
	EndTime            string `json:"endTime,omitempty"`
	TimeUnit           string `json:"timeUnit,omitempty"`
	IncludeSubMerchant bool   `json:"includeSubMerchant,omitempty"`
}

type MalfunctionStatistic struct {
	StatisticsTime          string    `json:"statisticsTime,omitempty"`
	ServerType              string    `json:"serverType,omitempty"`
	InstanceMalfunctionRate float64   `json:"instanceMalfunctionRate,omitempty"`
	Raw                     RawObject `json:"-"`
}

func (m *MalfunctionStatistic) UnmarshalJSON(data []byte) error {
	type alias MalfunctionStatistic
	var v alias
	if err := unmarshalRaw(data, (*RawObject)(&v.Raw), &v); err != nil {
		return err
	}
	*m = MalfunctionStatistic(v)
	return nil
}

func (c *Client) GetMalfunctionStatistics(ctx context.Context, req *MalfunctionStatisticsRequest) ([]MalfunctionStatistic, error) {
	var resp []MalfunctionStatistic
	if err := c.Do(ctx, pathMalfunctionStatistics, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type NetworkBandwidthRequest struct {
	IDCCode   string `json:"idcCode,omitempty"`
	BeginTime string `json:"beginTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	StatUnit  string `json:"statUnit,omitempty"`
}

type NetworkBandwidthListRequest = NetworkBandwidthRequest

type NetworkBandwidthStats struct {
	Peek          FlexibleString          `json:"peek,omitempty"`
	Peek95        FlexibleString          `json:"peek95,omitempty"`
	Average       FlexibleString          `json:"average,omitempty"`
	BandwidthList []NetworkBandwidthPoint `json:"bandwidthList,omitempty"`
	Raw           RawObject               `json:"-"`
}

func (r *NetworkBandwidthStats) UnmarshalJSON(data []byte) error {
	type alias NetworkBandwidthStats
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = NetworkBandwidthStats(a)
	return nil
}

type NetworkBandwidthPoint struct {
	RecordTime string         `json:"recordTime,omitempty"`
	Bandwidth  FlexibleString `json:"bandwidth,omitempty"`
	Raw        RawObject      `json:"-"`
}

func (r *NetworkBandwidthPoint) UnmarshalJSON(data []byte) error {
	type alias NetworkBandwidthPoint
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = NetworkBandwidthPoint(a)
	return nil
}

func (c *Client) ListNetworkBandwidth(ctx context.Context, req *NetworkBandwidthRequest) (*NetworkBandwidthStats, error) {
	var resp NetworkBandwidthStats
	if err := c.Do(ctx, pathNetworkBandwidthList, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type PadCountRequest struct {
	IDCCodes       []string       `json:"idcCodes,omitempty"`
	ServerType     string         `json:"serverType,omitempty"`
	MerchantPoolNo FlexibleString `json:"merchantPoolNo,omitempty"`
}

func (c *Client) GetAvailablePadCount(ctx context.Context, req *PadCountRequest) (RawObject, error) {
	var resp RawObject
	if err := c.Do(ctx, pathAvailablePadCount, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetEnableBindPadCount(ctx context.Context, req *PadCountRequest) (RawObject, error) {
	var resp RawObject
	if err := c.Do(ctx, pathEnableBindPadCount, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type DeviceMonitorInfoRequest struct {
	DeviceIPs []string `json:"deviceIps,omitempty"`
}

type DeviceMonitorInfoQueryRequest struct {
	DeviceIPs []string `json:"deviceIps,omitempty"`
	StartTime string   `json:"startTime,omitempty"`
	EndTime   string   `json:"endTime,omitempty"`
}

type DeviceMonitorInfo struct {
	DeviceIP    string         `json:"deviceIp,omitempty"`
	RecordTime  FlexibleString `json:"recordTime,omitempty"`
	CPURate     float64        `json:"cpuRate,omitempty"`
	GPURate     float64        `json:"gpuRate,omitempty"`
	MemRate     float64        `json:"memRate,omitempty"`
	StorageRate float64        `json:"storageRate,omitempty"`
	Raw         RawObject      `json:"-"`
}

func (d *DeviceMonitorInfo) UnmarshalJSON(data []byte) error {
	type alias DeviceMonitorInfo
	var v alias
	if err := unmarshalRaw(data, (*RawObject)(&v.Raw), &v); err != nil {
		return err
	}
	*d = DeviceMonitorInfo(v)
	return nil
}

func (c *Client) GetDeviceMonitorInfo(ctx context.Context, req *DeviceMonitorInfoRequest) ([]DeviceMonitorInfo, error) {
	var resp []DeviceMonitorInfo
	if err := c.Do(ctx, pathDeviceMonitorInfo, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) QueryDeviceMonitorInfo(ctx context.Context, req *DeviceMonitorInfoQueryRequest) ([]RawObject, error) {
	var resp []RawObject
	if err := c.Do(ctx, pathDeviceMonitorInfoQuery, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type InstanceMonitorInfoRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
}

type InstanceAppMonitorInfoRequest struct {
	InstanceCode string `json:"instanceCode,omitempty"`
	StartTime    string `json:"startTime,omitempty"`
	EndTime      string `json:"endTime,omitempty"`
}

func (c *Client) GetInstanceMonitorInfo(ctx context.Context, req *InstanceMonitorInfoRequest) ([]RawObject, error) {
	var resp []RawObject
	if err := c.Do(ctx, pathInstanceMonitorInfo, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type InstanceAppMonitorInfo struct {
	CPURateTopTen []RawObject `json:"cpuRateTopTen,omitempty"`
	MemRateTopTen []RawObject `json:"memRateTopTen,omitempty"`
	Raw           RawObject   `json:"-"`
}

func (r *InstanceAppMonitorInfo) UnmarshalJSON(data []byte) error {
	type alias InstanceAppMonitorInfo
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = InstanceAppMonitorInfo(a)
	return nil
}

func (c *Client) GetInstanceAppMonitorInfo(ctx context.Context, req *InstanceAppMonitorInfoRequest) (*InstanceAppMonitorInfo, error) {
	var resp InstanceAppMonitorInfo
	if err := c.Do(ctx, pathInstanceAppMonitorInfo, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type InstanceMetricDetailRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
	RecordTime    string   `json:"recordTime,omitempty"`
}

func (c *Client) GetInstanceMetricDetail(ctx context.Context, req *InstanceMetricDetailRequest) ([]RawObject, error) {
	var resp []RawObject
	if err := c.Do(ctx, pathInstanceMetricDetail, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
