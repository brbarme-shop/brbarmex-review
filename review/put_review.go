package review

import (
	"context"
	"errors"
)

var (
	ErrPutRatingInputInvalid = errors.New("the ReviewInput is invalid. The struct and your fields cann't be NIL. Check all fields or see more about in documentation")
)

type PutReviewRepository interface {
	PutNewReview(ctx context.Context, itemId string, comment string) error
}

type ReviewInput struct {
	ItemId  string `json:"item_id"`
	Comment string `json:"comment"`
}

func PutReview(ctx context.Context, reviewInput *ReviewInput, db PutReviewRepository) error {

	isInvalid := reviewInput == nil || len(reviewInput.ItemId) == 0 || len(reviewInput.Comment) == 0
	if isInvalid {
		return ErrPutRatingInputInvalid
	}

	return db.PutNewReview(ctx, reviewInput.ItemId, reviewInput.ItemId)
}
