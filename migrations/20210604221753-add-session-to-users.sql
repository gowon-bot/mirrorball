-- +migrate Up
ALTER TABLE users ADD last_fm_session text;

-- +migrate Down
ALTER TABLE users DROP last_fm_session;
