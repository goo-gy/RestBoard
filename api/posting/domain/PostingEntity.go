package domain

import (
	"time"
)

type Posting struct {
	Id        int64 `gorm:"primaryKey;autoIncrement:true"`
	Title     string
	Content   string
	CreatedAt time.Time `gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:true"`
}

func NewPosting(postingRequest PostingRequest) Posting {
	return Posting{
		Title:   postingRequest.Title,
		Content: postingRequest.Content,
	}
}
