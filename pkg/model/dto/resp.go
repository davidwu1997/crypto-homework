package dto

type Resp struct {
	Data    interface{} `json:"data,omitempty"`
	ErrMsg  string      `json:"errMsg,omitempty"`
	ErrCode string      `json:"errCode,omitempty"`
}
