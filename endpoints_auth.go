package bac

import "context"

const (
	pathCloudPhoneServerToken = "/auth/instance/cloud-phone-server-token"
	pathAuthorizedServerToken = "/auth/instance/authorized-server-token"
	pathDisconnectInstance    = "/auth/instance/disconnect"
	pathBatchDisconnect       = "/auth/instance/batch-disconnect"
	pathDisconnectAll         = "/auth/instance/disconnect-all"
	pathAuthorizedConnect     = "/auth/instance/authorized-connect"
)

type GetServerTokenRequest struct {
	UUID          string   `json:"uuid,omitempty"`
	InstanceCodes []string `json:"instanceCodes,omitempty"`
	OnlineTime    int      `json:"onlineTime,omitempty"`
	GrantControl  string   `json:"grantControl,omitempty"`
}

type ServerTokenResponse struct {
	ServerToken  string               `json:"serverToken,omitempty"`
	Token        string               `json:"token,omitempty"`
	InstanceCode string               `json:"instanceCode,omitempty"`
	ExpireTime   FlexibleString       `json:"expireTime,omitempty"`
	SuccessList  []ServerTokenSuccess `json:"successList,omitempty"`
	ErrorList    []ServerTokenError   `json:"errorList,omitempty"`
	Raw          RawObject            `json:"-"`
}

func (r *ServerTokenResponse) UnmarshalJSON(data []byte) error {
	type alias ServerTokenResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = ServerTokenResponse(a)
	return nil
}

type ServerTokenSuccess struct {
	InstanceCode string    `json:"instanceCode,omitempty"`
	ServerToken  string    `json:"serverToken,omitempty"`
	Raw          RawObject `json:"-"`
}

func (s *ServerTokenSuccess) UnmarshalJSON(data []byte) error {
	type alias ServerTokenSuccess
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*s = ServerTokenSuccess(a)
	return nil
}

type ServerTokenError struct {
	InstanceCode string    `json:"instanceCode,omitempty"`
	ErrorMsg     string    `json:"errorMsg,omitempty"`
	Raw          RawObject `json:"-"`
}

func (e *ServerTokenError) UnmarshalJSON(data []byte) error {
	type alias ServerTokenError
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*e = ServerTokenError(a)
	return nil
}

func (c *Client) GetServerToken(ctx context.Context, req *GetServerTokenRequest) (*ServerTokenResponse, error) {
	var resp ServerTokenResponse
	if err := c.Do(ctx, pathCloudPhoneServerToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetAuthorizedServerTokenRequest struct {
	UUID         string `json:"uuid,omitempty"`
	InstanceCode string `json:"instanceCode,omitempty"`
	OnlineTime   int    `json:"onlineTime,omitempty"`
	GrantControl string `json:"grantControl,omitempty"`
}

func (c *Client) GetAuthorizedServerToken(ctx context.Context, req *GetAuthorizedServerTokenRequest) (*ServerTokenResponse, error) {
	var resp ServerTokenResponse
	if err := c.Do(ctx, pathAuthorizedServerToken, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DisconnectInstanceRequest struct {
	UUID        string `json:"uuid,omitempty"`
	ServerToken string `json:"serverToken,omitempty"`
}

func (c *Client) DisconnectInstance(ctx context.Context, req *DisconnectInstanceRequest) error {
	return c.Do(ctx, pathDisconnectInstance, req, nil)
}

type BatchDisconnectInstancesRequest struct {
	DisconnectList []DisconnectCredential `json:"disconnectList,omitempty"`
}

type DisconnectCredential struct {
	UUID        string `json:"uuid,omitempty"`
	ServerToken string `json:"serverToken,omitempty"`
}

type BatchDisconnectResponse struct {
	SuccessList []string      `json:"successList,omitempty"`
	FailList    []TaskFailure `json:"failList,omitempty"`
	Raw         RawObject     `json:"-"`
}

func (r *BatchDisconnectResponse) UnmarshalJSON(data []byte) error {
	type alias BatchDisconnectResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = BatchDisconnectResponse(a)
	return nil
}

func (c *Client) BatchDisconnectInstances(ctx context.Context, req *BatchDisconnectInstancesRequest) (*BatchDisconnectResponse, error) {
	var resp BatchDisconnectResponse
	if err := c.Do(ctx, pathBatchDisconnect, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type DisconnectAllInstancesRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
}

func (c *Client) DisconnectAllInstances(ctx context.Context, req *DisconnectAllInstancesRequest) error {
	return c.Do(ctx, pathDisconnectAll, req, nil)
}

type AuthorizedConnectRequest struct {
	UUID         string `json:"uuid,omitempty"`
	InstanceCode string `json:"instanceCode,omitempty"`
	OnlineTime   int    `json:"onlineTime,omitempty"`
	GrantControl string `json:"grantControl,omitempty"`
}

type AuthorizedConnectResponse struct {
	URL         string         `json:"url,omitempty"`
	ServerURL   string         `json:"serverUrl,omitempty"`
	ServerToken string         `json:"serverToken,omitempty"`
	ExpireTime  FlexibleString `json:"expireTime,omitempty"`
	Raw         RawObject      `json:"-"`
}

func (r *AuthorizedConnectResponse) UnmarshalJSON(data []byte) error {
	type alias AuthorizedConnectResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = AuthorizedConnectResponse(a)
	return nil
}

func (c *Client) AuthorizedConnect(ctx context.Context, req *AuthorizedConnectRequest) (*AuthorizedConnectResponse, error) {
	var resp AuthorizedConnectResponse
	if err := c.Do(ctx, pathAuthorizedConnect, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
