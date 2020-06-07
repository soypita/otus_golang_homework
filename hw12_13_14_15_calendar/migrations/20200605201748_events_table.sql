-- +goose Up
CREATE TABLE IF NOT EXISTS events (
    id UUID NOT NULL PRIMARY KEY,
    header text,
    date timestamp with time zone,
    duration bigint,
    description text,
    ownerId UUID,
    notifyBefore bigint
);

-- +goose Down
DROP TABLE IF EXISTS events
