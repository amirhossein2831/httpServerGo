package Middleware

import (
	"bytes"
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/DB"
	"github.com/amirhossein2831/httpServerGo/src/Logger"
	"github.com/amirhossein2831/httpServerGo/src/model"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		url := r.RequestURI
		method := r.Method
		pro := r.Proto
		header := r.Header
		bodyBytes, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		Logger.GetInstance().GetLogger().Info("Incoming Request: ",
			zap.String("Host: ", host),
			zap.String("URL: ", url),
			zap.String("Method: ", method),
			zap.String("Protocol: ", pro),
			zap.Any("Header: ", header),
			zap.Any("Body: ", string(bodyBytes)),
			zap.Time("Timestamp", time.Now()),
		)

		headerJSON, _ := json.Marshal(header)

		DB.GetInstance().GetDb().Create(&model.Request{
			Host:     host,
			URL:      url,
			Method:   method,
			Protocol: pro,
			Header:   headerJSON,
			Body:     bodyBytes,
		})

		next.ServeHTTP(w, r)
	})
}
