-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE INDEX idx_wallets_user_id ON wallets(user_id);
CREATE INDEX idx_transactions_wallet_id ON transactions(wallet_id);
CREATE INDEX idx_transactions_time ON transactions(time);
CREATE INDEX idx_transactions_wallet_id_time ON transactions(wallet_id, time);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP INDEX IF EXISTS idx_wallets_user_id;
DROP INDEX IF EXISTS idx_transactions_wallet_id;
DROP INDEX IF EXISTS idx_transactions_time;
DROP INDEX IF EXISTS idx_transactions_wallet_id_time;
-- +goose StatementEnd
