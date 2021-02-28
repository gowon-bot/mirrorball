-- +migrate Up
CREATE TABLE artist_counts (
  id SERIAL PRIMARY KEY,
  playcount integer,
  artist_id integer REFERENCES artists(id),
  user_id integer REFERENCES "users"(id)
);
-- +migrate Down
DROP TABLE artist_counts;
