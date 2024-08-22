-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE IF NOT EXISTS wallet (
    id BIGSERIAL PRIMARY KEY,
    balance BIGINT,
    user_id BIGINT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE wallet;
-- +goose StatementEnd
