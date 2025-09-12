package main

import (
	"golang_rabbitmq_telegram/internal/config"
	"golang_rabbitmq_telegram/internal/processor"
	"golang_rabbitmq_telegram/internal/rabbitmq"
	"golang_rabbitmq_telegram/internal/telegram"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize Telegram client
	tgClient, err := telegram.NewClient(cfg.TelegramToken, cfg.TelegramChatID, cfg.DebugMode)
	if err != nil {
		log.Fatalf("Telegram client initialization failed: %v", err)
	}

	// Test Telegram connection
	err = tgClient.TestConnection()
	if err != nil {
		log.Fatalf("Telegram connection test failed: %v", err)
	}

	// Initialize RabbitMQ client
	rmqClient, err := rabbitmq.NewClient(cfg.RabbitMQURL, cfg.ExchangeName, cfg.QueueName)
	if err != nil {
		log.Fatalf("RabbitMQ client initialization failed: %v", err)
	}
	defer rmqClient.Close()

	// Start consuming messages
	msgs, err := rmqClient.Consume()
	if err != nil {
		log.Fatalf("Failed to start consuming messages: %v", err)
	}

	log.Printf("✅ Consumer is running. Waiting for messages. To exit press CTRL+C")

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Process messages
	for {
		select {
		case d := <-msgs:
			if cfg.DebugMode {
				log.Printf("Received raw message from RabbitMQ: %s", string(d.Body))
			}

			// Process message
			messageText, err := processor.ProcessMessage(d.Body, cfg.Timezone)
			if err != nil {
				log.Printf("Error processing message: %v", err)
				continue
			}

			// Send to Telegram
			err = tgClient.SendMessage(messageText, 3)
			if err != nil {
				log.Printf("Failed to send message to Telegram: %v", err)
			}

		case <-sigChan:
			log.Println("Shutting down gracefully...")
			return
		}
	}
}
