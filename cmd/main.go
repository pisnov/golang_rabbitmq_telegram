package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang_rabbitmq_telegram/internal/config"
	"golang_rabbitmq_telegram/internal/processor"
	"golang_rabbitmq_telegram/internal/rabbitmq"
	"golang_rabbitmq_telegram/internal/telegram"
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

			// Send to Telegram dengan retry mechanism yang baru
			err = tgClient.SendMessageWithRetry(messageText, 3)
			if err != nil {
				log.Printf("Failed to send message to Telegram after retries: %v", err)
			}

			// Tambahkan delay untuk menghindari rate limiting
			time.Sleep(1 * time.Second)

		case <-sigChan:
			log.Println("Shutting down gracefully...")
			return
		}
	}
}
