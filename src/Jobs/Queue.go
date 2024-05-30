package Jobs

import (
	"github.com/amirhossein2831/httpServerGo/src/Logger"
	"github.com/amirhossein2831/httpServerGo/src/rabbit"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"time"
)

func channelWithDefaultQueue(qName string) (*amqp.Channel, amqp.Queue, error) {
	ch, err := rabbit.GetChannel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	q, err := defaultQueue(ch, qName)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	return ch, q, nil
}

func channelWithCustomQueue(qName string, durable, delete, exclusive, nowWit bool, args amqp.Table) (*amqp.Channel, amqp.Queue, error) {
	ch, err := rabbit.GetChannel()
	defer ch.Close()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	q, err := customQueue(ch, qName, durable, delete, exclusive, nowWit, args)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	return ch, q, nil
}

func defaultQueue(ch *amqp.Channel, name string) (amqp.Queue, error) {
	qTemp, err := ch.QueueDeclare(name, false, false, false, false, nil)
	if err != nil {
		Logger.GetInstance().GetLogger().Error("Failed to declare queue", zap.Error(err), zap.Time("timestamp", time.Now()))
		return amqp.Queue{}, err
	}

	return qTemp, nil
}

func customQueue(ch *amqp.Channel, name string, durable, delete, exclusive, nowWit bool, args amqp.Table) (amqp.Queue, error) {

	qTemp, err := ch.QueueDeclare(name, durable, delete, exclusive, nowWit, args)
	if err != nil {
		Logger.GetInstance().GetLogger().Error("Failed to declare queue", zap.Error(err), zap.Time("timestamp", time.Now()))
		return amqp.Queue{}, err
	}

	return qTemp, nil
}
