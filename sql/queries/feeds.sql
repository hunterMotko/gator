-- name: AddFeed :one
INSERT INTO feeds (id, name, url, user_id) 
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetFeeds :many
SELECT f.name as feed_name, f.url, u.name as user_name
	FROM feeds f
	INNER JOIN users u
	ON f.user_id = u.id;

-- name: GetFeedId :one
SELECT f.id as feed_id 
	FROM feeds f
	INNER JOIN users u
	ON f.user_id = u.id
	WHERE f.url = $1 AND u.id = $2;

-- name: GetFeedByUrl :one
SELECT * FROM feeds WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
	SET last_fetched_at = NOW(), updated_at = NOW()
	WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST, created_at ASC
LIMIT 1;
