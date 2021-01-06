package model

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Status  int         `json:"status,omitempty"`
}
