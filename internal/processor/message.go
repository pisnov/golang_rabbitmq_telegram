package processor

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"golang_rabbitmq_telegram/pkg/utils" // Ganti dengan path yang benar
)

type MessageStructure struct {
	ID        int64   `json:"id"`
	UserID    int     `json:"userID"`
	RuleName  string  `json:"RuleName"`
	UserName  string  `json:"UserName"`
	Namespace string  `json:"Namespace"`
	Value     string  `json:"Value"`
	State     string  `json:"State"`
	Lng       float64 `json:"lng"`
	Lat       float64 `json:"lat"`
	Alt       float64 `json:"alt"`
	Utc       int64   `json:"utc"`
}

func ProcessMessage(body []byte, timezone string) (string, error) {
	var messageData MessageStructure
	err := json.Unmarshal(body, &messageData)
	if err != nil {
		log.Printf("Received non-JSON message: %s", string(body))
		return fmt.Sprintf("⚠️ Received non-JSON message:\n%s", string(body)), nil
	}

	return FormatMessage(messageData.Value, messageData.Utc, timezone), nil
}

func FormatMessage(value string, utc int64, timezoneStr string) string {
	// Pisahkan value menjadi bagian-bagian
	parts := strings.SplitN(value, "\r\n\r\n", 2)

	var title, content string
	if len(parts) >= 2 {
		title = parts[0]
		content = parts[1]
	} else {
		// Jika tidak ada pemisah, gunakan seluruh value sebagai konten
		title = "Notifikasi"
		content = value
	}

	// Bersihkan dan format konten
	content = strings.ReplaceAll(content, "\r\n", "\n")
	contentLines := strings.Split(content, "\n")

	// Cari dan ekstrak informasi unit
	var unit string
	for _, line := range contentLines {
		if strings.Contains(line, "Unit :") {
			unit = strings.TrimSpace(strings.TrimPrefix(line, "Unit :"))
			break
		}
	}

	// Format HTML untuk Telegram
	var htmlBuilder strings.Builder

	// Judul
	htmlBuilder.WriteString(fmt.Sprintf("<b>%s</b>\n\n", title))

	// Unit (jika ditemukan)
	if unit != "" {
		htmlBuilder.WriteString(fmt.Sprintf("Unit: %s\n", unit))
	}

	// Waktu kejadian (jika tersedia UTC)
	if utc != 0 {
		convertedTime := utils.ConvertUtcToLocal(utc, timezoneStr)
		htmlBuilder.WriteString(fmt.Sprintf("Waktu Kejadian: %s\n", convertedTime))
	}

	// Konten lainnya (jika ada)
	for _, line := range contentLines {
		if !strings.Contains(line, "Unit :") && !strings.Contains(line, "Waktu Kejadian :") && line != "" {
			htmlBuilder.WriteString(fmt.Sprintf("%s\n", line))
		}
	}

	return htmlBuilder.String()
}
