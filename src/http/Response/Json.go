package Response

import (
	"encoding/json"
	"net/http"
)

type JsonResponse interface {
	SetStatusCode(code int) JsonResponse
	SetSuccess(isSuccess bool) JsonResponse
	SetData(data interface{}) JsonResponse
	Send(w http.ResponseWriter)
}

type Json struct {
	StatusCode int         `json:"status-code"`
	IsSuccess  bool        `json:"is-success"`
	Data       interface{} `json:"data"`
}

func NewJson() JsonResponse {
	return &Json{
		StatusCode: http.StatusBadRequest,
		IsSuccess:  false,
		Data:       nil,
	}
}

func (j *Json) SetStatusCode(code int) JsonResponse {
	j.StatusCode = code
	return j
}

func (j *Json) SetSuccess(isSuccess bool) JsonResponse {
	j.IsSuccess = isSuccess
	return j
}

func (j *Json) SetData(data interface{}) JsonResponse {
	j.Data = data
	return j
}

func (j *Json) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(j.StatusCode)
	err := json.NewEncoder(w).Encode(j)
	if err != nil {
		NewJson().SetSuccess(false).SetStatusCode(http.StatusBadRequest).SetData(err).Send(w)
	}
}
