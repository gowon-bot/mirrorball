-- +migrate Up
CREATE TABLE track_counts (
  id SERIAL PRIMARY KEY,
  playcount integer,
  track_id integer REFERENCES tracks(id),
  user_id integer REFERENCES "users"(id)
);
-- +migrate Down
DROP TABLE track_counts;