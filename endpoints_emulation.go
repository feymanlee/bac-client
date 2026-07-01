package bac

import "context"

const pathEmulationAuthCodePage = "/resources/emulation/auth-code-page"

type ListEmulationAuthCodesRequest struct {
	AuthCodes     []string `json:"authCodes,omitempty"`
	AuthStatus    *int     `json:"authStatus,omitempty"`
	InstanceCodes []string `json:"instanceCodes,omitempty"`
	PackageName   string   `json:"packageName,omitempty"`
	Page          int      `json:"page,omitempty"`
	Rows          int      `json:"rows,omitempty"`
}

type EmulationAuthCode struct {
	AuthCode     string         `json:"authCode,omitempty"`
	ExpireTime   FlexibleString `json:"expireTime,omitempty"`
	AuthStatus   int            `json:"authStatus,omitempty"`
	ActivateTime FlexibleString `json:"activateTime,omitempty"`
	PackageName  string         `json:"packageName,omitempty"`
	DeviceIP     string         `json:"deviceIp,omitempty"`
	InstanceIP   string         `json:"instanceIp,omitempty"`
	InstanceCode string         `json:"instanceCode,omitempty"`
	CreateTime   FlexibleString `json:"createTime,omitempty"`
	Raw          RawObject      `json:"-"`
}

func (a *EmulationAuthCode) UnmarshalJSON(data []byte) error {
	type alias EmulationAuthCode
	var v alias
	if err := unmarshalRaw(data, (*RawObject)(&v.Raw), &v); err != nil {
		return err
	}
	*a = EmulationAuthCode(v)
	return nil
}

func (c *Client) ListEmulationAuthCodes(ctx context.Context, req *ListEmulationAuthCodesRequest) (*Page[EmulationAuthCode], error) {
	var resp Page[EmulationAuthCode]
	if err := c.Do(ctx, pathEmulationAuthCodePage, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
