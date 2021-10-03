-- +migrate Up
-- This fuckup will be immortalized in my migrations forever :suffering2:
CREATE INDEX tc_name_idx ON tracks ((name));
CREATE INDEX al_name_idx ON albums ((name));
CREATE INDEX tr_name_idx ON tracks ((name));
-- +migrate Down
DROP INDEX tc_name_idx;
DROP INDEX al_name_idx;
DROP INDEX tr_name_idx;