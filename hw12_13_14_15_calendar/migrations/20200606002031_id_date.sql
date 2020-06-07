-- +goose Up
CREATE UNIQUE INDEX IF NOT EXISTS idx_date  ON events (date);

-- +goose Down
DROP INDEX IF EXISTS idx_date;