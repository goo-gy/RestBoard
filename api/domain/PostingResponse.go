package domain

import (
	"time"
)

type PostingResponse struct {
	Id        int64
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPostingResponse(posting Posting) PostingResponse {
	return PostingResponse{
		Id:        posting.Id,
		Title:     posting.Title,
		Content:   posting.Content,
		CreatedAt: posting.CreatedAt,
		UpdatedAt: posting.UpdatedAt,
	}
}
