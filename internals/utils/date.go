package utils

import "time"

func GetCurrentDate() string {
	now := time.Now().UTC()
	return now.Format(time.RFC3339)
}
