package helpers

import (
	"strings"
	"time"
)

func NotEmpty(value string) bool {
	return strings.TrimSpace(value) != ""
}

func NotEmptyTime(value time.Time) bool {
	return !value.IsZero()
}
