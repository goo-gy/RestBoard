package repository

import (
	"github.com/restBoard/api/domain"
	"github.com/restBoard/database"
)

func SavePosting(posting domain.Posting) (domain.Posting, error) {
	result := database.Db.Create(&posting)
	return posting, result.Error
}

func FindPostingById(postingId int64) (domain.Posting, error) {
	posting := domain.Posting{}
	result := database.Db.First(&posting, postingId)
	return posting, result.Error
}

func FindAllPosting() ([]domain.Posting, error) {
	postingList := []domain.Posting{}
	result := database.Db.Find(&postingList)
	return postingList, result.Error
}

func UpdatePosting(postingId int64, postingRequest domain.PostingRequest) (domain.Posting, error) {
	posting := domain.Posting{}
	result := database.Db.First(&posting, postingId)
	posting.Title = postingRequest.Title
	posting.Content = postingRequest.Content
	database.Db.Save(&posting)
	return posting, result.Error
}

func DeletePostingById(postingId int64) error {
	result := database.Db.Delete(&domain.Posting{}, postingId)
	return result.Error
}
