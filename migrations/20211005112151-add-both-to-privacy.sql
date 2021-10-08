-- +migrate Up
ALTER TYPE privacy
RENAME TO privacy__;
CREATE TYPE privacy AS ENUM (
  'PRIVATE',
  'DISCORD',
  'FMUSERNAME',
  'UNSET',
  'BOTH'
);
ALTER TABLE users
ALTER COLUMN privacy TYPE privacy USING privacy::text::privacy;
DROP TYPE privacy__;
-- +migrate Down
ALTER TYPE privacy
RENAME TO privacy__;
CREATE TYPE privacy AS ENUM (
  'PRIVATE',
  'DISCORD',
  'FMUSERNAME',
  'UNSET',
);
ALTER TABLE users
ALTER COLUMN privacy TYPE privacy USING privacy::text::privacy;
DROP TYPE privacy__;