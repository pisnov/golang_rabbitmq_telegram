abbitMQ to Telegram Bot
Sebuah aplikasi Go yang mengonsumsi pesan dari RabbitMQ dan mengirimkannya ke channel Telegram dengan format yang rapi.

📋 Deskripsi
Aplikasi ini bertindak sebagai konsumen dari queue RabbitMQ yang menerima pesan dalam format JSON, memprosesnya, dan mengirimkannya ke channel Telegram. Aplikasi ini sangat berguna untuk notifikasi real-time dari sistem Anda ke Telegram.

✨ Fitur
✅ Konsumsi pesan dari RabbitMQ

✅ Kirim notifikasi ke Telegram

✅ Format pesan yang rapi dan informatif

✅ Konversi waktu UTC ke zona waktu lokal

✅ Konfigurasi melalui environment variables

✅ Mode debug untuk troubleshooting

✅ Retry mechanism untuk pengiriman pesan

✅ Penanganan error yang robust

🛠️ Teknologi
Golang - Bahasa pemrograman utama

RabbitMQ - Message broker

Telegram Bot API - Untuk mengirim pesan ke Telegram

Godotenv - Untuk mengelola environment variables

📦 Prasyarat
Go 1.18 atau lebih baru

Server RabbitMQ

Akun Telegram dan Bot (dibuat via BotFather)

🚀 Instalasi
Clone repository:

bash
git clone https://github.com/username/rabbitmq-telegram-bot.git
cd rabbitmq-telegram-bot
Install dependencies:

bash
go mod download
Setup environment variables:

bash
cp .env.example .env
# Edit file .env dengan konfigurasi Anda
⚙️ Konfigurasi
Edit file .env dengan nilai yang sesuai:

env
# Telegram Bot Configuration
TELEGRAM_BOT_TOKEN=your_telegram_bot_token_here
TELEGRAM_CHAT_ID=your_chat_id_here

# RabbitMQ Configuration
RABBITMQ_URL=amqp://admin:password@rabbitmq-server:5672/
RABBITMQ_EXCHANGE=your_exchange_name
RABBITMQ_QUEUE=your_queue_name

# Timezone Configuration (contoh: GMT+8, Asia/Jakarta)
TIMEZONE=GMT+8

# Debug Mode
DEBUG_MODE=false
Cara Mendapatkan Telegram Bot Token
Buka BotFather di Telegram

Kirim perintah /newbot dan ikuti instruksinya

Salin token yang diberikan

Cara Mendapatkan Chat ID
Tambahkan bot Anda ke channel/group

Kirim pesan apa saja ke channel/group

Akses URL berikut di browser:

text
https://api.telegram.org/bot<YOUR_BOT_TOKEN>/getUpdates
Cari chat.id dalam respons JSON

🎯 Cara Penggunaan
Jalankan aplikasi:

bash
go run main.go
Aplikasi akan:

Terhubung ke RabbitMQ

Mengonsumsi pesan dari queue yang ditentukan

Mengirim pesan ke Telegram

Menampilkan log aktivitas

📝 Format Pesan
Pesan dari RabbitMQ harus dalam format JSON dengan struktur berikut:

json
{
  "id": 123456,
  "userID": 789,
  "RuleName": "ignition",
  "UserName": "LV_8014",
  "Namespace": "default",
  "Value": "Notifikasi ignition Unit PPA-BIB\r\n\r\nUnit : LV_8014\r\nWaktu Kejadian : 2025-09-10T05:58:19",
  "State": "New",
  "lng": 115.6061116,
  "lat": -3.5127966,
  "alt": 89.0,
  "utc": 1757483899000
}
Aplikasi akan memformat pesan tersebut menjadi:

text
<b>Notifikasi ignition Unit PPA-BIB</b>

Unit: LV_8014
Waktu Kejadian: 2025-09-10 13:58:19 GMT+8
🏗️ Struktur Project
text
rabbitmq-telegram-bot/
├── main.go          # Kode utama aplikasi
├── go.mod           # File dependensi Go
├── go.sum           # Checksum dependensi
├── .env.example     # Template environment variables
└── README.md        # Dokumentasi
🔧 Troubleshooting
Error "chat not found"
Pastikan bot sudah ditambahkan sebagai admin di channel

Pastikan chat ID benar

Pastikan user sudah memulai chat dengan bot

Error koneksi RabbitMQ
Periksa kredensial dan URL RabbitMQ

Pastikan server RabbitMQ dapat diakses

Pesan tidak terkirim
Aktifkan debug mode untuk melihat log detail

Periksa token bot Telegram

🤝 Kontribusi
Kontribusi selalu diterima! Silakan:

Fork project ini

Buat branch untuk fitur Anda (git checkout -b feature/AmazingFeature)

Commit perubahan Anda (git commit -m 'Add some AmazingFeature')

Push ke branch (git push origin feature/AmazingFeature)

Buat Pull Request

📜 Lisensi
Distributed under the MIT License. Lihat file LICENSE untuk informasi lebih lanjut.

💬 Dukungan
Jika Anda memiliki pertanyaan atau masalah, silakan buat issue di GitHub atau hubungi melalui:

Email: nofisetiawan88@gmail.com

Telegram: @pisnov

🙏 Penghargaan
Go Telegram Bot API

RabbitMQ Go Client

Godotenv

⭐ Jangan lupa memberikan bintang di repository ini jika Anda merasa terbantu!