-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE wallet
    ALTER COLUMN user_id TYPE VARCHAR(255) USING user_id::VARCHAR;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE wallet
    ALTER COLUMN user_id TYPE BIGINT USING user_id::BIGINT;
-- +goose StatementEnd
