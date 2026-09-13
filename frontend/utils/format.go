package utils

import (
	"fmt"
	"strings"
	"time"
)

func Size(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	value := strings.TrimSuffix(fmt.Sprintf("%.1f", float64(bytes)/float64(div)), ".0")
	return fmt.Sprintf("%s %cB", value, "KMGTPE"[exp])
}

func Count(n int, noun string) string {
	if n == 1 {
		return fmt.Sprintf("%d %s", n, noun)
	}
	return fmt.Sprintf("%d %ss", n, noun)
}

func Date(t time.Time) string {
	return t.Format(time.DateTime)
}

func Timestamp(t time.Time) string {
	return t.Format(time.RFC3339)
}
