module golang_rabbitmq_telegram

replace github.com/golang_rabbitmq_telegram => ./

go 1.24.6

require (
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/joho/godotenv v1.5.1
)
