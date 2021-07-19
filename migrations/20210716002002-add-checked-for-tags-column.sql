-- +migrate Up
ALTER TABLE artists ADD checked_for_tags boolean;

-- +migrate Down
ALTER TABLE artists DROP checked_for_tags;
