package main

import (
	"background-order-service/config"
	"background-order-service/domain"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	cfg := config.ReadConfig()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":         cfg.Servers,
		"broker.address.family":     "v4",
		"group.id":                  cfg.Group,
		"session.timeout.ms":        6000,
		"auto.offset.reset":         cfg.Offset,
		"enable.auto.offset.store":  false,
		"receive.message.max.bytes": "1313486160",
		"security.protocol":         cfg.Security,
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	err = c.SubscribeTopics([]string{"microshop.order-outbox"}, nil)
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
			ev := c.Poll(cfg.Poll)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:

				// Create an instance of Outbox to store the unmarshaled data
				var outbox domain.Outbox

				// Unmarshal the JSON message into the Outbox struct
				err := json.Unmarshal(e.Value, &outbox)
				if err != nil {
					fmt.Println("Error decoding message:", err)
					return
				}

				fmt.Println("Payload:", outbox.Payload)

				_, err = c.CommitMessage(e)
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
