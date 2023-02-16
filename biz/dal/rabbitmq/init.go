package rabbitmq

import (
	"douyin_backend/biz/config"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	amqp "github.com/rabbitmq/amqp091-go"
)

var Rmq *amqp.Connection
var Db *amqp.Channel

func Init() {
	dial, err := amqp.Dial(config.Cfg.RabbitMQ.Address)
	if err != nil {
		panic(err)
	}
	Rmq = dial

	channel, err := dial.Channel()
	if err != nil {
		panic(err)
	}
	Db = channel
}

func Close() {
	if err := Db.Close(); err != nil {
		hlog.Error("close rabbitmq channel error")
	}

	if err := Rmq.Close(); err != nil {
		hlog.Error("close rabbitmq connection error")
	}
}
