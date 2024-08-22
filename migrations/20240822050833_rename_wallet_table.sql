-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE wallet RENAME TO wallets;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE wallets RENAME TO wallet;
-- +goose StatementEnd
