package helpers

import (
	"fmt"
	"strings"
	"time"
)

func NotEmpty(value string) bool {
	return strings.TrimSpace(value) != ""
}

func NotEmptyTime(value time.Time) bool {
	return !value.IsZero()
}

func TimeIsValid(value time.Time) bool {
	fmt.Println(value)
	fmt.Println(value.Unix())
	if value.Unix() <= time.Now().Unix() {
		return false
	}
	return true
}
