package bac

import "context"

const (
	pathAddAlarmStrategy          = "/merchant/alarm-strategy/add"
	pathUpdateAlarmStrategy       = "/merchant/alarm-strategy/update"
	pathRemoveAlarmStrategy       = "/merchant/alarm-strategy/remove"
	pathListAlarmStrategies       = "/merchant/alarm-strategy/page"
	pathUpdateAlarmStrategyEnable = "/merchant/alarm-strategy/update-enable-status"
	pathOpenAccount               = "/merchant/open-account/save"
	pathAddSubMerchant            = "/merchant/sub-merchant/add"
)

type AlarmStrategyRequest struct {
	AlarmStrategyID      FlexibleString `json:"alarmStrategyId,omitempty"`
	AlarmStrategyName    string         `json:"alarmStrategyName,omitempty"`
	AlarmStrategyDesc    string         `json:"alarmStrategyDesc,omitempty"`
	AlarmResourceType    string         `json:"alarmResourceType,omitempty"`
	IDCCodes             []string       `json:"idcCodes,omitempty"`
	AlarmMetrics         string         `json:"alarmMetrics,omitempty"`
	AlarmThreshold       int            `json:"alarmThreshold,omitempty"`
	AlarmSilencePeriod   int            `json:"alarmSilencePeriod,omitempty"`
	SMSNotifyStatus      *int           `json:"smsNotifyStatus,omitempty"`
	SMSNotifyObjects     []string       `json:"smsNotifyObjects,omitempty"`
	CallbackNotifyStatus *int           `json:"callbackNotifyStatus,omitempty"`
	EnableStatus         *int           `json:"enableStatus,omitempty"`
	Page                 int            `json:"page,omitempty"`
	Rows                 int            `json:"rows,omitempty"`
}

type AlarmStrategy struct {
	AlarmStrategyID      FlexibleString `json:"alarmStrategyId,omitempty"`
	AlarmStrategyName    string         `json:"alarmStrategyName,omitempty"`
	AlarmStrategyDesc    string         `json:"alarmStrategyDesc,omitempty"`
	IDCNames             []string       `json:"idcNames,omitempty"`
	AlarmResourceType    string         `json:"alarmResourceType,omitempty"`
	AlarmMetrics         string         `json:"alarmMetrics,omitempty"`
	AlarmThreshold       FlexibleString `json:"alarmThreshold,omitempty"`
	AlarmSilencePeriod   FlexibleString `json:"alarmSilencePeriod,omitempty"`
	SMSNotifyStatus      FlexibleString `json:"smsNotifyStatus,omitempty"`
	SMSNotifyObjects     []string       `json:"smsNotifyObjects,omitempty"`
	CallbackNotifyStatus FlexibleString `json:"callbackNotifyStatus,omitempty"`
	EnableStatus         FlexibleString `json:"enableStatus,omitempty"`
	CreateTime           FlexibleString `json:"createTime,omitempty"`
	Raw                  RawObject      `json:"-"`
}

func (a *AlarmStrategy) UnmarshalJSON(data []byte) error {
	type alias AlarmStrategy
	var v alias
	if err := unmarshalRaw(data, (*RawObject)(&v.Raw), &v); err != nil {
		return err
	}
	*a = AlarmStrategy(v)
	return nil
}

func (c *Client) AddAlarmStrategy(ctx context.Context, req *AlarmStrategyRequest) error {
	return c.Do(ctx, pathAddAlarmStrategy, req, nil)
}

func (c *Client) UpdateAlarmStrategy(ctx context.Context, req *AlarmStrategyRequest) error {
	return c.Do(ctx, pathUpdateAlarmStrategy, req, nil)
}

type RemoveAlarmStrategyRequest struct {
	AlarmStrategyIDs []FlexibleString `json:"alarmStrategyIds,omitempty"`
}

func (c *Client) RemoveAlarmStrategy(ctx context.Context, req *RemoveAlarmStrategyRequest) error {
	return c.Do(ctx, pathRemoveAlarmStrategy, req, nil)
}

func (c *Client) ListAlarmStrategies(ctx context.Context, req *AlarmStrategyRequest) (*Page[AlarmStrategy], error) {
	var resp Page[AlarmStrategy]
	if err := c.Do(ctx, pathListAlarmStrategies, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateAlarmStrategyEnableStatusRequest struct {
	AlarmStrategyIDs []FlexibleString `json:"alarmStrategyIds,omitempty"`
	EnableStatus     int              `json:"enableStatus"`
}

func (c *Client) UpdateAlarmStrategyEnableStatus(ctx context.Context, req *UpdateAlarmStrategyEnableStatusRequest) error {
	return c.Do(ctx, pathUpdateAlarmStrategyEnable, req, nil)
}

type OpenAccountRequest struct {
	UserName        string           `json:"userName,omitempty"`
	Phone           string           `json:"phone,omitempty"`
	Nickname        string           `json:"nickname,omitempty"`
	RoleNames       []string         `json:"roleNames,omitempty"`
	MerchantPoolNos []FlexibleString `json:"merchantPoolNos,omitempty"`
}

type OpenAccountResponse struct {
	Password string    `json:"password,omitempty"`
	Raw      RawObject `json:"-"`
}

func (r *OpenAccountResponse) UnmarshalJSON(data []byte) error {
	type alias OpenAccountResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = OpenAccountResponse(a)
	return nil
}

func (c *Client) OpenAccount(ctx context.Context, req *OpenAccountRequest) (*OpenAccountResponse, error) {
	var resp OpenAccountResponse
	if err := c.Do(ctx, pathOpenAccount, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddSubMerchantRequest struct {
	ParentMerchantCode string `json:"parentMerchantCode,omitempty"`
	MerchantCode       string `json:"merchantCode,omitempty"`
	MerchantName       string `json:"merchantName,omitempty"`
	MerchantType       string `json:"merchantType,omitempty"`
	MerchantPhone      string `json:"merchantPhone,omitempty"`
	AdminUserName      string `json:"adminUserName,omitempty"`
	AdminPhone         string `json:"adminPhone,omitempty"`
}

type AddSubMerchantResponse struct {
	AdminPassword string    `json:"adminPassword,omitempty"`
	Raw           RawObject `json:"-"`
}

func (r *AddSubMerchantResponse) UnmarshalJSON(data []byte) error {
	type alias AddSubMerchantResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = AddSubMerchantResponse(a)
	return nil
}

func (c *Client) AddSubMerchant(ctx context.Context, req *AddSubMerchantRequest) (*AddSubMerchantResponse, error) {
	var resp AddSubMerchantResponse
	if err := c.Do(ctx, pathAddSubMerchant, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
