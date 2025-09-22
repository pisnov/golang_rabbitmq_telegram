package telegram

import (
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Client struct {
	Bot    *tgbotapi.BotAPI
	ChatID int64
}

func NewClient(token string, chatID int64, debug bool) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create Telegram bot: %v", err)
	}
	bot.Debug = debug

	if debug {
		log.Printf("Authorized on account %s", bot.Self.UserName)
	}

	return &Client{Bot: bot, ChatID: chatID}, nil
}

// extractRetryAfterTime mengekstrak waktu retry dari error message Telegram
func extractRetryAfterTime(errorMsg string) int {
	// Error message format: "Too Many Requests: retry after 15"
	if strings.Contains(errorMsg, "retry after") {
		parts := strings.Split(errorMsg, "retry after")
		if len(parts) > 1 {
			var retryAfter int
			_, err := fmt.Sscanf(strings.TrimSpace(parts[1]), "%d", &retryAfter)
			if err == nil {
				return retryAfter
			}
		}
	}
	return 0
}

// SendMessageWithRetry mengirim pesan dengan retry mechanism yang menangani rate limiting
func (c *Client) SendMessageWithRetry(message string, maxRetries int) error {
	for i := 0; i < maxRetries; i++ {
		msg := tgbotapi.NewMessage(c.ChatID, message)
		msg.ParseMode = "HTML"

		_, err := c.Bot.Send(msg)

		if err != nil {
			// Handle rate limiting error (429)
			if strings.Contains(err.Error(), "Too Many Requests") {
				retryAfter := extractRetryAfterTime(err.Error())
				if retryAfter > 0 {
					log.Printf("Rate limited by Telegram. Waiting %d seconds before retry...", retryAfter)
					time.Sleep(time.Duration(retryAfter) * time.Second)
					continue // Retry after waiting
				}
			}

			if strings.Contains(err.Error(), "chat not found") {
				return fmt.Errorf("chat not found. Please check: 1) Bot is admin in channel, 2) Correct chat ID, 3) User started chat with bot")
			}

			log.Printf("Retry %d/%d failed: %v", i+1, maxRetries, err)

			// Exponential backoff for other errors
			backoffTime := time.Duration(i+1) * 2 * time.Second
			time.Sleep(backoffTime)
			continue
		}

		log.Printf("✅ Message sent to Telegram successfully")
		return nil
	}

	return fmt.Errorf("failed to send message after %d attempts", maxRetries)
}

func (c *Client) TestConnection() error {
	testMsg := tgbotapi.NewMessage(c.ChatID, "🤖 Bot connected successfully! Ready to receive messages from RabbitMQ.")
	_, err := c.Bot.Send(testMsg)
	return err
}
