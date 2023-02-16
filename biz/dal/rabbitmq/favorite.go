package rabbitmq

import (
	"context"
	"douyin_backend/biz/dal/mysql"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

const favoriteQueueName = "favorite"

func generateMessageFavorite(create bool, userId int64, videoId int64) string {
	return fmt.Sprintf("%d,%d,%t", userId, videoId, create)
}

func extractMessageFavorite(message string) (userId int64, videoId int64, create bool) {
	fmt.Sscanf(message, "%d,%d,%t", &userId, &videoId, &create)
	return
}

func PublishCrateFavorite(userId int64, videoId int64) {
	_, err := Db.QueueDeclare(
		favoriteQueueName,
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

	err = Db.PublishWithContext(ctx, "", favoriteQueueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(generateMessageFavorite(true, userId, videoId)),
	})

	if err != nil {
		panic(err)
	}
}

func PublishDeleteFavorite(userId int64, videoId int64) {
	_, err := Db.QueueDeclare(
		favoriteQueueName,
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

	err = Db.PublishWithContext(ctx, "", favoriteQueueName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(generateMessageFavorite(false, userId, videoId)),
	})

	if err != nil {
		panic(err)
	}
}

func ConsumerFavorite() {
	_, err := Db.QueueDeclare(favoriteQueueName, false, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	msg, err := Db.Consume(
		favoriteQueueName,
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
			userId, videoId, create := extractMessageFavorite(msg)
			if create {
				err := mysql.AddFavorite(userId, videoId)
				if err != nil {
					log.Println(err)
				}
			} else {
				err := mysql.RemoveFavorite(userId, videoId)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	<-forever
}
