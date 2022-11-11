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
	for _, element := range question.Answers {
		if element.Correct == true {
			return true
		}
	}
	return false
}
func NoDescription(question *models.QuestionModel) bool {
	if len(strings.Trim(question.Description, "")) == 0 || question.Description == "" {
		return false
	}
	return true
}
