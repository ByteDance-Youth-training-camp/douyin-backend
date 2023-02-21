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

const commentQueueName = "comment"

func generateMessageComment(userId int64, videoId int64, comment string) string {
	return fmt.Sprintf("%d,%d,%s", userId, videoId, comment)
}

func extractMessageComment(message string) (userId int64, videoId int64, comment string) {
	fmt.Sscanf(message, "%d,%d,%s", &userId, &videoId, &comment)
	return
}

func PublishCrateComment(userId int64, videoId int64, comment string) {
	_, err := commentChannel.QueueDeclare(
		commentQueueName,
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

	err = commentChannel.PublishWithContext(ctx, "", commentQueueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(generateMessageComment(userId, videoId, comment)),
	})

	if err != nil {
		panic(err)
	}
}

func PublishUncomment(commentId int64) {
	_, err := commentChannel.QueueDeclare(
		commentQueueName,
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

	err = commentChannel.PublishWithContext(ctx, "", commentQueueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(generateMessageComment(commentId, -1, "")),
	})

	if err != nil {
		panic(err)
	}
}

func ConsumerComment() {
	_, err := commentChannel.QueueDeclare(commentQueueName, false, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	msg, err := commentChannel.Consume(
		commentQueueName,
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
			userId, videoId, comment := extractMessageComment(msg)
			if videoId != -1 {
				comment := model.Comment{
					UserId:  userId,
					VideoId: videoId,
					Content: comment,
				}
				_, err := mysql.CreateComment(&comment)
				if err != nil {
					log.Println(err)
				}
			} else {
				err := mysql.DeleteCommentByID(userId)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	<-forever
}
