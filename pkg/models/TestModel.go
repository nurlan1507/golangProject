package models

import (
	"time"
)

type TestModel struct {
	Id          int       `json:"Id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SubjectId   int       `json:"subjectId"`
	AuthorId    int       `json:"authorId"`
	GroupId     string    `json:"group"`
	CreatedAt   time.Time `json:"create_at"`
	StartAt     time.Time `json:"st"`
	ExpiresAt   time.Time `json:"expires_at"`
}
