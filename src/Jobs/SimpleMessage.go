package Jobs

import (
	"context"
	"github.com/amirhossein2831/httpServerGo/src/Logger"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
	"time"
)

type SimpleMessage struct {
}

func NewSimpleMessageJob() Job {
	return &SimpleMessage{}
}

func (it *SimpleMessage) Publish() {
	ch, q, err := channelWithDefaultQueue("simpleMessage")
	defer ch.Close()
	if err != nil {
		Logger.GetInstance().GetLogger().Error("Failed to open rabbit channel or queue", zap.Error(err), zap.Time("timestamp", time.Now()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "Hello World!"
	err = ch.PublishWithContext(ctx, "", q.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		Logger.GetInstance().GetLogger().Error("Failed to publish a message", zap.Error(err), zap.Time("timestamp", time.Now()))
		return
	}
	Logger.GetInstance().GetLogger().Info("Published a message: ", zap.String("body", body), zap.Time("timestamp", time.Now()))
}

func (it *SimpleMessage) Consume() {
	ch, q, err := channelWithDefaultQueue("simpleMessage")
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
			Logger.GetInstance().GetLogger().Info("Received a message", zap.String("queue", q.Name), zap.Any("data", d.Body), zap.Time("timestamp", time.Now()))
			log.Printf("Received a message from %s Queue: %s", q.Name, d.Body)
		}
	}()

	Logger.GetInstance().GetLogger().Info("Waiting for Publisher", zap.String("queue", q.Name), zap.Time("timestamp", time.Now()))
	log.Printf(" [****] Waiting for messages form %v queue. To exit press CTRL+C", q.Name)
	<-forever
}
