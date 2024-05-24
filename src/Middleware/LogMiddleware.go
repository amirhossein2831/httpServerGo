package Middleware

import (
	"bytes"
	"github.com/amirhossein2831/httpServerGo/src/Logger"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bodyBytes, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		body := string(bodyBytes)

		Logger.GetInstance().GetLogger().Info("Incoming Request: ",
			zap.String("Host: ", r.Host),
			zap.String("URL: ", r.RequestURI),
			zap.String("Method: ", r.Method),
			zap.String("Protocol: ", r.Proto),
			zap.Any("Header: ", r.Header),
			zap.Any("Body: ", body),
			zap.Time("Timestamp", time.Now()),
		)

		next.ServeHTTP(w, r)
	})
}
