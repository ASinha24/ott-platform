package api

import "fmt"

type OTTError struct {
	Code        ErrorCode `json:"code,omitempty"`
	Message     string    `json:"message,omitempty"`
	Description string    `json:"description,omitempty"`
}

func (e OTTError) Error() string {
	return fmt.Sprintf("code: %d msg: %s, err: %s", e.Code, e.Message, e.Description)
}
