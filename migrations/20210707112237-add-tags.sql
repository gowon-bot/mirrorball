-- +migrate Up
CREATE TABLE tags (
  id SERIAL PRIMARY KEY,
  name text
);
CREATE TABLE artist_tags (
  artist_id integer REFERENCES artists(id),
  tag_id integer REFERENCES tags(id)
);

-- +migrate Down
DROP TABLE tags;
DROP TABLE artists;