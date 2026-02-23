package models

type ErrorMessage struct {
	Error   string   `json:"error,omitempty"`
	Details []string `json:"details,omitempty"`
	Code    string   `json:"code,omitempty"`
}
