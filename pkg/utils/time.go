package utils

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func ConvertUtcToLocal(utcMillis int64, timezoneStr string) string {
	// Parse timestamp (dalam milidetik)
	utcTime := time.Unix(0, utcMillis*int64(time.Millisecond))

	// Tentukan lokasi berdasarkan timezone
	var loc *time.Location
	var err error

	// Coba parsing sebagai offset (e.g., GMT+8)
	if strings.HasPrefix(strings.ToUpper(timezoneStr), "GMT") {
		offsetStr := strings.TrimPrefix(strings.ToUpper(timezoneStr), "GMT")
		if offsetStr == "" {
			offsetStr = "+0"
		}

		// Parse offset
		var hours int
		_, err := fmt.Sscanf(offsetStr, "%d", &hours)
		if err == nil {
			// Buat location berdasarkan offset
			loc = time.FixedZone(fmt.Sprintf("GMT%s", offsetStr), hours*3600)
		} else {
			// Fallback ke UTC
			loc = time.UTC
		}
	} else {
		// Coba parsing sebagai IANA timezone (e.g., Asia/Jakarta)
		loc, err = time.LoadLocation(timezoneStr)
		if err != nil {
			log.Printf("Error loading timezone %s: %v. Using UTC as fallback.", timezoneStr, err)
			loc = time.UTC
		}
	}

	// Konversi ke waktu lokal
	localTime := utcTime.In(loc)

	// Format waktu
	return localTime.Format("2006-01-02 15:04:05 MST")
}
