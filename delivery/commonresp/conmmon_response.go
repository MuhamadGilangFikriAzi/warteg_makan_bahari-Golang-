package commonresp

import "encoding/json"

type AppHttpResponse interface {
	SendData(message ResponMessage)
	SendError(message ErrorMessage)
}

type ResponMessage struct {
	Status       string      `json:"status"`
	ResponseCode string      `json:"response_code"`
	Description  string      `json:"description"`
	Data         interface{} `json:"data"`
}

type ErrorMessage struct {
	HttpCode int `json:"http_code"`
	ErrorDescription
}

type ErrorDescription struct {
	Status       string `json:"status"`
	Service      string `json:"service"`
	ResponseCode string `json:"response_code"`
	Description  string `json:"description"`
}

func (e ErrorMessage) ToJson() string {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}

func (e ResponMessage) ToJson() string {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}

func NewResponMessage(responCode string, description string, data interface{}) ResponMessage {
	return ResponMessage{
		"Success",
		responCode,
		description,
		data,
	}
}

func NewErrorMessage(httpCode int, service string, resCode string, description string) ErrorMessage {
	return ErrorMessage{
		httpCode, ErrorDescription{
			"Error",
			service,
			resCode,
			description,
		},
	}
}
