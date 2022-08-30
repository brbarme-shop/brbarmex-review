package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/brbarme-shop/brbarmex-review/review"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sql.DB
}

func (r *repository) PutNewReview(ctx context.Context, itemId, comment, curstomerId, orderId string) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_sql := `
INSERT INTO reviews(review_hash_id, review_item_id)
    SELECT uuid_generate_v1(), $1
WHERE NOT EXISTS (
    SELECT 1 FROM reviews WHERE review_item_id = $2
) RETURNING review_id ;
`

	var review_id int64
	err = tx.QueryRowContext(ctx, _sql, itemId, itemId).Scan(&review_id)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	if review_id > 0 {

		_sql = `
INSERT INTO review_comments
	(review_id, review_comment, review_customer_id, review_order_id, review_datatime)
VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP);
`
		row, err := tx.ExecContext(ctx, _sql, review_id, comment, curstomerId, orderId)
		if err != nil {
			tx.Rollback()
			return err
		}

		rowsAffected, err := row.RowsAffected()
		if err != nil || rowsAffected <= 0 {
			tx.Rollback()
			return err
		}

		return tx.Commit()
	}

	_sql = `
INSERT INTO reviews
	(review_hash_id, review_item_id)
VALUES(uuid_generate_v1(), $1) RETURNING review_id ;
	`

	err = tx.QueryRowContext(ctx, _sql, itemId).Scan(&review_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if review_id <= 0 {
		tx.Rollback()
		return err
	}

	_sql = `
INSERT INTO review_comments
		(review_id, review_comment, review_customer_id, review_order_id, review_datatime)
VALUES($1, $2, $3, $4, CURRENT_TIMESTAMP);`

	row, err := tx.ExecContext(ctx, _sql, review_id, comment, curstomerId, orderId)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := row.RowsAffected()
	if err != nil || rowsAffected <= 0 {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func NewPutReviewRepository(db *sql.DB) review.PutReviewRepository {
	return &repository{db: db}
}
