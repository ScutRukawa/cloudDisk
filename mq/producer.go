package mq

import (
	"cloudDisk/config"
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel

var notifyClose chan *amqp.Error

func init() {
	// 是否开启异步转移功能，开启时才初始化rabbitMQ连接
	log.Println("xxxxxxxxxxxxx")

	if !config.AsyncTransferEnable {
		return
	}
	if initChannel() {
		channel.NotifyClose(notifyClose)
	}
	// 断线自动重连
	go func() {
		for {
			select {
			case msg := <-notifyClose:
				conn = nil
				channel = nil
				log.Printf("onNotifyChannelClosed: %+v\n", msg)
				initChannel()
			}
		}
	}()
}
func initChannel() bool {
	if channel != nil {
		return true
	}

	conn, err := amqp.Dial(config.RabbitURL)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	channel, err = conn.Channel()
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

func Publish(exchange, routingKey string, msg []byte) bool {
	if !initChannel() {
		return false
	}

	err := channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
	if err != nil {
		log.Println(err.Error())
	}
	return true
}
