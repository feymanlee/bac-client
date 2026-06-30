package bac

import (
	"errors"
	"fmt"
)

var (
	ErrMissingAppKey    = errors.New("bac: app key is required")
	ErrMissingAppSecret = errors.New("bac: app secret is required")
	ErrInvalidDESKey    = errors.New("bac: invalid des key")
)

type APIError struct {
	Code      int
	Message   string
	Timestamp FlexibleString
	RawBody   []byte
}

func (e *APIError) Error() string {
	return fmt.Sprintf("bac: api error code=%d message=%q ts=%s", e.Code, e.Message, e.Timestamp.String())
}

type HTTPError struct {
	StatusCode int
	Status     string
	RawBody    []byte
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("bac: http error status=%s", e.Status)
}

type DecodeError struct {
	Err     error
	RawBody []byte
}

func (e *DecodeError) Error() string {
	return fmt.Sprintf("bac: decode response: %v", e.Err)
}

func (e *DecodeError) Unwrap() error {
	return e.Err
}
