package rabbitmq

import (
	"douyin_backend/biz/config"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	amqp "github.com/rabbitmq/amqp091-go"
)

var Rmq *amqp.Connection
var commentChannel *amqp.Channel
var favoriteChannel *amqp.Channel
var followChannel *amqp.Channel

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
	commentChannel = channel

	channel, err = dial.Channel()
	if err != nil {
		panic(err)
	}
	favoriteChannel = channel

	channel, err = dial.Channel()
	if err != nil {
		panic(err)
	}
	followChannel = channel

	go ConsumerComment()
	go ConsumerFavorite()
	go ConsumerFollow()
}

func Close() {
	if err := commentChannel.Close(); err != nil {
		hlog.Error("close comment rabbitmq channel channel error")
	}
	if err := favoriteChannel.Close(); err != nil {
		hlog.Error("close favorite rabbitmq channel error")
	}
	if err := followChannel.Close(); err != nil {
		hlog.Error("close follow rabbitmq channel error")
	}

	if err := Rmq.Close(); err != nil {
		hlog.Error("close rabbitmq connection error")
	}
}
