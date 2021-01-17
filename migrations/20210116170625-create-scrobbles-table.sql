-- +migrate Up
CREATE TABLE scrobbles (
  id SERIAL PRIMARY KEY,
  "timestamp" timestamp with time zone,
  user_id integer REFERENCES "users"(id),
  track_id integer REFERENCES tracks(id)
);
-- +migrate Down
DROP TABLE scrobbles;