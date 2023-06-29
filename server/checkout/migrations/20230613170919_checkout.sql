-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Checkout (
    user_id bigint NOT NULL PRIMARY KEY,
    totalPrice integer NOT NULL,
    orderID bigint NOT NULL
);

CREATE TABLE IF NOT EXISTS CartItem (
    user_id bigint NOT NULL,
    sku integer NOT NULL,
    count smallint NOT NULL,
    PRIMARY KEY (user_id, sku),
    FOREIGN KEY (user_id) REFERENCES Checkout (user_id)
);

CREATE TYPE Item AS (
    sku integer,
    count smallint,
    name text,
    price integer
);

CREATE TABLE IF NOT EXISTS CartItemDetails (
    user_id bigint NOT NULL,
    item Item[] NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Checkout (user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS CartItemDetails;
DROP TABLE IF EXISTS CartItem;
DROP TABLE IF EXISTS Checkout;

-- +goose StatementEnd
