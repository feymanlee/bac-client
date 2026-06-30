package bac

import "context"

const (
	pathScreenshot     = "/command/pad/screenshot.html"
	pathScreenshotInfo = "/command/pad/screenshot-info.html"
	pathAppOperate     = "/command/apps/app-operate.html"
	pathExecuteScript  = "/command/pad/execute-script.html"
	pathTaskInfo       = "/command/pad/execute-task-info.html"
	pathTaskPage       = "/command/pad/execute-task-page"
	pathTaskTypeList   = "/command/pad/execute-task-type-list"
	pathCommandReboot  = "/command/pad/reboot.html"
	pathCommandReset   = "/command/pad/reset.html"
)

type ScreenshotRequest struct {
	PadCodes    []string `json:"padCodes,omitempty"`
	Quality     int      `json:"quality,omitempty"`
	PictureType string   `json:"pictureType,omitempty"`
}

type ScreenshotResponse struct {
	TaskID FlexibleString `json:"taskId,omitempty"`
	URL    string         `json:"url,omitempty"`
	Raw    RawObject      `json:"-"`
}

func (r *ScreenshotResponse) UnmarshalJSON(data []byte) error {
	type alias ScreenshotResponse
	var a alias
	if err := unmarshalRaw(data, &a.Raw, &a); err != nil {
		return err
	}
	*r = ScreenshotResponse(a)
	return nil
}

func (c *Client) Screenshot(ctx context.Context, req *ScreenshotRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathScreenshot, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type ScreenshotResultRequest struct {
	TaskIDs []FlexibleString `json:"taskIds,omitempty"`
}

func (c *Client) GetScreenshotResult(ctx context.Context, req *ScreenshotResultRequest) ([]TaskResult, error) {
	var resp []TaskResult
	if err := c.Do(ctx, pathScreenshotInfo, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type AppOperateRequest struct {
	PadCodes    []string `json:"padCodes,omitempty"`
	PackageName string   `json:"packageName,omitempty"`
	OperateType string   `json:"operateType,omitempty"`
}

func (c *Client) OperateApp(ctx context.Context, req *AppOperateRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathAppOperate, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) StartApp(ctx context.Context, req *AppOperateRequest) ([]TaskResponse, error) {
	cp := AppOperateRequest{}
	if req != nil {
		cp = *req
	}
	cp.OperateType = "start"
	return c.OperateApp(ctx, &cp)
}

func (c *Client) StopApp(ctx context.Context, req *AppOperateRequest) ([]TaskResponse, error) {
	cp := AppOperateRequest{}
	if req != nil {
		cp = *req
	}
	cp.OperateType = "stop"
	return c.OperateApp(ctx, &cp)
}

type ExecuteScriptRequest struct {
	PadCodes      []string `json:"padCodes,omitempty"`
	ScriptContent string   `json:"scriptContent,omitempty"`
}

func (c *Client) ExecuteScript(ctx context.Context, req *ExecuteScriptRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathExecuteScript, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetTaskResultRequest struct {
	TaskIDs []FlexibleString `json:"taskIds,omitempty"`
}

func (c *Client) GetTaskResult(ctx context.Context, req *GetTaskResultRequest) ([]TaskResult, error) {
	var resp []TaskResult
	if err := c.Do(ctx, pathTaskInfo, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type ListTaskResultsRequest struct {
	Page     int              `json:"page,omitempty"`
	Rows     int              `json:"rows,omitempty"`
	TaskIDs  []FlexibleString `json:"taskIds,omitempty"`
	TaskType string           `json:"taskType,omitempty"`
}

func (c *Client) ListTaskResults(ctx context.Context, req *ListTaskResultsRequest) (*Page[TaskResult], error) {
	var resp Page[TaskResult]
	if err := c.Do(ctx, pathTaskPage, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

type TaskType struct {
	Type string    `json:"type,omitempty"`
	Name string    `json:"name,omitempty"`
	Raw  RawObject `json:"-"`
}

func (t *TaskType) UnmarshalJSON(data []byte) error {
	type alias TaskType
	var a alias
	if err := unmarshalRaw(data, &a.Raw, &a); err != nil {
		return err
	}
	*t = TaskType(a)
	return nil
}

func (c *Client) ListTaskTypes(ctx context.Context, req any) ([]TaskType, error) {
	var resp []TaskType
	if err := c.Do(ctx, pathTaskTypeList, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) CommandReboot(ctx context.Context, req *InstanceActionRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathCommandReboot, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) CommandReset(ctx context.Context, req *InstanceActionRequest) ([]TaskResponse, error) {
	var resp []TaskResponse
	if err := c.Do(ctx, pathCommandReset, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
