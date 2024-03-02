-- +goose Up
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        surname VARCHAR(255) NOT NULL,
                        code INT NOT NULL,
                        operation_code INT NOT NULL,
                        transaction_code VARCHAR(255) NOT NULL,
                        status VARCHAR(20) DEFAULT 'NEW' NOT NULL
);

CREATE INDEX idx_transaction_code ON orders (transaction_code);

-- +goose Down
DROP TABLE orders;
