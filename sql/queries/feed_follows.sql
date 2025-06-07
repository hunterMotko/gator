-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, user_id, feed_id)
	VALUES ($1, $2, $3)
  RETURNING *
) SELECT iff.*, f.name AS feed_name, u.name AS user_name
	FROM inserted_feed_follow iff
	INNER JOIN users u ON iff.user_id = u.id
	INNER JOIN feeds f ON iff.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT iff.id, u.name as user_name, f.name as feed_name
	FROM feed_follows iff
	INNER JOIN users u ON iff.user_id = u.id
	INNER JOIN feeds f ON iff.feed_id = f.id
	WHERE iff.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows ff
USING feeds fe
	WHERE ff.user_id = $1 AND fe.url = $2;
