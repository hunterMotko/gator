-- +goose Up
CREATE TABLE feed_follows (
	id uuid PRIMARY KEY,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp,
	user_id uuid NOT NULL,
	feed_id uuid NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
	FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose down
DROP TABLE feed_follows;
