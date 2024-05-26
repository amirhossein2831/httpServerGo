package Response

import (
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/Logger"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type JsonResponse interface {
	GetStatusCode() int
	GetSuccess() bool
	GetData() interface{}
	SetStatusCode(code int) JsonResponse
	SetSuccess(isSuccess bool) JsonResponse
	SetData(data interface{}) JsonResponse
	Log() JsonResponse
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

func (j *Json) GetStatusCode() int {
	return j.StatusCode
}

func (j *Json) GetSuccess() bool {
	return j.IsSuccess
}

func (j *Json) GetData() interface{} {
	return j.Data
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

func (j *Json) Log() JsonResponse {
	logger := Logger.GetInstance().GetLogger()

	// Log the Response
	if j.IsSuccess {
		logger.Info("Response sent",
			zap.Int("StatusCode", j.StatusCode),
			zap.Bool("IsSuccess", j.IsSuccess),
			zap.Any("Data", j.Data),
			zap.Time("Timestamp", time.Now()),
		)
	} else {
		logger.Error("Response sent with error",
			zap.Int("StatusCode", j.StatusCode),
			zap.Bool("IsSuccess", j.IsSuccess),
			zap.Error(j.Data.(error)),
			zap.Time("Timestamp", time.Now()),
		)
	}

	// save the Response in DB
	data, _ := json.Marshal(j.Data)
	DB.GetInstance().GetDb().Create(&model.Response{
		Status:    j.StatusCode,
		IsSuccess: j.IsSuccess,
		Data:      data,
	})

	return j
}

func (j *Json) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(j.StatusCode)
	if !j.IsSuccess {
		j.Data = j.Data.(error).Error()
	}
	err := json.NewEncoder(w).Encode(j)
	if err != nil {
		NewJson().SetSuccess(false).SetStatusCode(http.StatusBadRequest).SetData(err).Send(w)
	}
}
