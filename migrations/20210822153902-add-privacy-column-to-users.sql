-- +migrate Up
CREATE TYPE privacy AS ENUM ('PRIVATE', 'DISCORD', 'FMUSERNAME', 'UNSET');
ALTER TABLE users
ADD "privacy" privacy;
UPDATE users
SET privacy = 'UNSET';
-- +migrate Down
ALTER TABLE users DROP privacy;
DROP TYPE privacy;