-- +migrate Up
CREATE TABLE album_counts (
  id SERIAL PRIMARY KEY,
  playcount integer,
  album_id integer REFERENCES albums(id),
  user_id integer REFERENCES "users"(id)
);
-- +migrate Down
DROP TABLE album_counts;
