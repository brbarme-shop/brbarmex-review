package review

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestPutReview(t *testing.T) {

	ctx := context.TODO()

	t.Run("Should abort process when ReviewInput is invalid", func(t *testing.T) {

		var testTable = []struct {
			input *ReviewInput
		}{
			{
				input: &ReviewInput{ItemId: "", Comment: ""},
			},
			{
				input: &ReviewInput{},
			},
			{
				input: nil,
			},
		}

		db := &repositoryMock{
			putNewReview: func() error {
				return fmt.Errorf("failed")
			},
		}

		for _, tb := range testTable {

			err := PutReview(ctx, tb.input, db)

			if err == nil {
				t.Fatal("any message here")
			}

			if !errors.Is(err, ErrPutRatingInputInvalid) {
				t.Fatal("any message here")
			}
		}

	})

	t.Run("Should return error when Repository failed", func(t *testing.T) {

		db := &repositoryMock{
			putNewReview: func() error {
				return fmt.Errorf("failed")
			},
		}

		input := &ReviewInput{
			ItemId:  "DUMMY",
			Comment: "DUMMY",
		}

		err := PutReview(ctx, input, db)

		if err == nil {
			t.Fatal("")
		}

	})

}

type repositoryMock struct {
	putNewReview func() error
}

func (r *repositoryMock) PutNewReview(ctx context.Context, itemId string, comment string) error {
	return r.putNewReview()
}
