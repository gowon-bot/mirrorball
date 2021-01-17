-- +migrate Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  last_fm_username text,
  last_indexed timestamp with time zone
);
-- +migrate Down
DROP TABLE users;