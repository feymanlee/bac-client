package bac

import "context"

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

func (c *Client) QueryFlow(ctx context.Context, req *QueryFlowRequest) (RawObject, error) {
	var resp RawObject
	if err := c.Do(ctx, pathQueryFlow, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type QueryFlowResultRequest struct {
	TaskID FlexibleString `json:"taskId,omitempty"`
}

func (c *Client) QueryFlowResult(ctx context.Context, req *QueryFlowResultRequest) (RawObject, error) {
	var resp RawObject
	if err := c.Do(ctx, pathQueryFlowResult, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
