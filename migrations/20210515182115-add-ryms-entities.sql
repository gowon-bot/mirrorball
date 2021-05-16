-- +migrate Up
CREATE TABLE rate_your_music_albums (
  id SERIAL PRIMARY KEY,
  rate_your_music_id TEXT,
  release_year INTEGER
);

CREATE TABLE rate_your_music_album_albums (
  rate_your_music_album_id INTEGER REFERENCES rate_your_music_albums(id), 
  album_id INTEGER REFERENCES albums(id)
);

CREATE TABLE ratings (
  id SERIAL PRIMARY KEY,
  rating INTEGER,
  user_id  INTEGER REFERENCES "users"(id),
  rate_your_music_album_id INTEGER REFERENCES rate_your_music_albums(id)
);

-- +migrate Down
DROP TABLE ratings;
DROP TABLE rate_your_music_album_albums;
DROP TABLE rate_your_music_albums;
