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

func (c *Client) ListEmulationAuthCodes(ctx context.Context, req *ListEmulationAuthCodesRequest) (*Page[RawObject], error) {
	var resp Page[RawObject]
	if err := c.Do(ctx, pathEmulationAuthCodePage, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
