package bac

import "context"

const pathEmulationAuthCodePage = "/resources/emulation/auth-code-page"

type ListEmulationAuthCodesRequest struct {
	PageRequest
	Rows         int            `json:"rows,omitempty"`
	AuthCode     string         `json:"authCode,omitempty"`
	InstanceCode string         `json:"instanceCode,omitempty"`
	Status       FlexibleString `json:"status,omitempty"`
}

func (c *Client) ListEmulationAuthCodes(ctx context.Context, req *ListEmulationAuthCodesRequest) (*Page[RawObject], error) {
	var resp Page[RawObject]
	if err := c.Do(ctx, pathEmulationAuthCodePage, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
