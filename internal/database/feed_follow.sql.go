// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feed_follow.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id,createdAt,updatedAt,userId,feedId)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, createdat, updatedat, userid, feedid
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	Createdat time.Time
	Updatedat time.Time
	Userid    uuid.UUID
	Feedid    uuid.UUID
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.Createdat,
		arg.Updatedat,
		arg.Userid,
		arg.Feedid,
	)
	var i FeedFollow
	err := row.Scan(
		&i.ID,
		&i.Createdat,
		&i.Updatedat,
		&i.Userid,
		&i.Feedid,
	)
	return i, err
}

const deleteFeedFollow = `-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id =$1 AND userId =$2
`

type DeleteFeedFollowParams struct {
	ID     uuid.UUID
	Userid uuid.UUID
}

func (q *Queries) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollow, arg.ID, arg.Userid)
	return err
}

const getFeedFollow = `-- name: GetFeedFollow :many
SELECT id, createdat, updatedat, userid, feedid FROM feed_follows WHERE userId= $1
`

func (q *Queries) GetFeedFollow(ctx context.Context, userid uuid.UUID) ([]FeedFollow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollow, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FeedFollow
	for rows.Next() {
		var i FeedFollow
		if err := rows.Scan(
			&i.ID,
			&i.Createdat,
			&i.Updatedat,
			&i.Userid,
			&i.Feedid,
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