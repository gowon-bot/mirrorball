-- +migrate Up
CREATE TABLE albums (
  id SERIAL PRIMARY KEY,
  name text,
  artist_id integer REFERENCES artists(id)
);
-- +migrate Down
DROP TABLE albums;
