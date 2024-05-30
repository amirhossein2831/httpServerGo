package Jobs

import (
	"context"
	"encoding/json"
	"github.com/amirhossein2831/httpServerGo/src/Logger"
	"github.com/amirhossein2831/httpServerGo/src/http/Response"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
	"time"
)

type Logging struct {
	json Response.Json
}

func NewLoggingJob(json Response.Json) Job {
	return &Logging{
		json: json,
	}
}

func (it *Logging) Publish() {
	ch, q, err := channelWithDefaultQueue("Logging")
	defer ch.Close()
	if err != nil {
		Logger.GetInstance().GetLogger().Error("Failed to open rabbit channel or queue", zap.Error(err), zap.Time("timestamp", time.Now()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(it.json)
	if err != nil {
		Logger.GetInstance().GetLogger().Error("Failed to marshal JSON", zap.Error(err), zap.Time("timestamp", time.Now()))
		return
	}

	err = ch.PublishWithContext(ctx, "", q.Name, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		Logger.GetInstance().GetLogger().Error("Failed to publish a message", zap.Error(err), zap.Time("timestamp", time.Now()))
		return
	}
	Logger.GetInstance().GetLogger().Info("Published a ResponseLog: ", zap.Any("body", it.json), zap.Time("timestamp", time.Now()))
}

func (it *Logging) Consume() {
	ch, q, err := channelWithDefaultQueue("Logging")
	defer ch.Close()
	if err != nil {
		Logger.GetInstance().GetLogger().Error("Failed to open rabbit channel or queue", zap.Error(err), zap.Time("timestamp", time.Now()))
		return
	}

	message, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		Logger.GetInstance().GetLogger().Error("Failed to register a consumer", zap.Error(err), zap.Time("timestamp", time.Now()))
		return
	}

	var forever chan struct{}
	go func() {
		for d := range message {
			var j Response.Json
			err := json.Unmarshal(d.Body, &j)
			if err != nil {
				Logger.GetInstance().GetLogger().Error("Failed to unmarshal JSON", zap.Error(err), zap.Time("timestamp", time.Now()))
				continue
			}
			if j.IsSuccess {
				Logger.GetInstance().GetLogger().Info("Response sent",
					zap.Int("StatusCode", j.StatusCode),
					zap.Bool("IsSuccess", j.IsSuccess),
					zap.Any("Data", j.Data),
					zap.Time("Timestamp", time.Now()),
				)
			} else {
				Logger.GetInstance().GetLogger().Error("Response sent with error",
					zap.Int("StatusCode", j.StatusCode),
					zap.Bool("IsSuccess", j.IsSuccess),
					zap.Error(j.Data.(error)),
					zap.Time("Timestamp", time.Now()),
				)
			}

			Logger.GetInstance().GetLogger().Info("Received a message", zap.String("queue", q.Name), zap.Any("data", d.Body), zap.Time("timestamp", time.Now()))
			log.Printf("Received a message from %v Queue: %v", q.Name, d.Body)
		}
	}()

	Logger.GetInstance().GetLogger().Info("Waiting for Publisher", zap.String("queue", q.Name), zap.Time("timestamp", time.Now()))
	log.Printf(" [****] Waiting for messages form %v queue. To exit press CTRL+C", q.Name)
	<-forever
}
