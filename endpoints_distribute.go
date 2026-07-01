package bac

import (
	"context"
)

const (
	pathQueryFlow       = "/distribute/task/query-flow"
	pathQueryFlowResult = "/distribute/task/query-flow-result"
)

type QueryFlowRequest struct {
	InstanceCodes []string `json:"instanceCodes,omitempty"`
	StartTime     string   `json:"startTime,omitempty"`
	EndTime       string   `json:"endTime,omitempty"`
	BillingType   string   `json:"billingType,omitempty"`
	TaskDesc      string   `json:"taskDesc,omitempty"`
}

type QueryFlowResponse struct {
	TaskID FlexibleString `json:"taskId,omitempty"`
	Raw    RawObject      `json:"-"`
}

func (r *QueryFlowResponse) UnmarshalJSON(data []byte) error {
	type alias QueryFlowResponse
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = QueryFlowResponse(a)
	return nil
}

func (c *Client) QueryFlow(ctx context.Context, req *QueryFlowRequest) (*QueryFlowResponse, error) {
	var resp QueryFlowResponse
	if err := c.Do(ctx, pathQueryFlow, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type QueryFlowResultRequest struct {
	TaskID FlexibleString `json:"taskId,omitempty"`
}

type QueryFlowResult struct {
	TaskID        FlexibleString       `json:"taskId,omitempty"`
	TaskDesc      string               `json:"taskDesc,omitempty"`
	BillingType   string               `json:"billingType,omitempty"`
	TaskStatus    string               `json:"taskStatus,omitempty"`
	Msg           string               `json:"msg,omitempty"`
	BillingValue  FlexibleString       `json:"billingValue,omitempty"`
	BandwidthList []FlowBandwidthPoint `json:"bandwidthList,omitempty"`
	Raw           RawObject            `json:"-"`
}

func (r *QueryFlowResult) UnmarshalJSON(data []byte) error {
	type alias QueryFlowResult
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*r = QueryFlowResult(a)
	return nil
}

type FlowBandwidthPoint struct {
	RecordTime FlexibleString `json:"recordTime,omitempty"`
	Send       FlexibleString `json:"send,omitempty"`
	Receive    FlexibleString `json:"receive,omitempty"`
	Raw        RawObject      `json:"-"`
}

func (p *FlowBandwidthPoint) UnmarshalJSON(data []byte) error {
	type alias FlowBandwidthPoint
	var a alias
	if err := unmarshalRaw(data, (*RawObject)(&a.Raw), &a); err != nil {
		return err
	}
	*p = FlowBandwidthPoint(a)
	return nil
}

func (c *Client) QueryFlowResult(ctx context.Context, req *QueryFlowResultRequest) (*QueryFlowResult, error) {
	var resp QueryFlowResult
	if err := c.Do(ctx, pathQueryFlowResult, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
