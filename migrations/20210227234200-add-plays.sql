-- +migrate Up
CREATE TABLE plays (
  id SERIAL PRIMARY KEY,
  scrobbled_at timestamp with time zone,
  user_id integer REFERENCES "users"(id) ON DELETE CASCADE,
  track_id integer REFERENCES tracks(id),
  
  CONSTRAINT uniqueness UNIQUE (scrobbled_at, track_id, user_id)
);
-- +migrate Down
DROP TABLE plays;
