// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, user_id, feed_id, created_at, updated_at
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.UserID,
		arg.FeedID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows 
WHERE id = $1
`

func (q *Queries) DeleteFeedFollow(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, id)
	return err
}

const findFeedFollowById = `-- name: FindFeedFollowById :one
SELECT id, user_id, feed_id, created_at, updated_at FROM feed_follows
WHERE id = $1
`

func (q *Queries) FindFeedFollowById(ctx context.Context, id uuid.UUID) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, findFeedFollowById, id)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllUserFeedFollows = `-- name: GetAllUserFeedFollows :many
SELECT id, user_id, feed_id, created_at, updated_at FROM feed_follows
WHERE user_id = $1
`

func (q *Queries) GetAllUserFeedFollows(ctx context.Context, userID uuid.UUID) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getAllUserFeedFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.FeedID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
