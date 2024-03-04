package main

import (
	"background-order-service/domain"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":         "localhost:8083",
		"group.id":                  "background-order-service",
		"auto.offset.reset":         "latest",
		"enable.auto.commit":        false,
		"receive.message.max.bytes": "31457280",
		"fetch.max.bytes":           "31451280",
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	err = c.SubscribeTopics([]string{"microshop.orders"}, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to subscribe to topics: %s\n", err)
		os.Exit(1)
	}

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				var msg domain.Message
				if err := json.Unmarshal(e.Value, &msg); err != nil {
					fmt.Printf("Message on %s cannot be decoded: %v\n", e.TopicPartition, err)
					continue
				}

				fmt.Printf("Message on %s: %s\n", e.TopicPartition, msg.Payload)
				_, err := c.CommitMessage(e)
				if err != nil {
					fmt.Printf("Failed to commit message: %v\n", err)
				}
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Println("Closing consumer")
	c.Close()
}
