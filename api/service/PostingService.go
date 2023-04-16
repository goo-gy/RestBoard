package service

import (
	"github.com/restBoard/api/domain"
	"github.com/restBoard/api/repository"
)

func CreatePosting(postingRequest domain.PostingRequest) (domain.PostingResponse, error) {
	posting := domain.NewPosting(postingRequest)
	savedPosting, err := repository.SavePosting(posting)
	if err != nil {
		return domain.PostingResponse{}, err
	}
	return domain.NewPostingResponse(savedPosting), nil
}

func DetailPosting(postingId int64) (domain.PostingResponse, error) {
	posting, err := repository.FindPostingById(postingId)
	if err != nil {
		return domain.PostingResponse{}, err
	}
	return domain.NewPostingResponse(posting), nil
}

func ListPosting() ([]domain.PostingResponse, error) {
	postingList, err := repository.FindAllPosting()
	if err != nil {
		return []domain.PostingResponse{}, err
	}
	return convertPostingResponseList(postingList), nil
}

func DeletePosting(postingId int64) error {
	return repository.DeletePostingById(postingId)
}

func UpdatePosting(postingId int64, postingRequest domain.PostingRequest) (domain.PostingResponse, error) {
	posting, err := repository.UpdatePosting(postingId, postingRequest)
	if err != nil {
		return domain.PostingResponse{}, err
	}
	return domain.NewPostingResponse(posting), nil
}

func convertPostingResponseList(postings []domain.Posting) []domain.PostingResponse {
	responseList := make([]domain.PostingResponse, len(postings))
	for i, posting := range postings {
		responseList[i] = domain.NewPostingResponse(posting)
	}
	return responseList
}
