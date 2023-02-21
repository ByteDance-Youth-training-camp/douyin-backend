package rabbitmq

import (
	"context"
	"douyin_backend/biz/dal/mysql"
	"douyin_backend/biz/model"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

const followQueueName = "follow"

func generateMessageFollow(create bool, userId int64, followerId int64) string {
	return fmt.Sprintf("%d,%d,%t", userId, followerId, create)
}

func extractMessageFollow(message string) (userId int64, followerId int64, create bool) {
	fmt.Sscanf(message, "%d,%d,%t", &userId, &followerId, &create)
	return
}

func PublishCrateFollow(userId int64, followerId int64) {
	_, err := followChannel.QueueDeclare(
		followQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = followChannel.PublishWithContext(ctx, "", followQueueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(generateMessageFollow(true, userId, followerId)),
	})

	if err != nil {
		panic(err)
	}
}

func PublishUnfollow(userId int64, followerId int64) {
	_, err := followChannel.QueueDeclare(
		followQueueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = followChannel.PublishWithContext(ctx, "", followQueueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(generateMessageFollow(false, userId, followerId)),
	})

	if err != nil {
		panic(err)
	}
}

func ConsumerFollow() {
	_, err := followChannel.QueueDeclare(followQueueName, false, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	msg, err := followChannel.Consume(
		followQueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msg {
			msg := string(d.Body)
			userId, followerId, create := extractMessageFollow(msg)
			if create {
				follow := model.Follow{
					UserId:     userId,
					FollowerId: followerId,
					Canceled:   false,
				}
				err := mysql.Follow(&follow)
				if err != nil {
					log.Println(err)
				}
			} else {
				err := mysql.Unfollow(userId, followerId)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	<-forever
}
