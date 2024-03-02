-- +goose Up
CREATE TABLE outbox (
                        id SERIAL PRIMARY KEY,
                        uuid VARCHAR(36) NOT NULL,
                        user_id VARCHAR(36) NOT NULL,
                        transaction_code VARCHAR(255) NOT NULL,
                        json TEXT NOT NULL,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE outbox;
