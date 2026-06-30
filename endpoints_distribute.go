package bac

import (
	"context"
	"encoding/json"
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
	Items  []RawObject    `json:"items,omitempty"`
	Raw    RawObject      `json:"-"`
}

func (r *QueryFlowResponse) UnmarshalJSON(data []byte) error {
	var list []RawObject
	if err := json.Unmarshal(data, &list); err == nil {
		r.Items = list
		if len(list) > 0 {
			if taskID, ok := list[0]["taskId"]; ok {
				_ = json.Unmarshal(taskID, &r.TaskID)
			}
		}
		return nil
	}
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

func (c *Client) QueryFlowResult(ctx context.Context, req *QueryFlowResultRequest) (RawObject, error) {
	var resp RawObject
	if err := c.Do(ctx, pathQueryFlowResult, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
