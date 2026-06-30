package bac

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type Response[T any] struct {
	Code      int            `json:"code"`
	Message   string         `json:"msg"`
	Data      T              `json:"data"`
	Timestamp FlexibleString `json:"ts"`
}

type FlexibleString string

func (s *FlexibleString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*s = ""
		return nil
	}
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		*s = FlexibleString(str)
		return nil
	}
	var num json.Number
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	if err := dec.Decode(&num); err != nil {
		return err
	}
	*s = FlexibleString(num.String())
	return nil
}

func (s FlexibleString) String() string {
	return string(s)
}

func (s FlexibleString) Int64() (int64, error) {
	return strconv.ParseInt(string(s), 10, 64)
}

type PageRequest struct {
	Page int `json:"page,omitempty"`
	Size int `json:"size,omitempty"`
}

type Page[T any] struct {
	Records   []T             `json:"records,omitempty"`
	List      []T             `json:"list,omitempty"`
	PageData  []T             `json:"pageData,omitempty"`
	Total     FlexibleString  `json:"total,omitempty"`
	TotalPage FlexibleString  `json:"totalPage,omitempty"`
	Page      FlexibleString  `json:"page,omitempty"`
	Size      FlexibleString  `json:"size,omitempty"`
	Rows      FlexibleString  `json:"rows,omitempty"`
	Raw       json.RawMessage `json:"-"`
}

func (p *Page[T]) UnmarshalJSON(data []byte) error {
	type alias Page[T]
	var a alias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	*p = Page[T](a)
	p.Raw = append(p.Raw[:0], data...)
	return nil
}

type TaskResult struct {
	TaskID FlexibleString  `json:"taskId,omitempty"`
	Status FlexibleString  `json:"status,omitempty"`
	Code   FlexibleString  `json:"code,omitempty"`
	Msg    string          `json:"msg,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`
	Raw    json.RawMessage `json:"-"`
}

func (r *TaskResult) UnmarshalJSON(data []byte) error {
	type alias TaskResult
	var a alias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	*r = TaskResult(a)
	r.Raw = append(r.Raw[:0], data...)
	return nil
}

type RawObject map[string]json.RawMessage

func unmarshalRaw(data []byte, raw *RawObject, dst any) error {
	if err := json.Unmarshal(data, dst); err != nil {
		return err
	}
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err == nil {
		*raw = obj
	}
	return nil
}
