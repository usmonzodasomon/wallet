-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE wallets
    ADD COLUMN is_identified BOOLEAN DEFAULT FALSE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE wallets
    DROP COLUMN is_identified;
-- +goose StatementEnd
