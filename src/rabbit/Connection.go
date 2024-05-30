package rabbit

import (
	"github.com/amirhossein2831/httpServerGo/src/Logger"
	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"log"
	"time"

	"sync"
)

var (
	conInstance Connection
	once        sync.Once
)

type Connection interface {
	GetConnection() *amqp.Connection
	SetConnection(conn *amqp.Connection)
}

type RbConnection struct {
	conn *amqp091.Connection
}

func GetInstance() Connection {
	once.Do(func() {
		conn, err := amqp.Dial("amqp://root:password@localhost:5672/")
		if err != nil {
			Logger.GetInstance().GetLogger().Error("Failed to connect to RabbitMQ", zap.Error(err), zap.Time("timestamp", time.Now()))
			log.Fatal(err)
		}
		Logger.GetInstance().GetLogger().Info("Connect to RabbitMQ Successfully", zap.Time("timestamp", time.Now()))
		conInstance = &RbConnection{
			conn: conn,
		}
	})
	return conInstance
}

func GetChannel() (*amqp.Channel, error) {
	conn := GetInstance()
	ch, err := conn.GetConnection().Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

func (c *RbConnection) GetConnection() *amqp.Connection {
	return c.conn
}

func (c *RbConnection) SetConnection(conn *amqp.Connection) {
	c.conn = conn
}
