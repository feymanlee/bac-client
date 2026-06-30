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
	SMSNotifyStatus      int            `json:"smsNotifyStatus,omitempty"`
	SMSNotifyObjects     []string       `json:"smsNotifyObjects,omitempty"`
	CallbackNotifyStatus int            `json:"callbackNotifyStatus,omitempty"`
	EnableStatus         int            `json:"enableStatus,omitempty"`
	Page                 int            `json:"page,omitempty"`
	Rows                 int            `json:"rows,omitempty"`
}

type AlarmStrategy struct {
	AlarmStrategyID   FlexibleString `json:"alarmStrategyId,omitempty"`
	AlarmStrategyName string         `json:"alarmStrategyName,omitempty"`
	Raw               RawObject      `json:"-"`
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

func (c *Client) RemoveAlarmStrategy(ctx context.Context, req *AlarmStrategyRequest) error {
	return c.Do(ctx, pathRemoveAlarmStrategy, req, nil)
}

func (c *Client) ListAlarmStrategies(ctx context.Context, req *AlarmStrategyRequest) (*Page[AlarmStrategy], error) {
	var resp Page[AlarmStrategy]
	if err := c.Do(ctx, pathListAlarmStrategies, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateAlarmStrategyEnableStatus(ctx context.Context, req *AlarmStrategyRequest) error {
	return c.Do(ctx, pathUpdateAlarmStrategyEnable, req, nil)
}

type OpenAccountRequest struct {
	MerchantName string    `json:"merchantName,omitempty"`
	ContactName  string    `json:"contactName,omitempty"`
	Phone        string    `json:"phone,omitempty"`
	Email        string    `json:"email,omitempty"`
	Extra        RawObject `json:"extra,omitempty"`
}

func (c *Client) OpenAccount(ctx context.Context, req *OpenAccountRequest) error {
	return c.Do(ctx, pathOpenAccount, req, nil)
}

type AddSubMerchantRequest struct {
	MerchantNo   FlexibleString `json:"merchantNo,omitempty"`
	MerchantName string         `json:"merchantName,omitempty"`
	ContactName  string         `json:"contactName,omitempty"`
	Phone        string         `json:"phone,omitempty"`
	Email        string         `json:"email,omitempty"`
	Extra        RawObject      `json:"extra,omitempty"`
}

func (c *Client) AddSubMerchant(ctx context.Context, req *AddSubMerchantRequest) error {
	return c.Do(ctx, pathAddSubMerchant, req, nil)
}
