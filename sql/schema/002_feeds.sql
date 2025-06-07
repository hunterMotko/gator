-- +goose Up
CREATE TABLE feeds (
	id uuid PRIMARY KEY,
	created_at timestamp NOT NULL DEFAULT current_timestamp,
	updated_at timestamp NOT NULL DEFAULT current_timestamp,
	name varchar(256) UNIQUE NOT NULL,
	url text NOT NULL,
	user_id uuid NOT NULL,
	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
