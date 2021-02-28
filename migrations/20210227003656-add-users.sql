-- +migrate Up
CREATE TYPE usertype AS ENUM ('Lastfm', 'Wavy');
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  discord_id text,
  username text,
  user_type usertype,
  last_indexed timestamp with time zone
);
-- +migrate Down
DROP TABLE users;
DROP TYPE usertype;