-- +migrate Up
DROP INDEX tc_name_idx;
CREATE INDEX ar_name_idx ON artists ((name));

-- +migrate Down
CREATE INDEX tc_name_idx ON tracks ((name));
DROP INDEX ar_name_idx;
