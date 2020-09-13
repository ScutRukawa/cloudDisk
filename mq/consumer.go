package mq

import (
	"log"
)

func StartConsume(qName, cName string, callback func(msg []byte) bool) {
	log.Println("kkkkk:")

	msgs, err := channel.Consume(
		qName,
		cName,
		true,
		false,
		false,
		false,
		nil,
	)
	log.Println("kkkkk")
	if err != nil {
		log.Println(err.Error())
		return
	}

	done := make(chan bool)
	go func() {
		for msg := range msgs {
			processSuc := callback(msg.Body)
			if !processSuc {
				//to do : 将任务写到另一个队列，用于异常情况的重试
			}
		}

	}()

	<-done

	channel.Close() //为何要关闭 ， 消费者 生产者都只用了这一个信道? to do
}
