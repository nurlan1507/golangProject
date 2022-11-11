package models

type TestModel struct {
	Id          int    `json:"Id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SubjectId   int    `json:"subjectId"`
	AuthorId    int    `json:"authorId"`
	GroupId     int    `json:"groupId"`
	//CreatedAt   time.Time `json:"create_at"`
	//StartAt     time.Time `json:"startA"`
	//ExpiresAt   time.Time `json:"expires_at"`
}
