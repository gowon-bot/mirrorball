-- +migrate Up
CREATE TABLE guild_members (
  guild_id text,
  user_id integer REFERENCES "users"(id) ON DELETE CASCADE,
  
  CONSTRAINT gm_uniqueness UNIQUE (user_id, guild_id)
);
-- +migrate Down
DROP TABLE guild_members;
