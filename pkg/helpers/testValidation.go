package helpers

import (
	"fmt"
	"strings"
	"testApp/pkg/models"
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

func QuestionWithoutAnswers(question *models.QuestionModel) bool {
	if len(question.Answers) == 0 {
		return false
	}
	return true
}
func NoDescription(question *models.QuestionModel) bool {
	if len(strings.Trim(question.Description, "")) == 0 {
		return false
	}
	return true
}
