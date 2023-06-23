-- +goose Up
-- +goose StatementBegin
-- Таблица "Orders" для хранения информации о заказах
CREATE TABLE IF NOT EXISTS Orders (
    orderID bigint PRIMARY KEY,
    status text,
    user_id bigint
);

-- Таблица "OrderItems" для хранения информации о товарах в заказе
CREATE TABLE IF NOT EXISTS OrderItems (
    orderID bigint,
    sku bigint,
    count smallint,
    PRIMARY KEY (orderID, sku),
    FOREIGN KEY (orderID) REFERENCES Orders(orderID)
);

-- Таблица "Stocks" для хранения информации о доступном количестве товаров на складах
CREATE TABLE IF NOT EXISTS Stocks (
    warehouseID bigint,
    sku bigint,
    count smallint,
    PRIMARY KEY (warehouseID, sku)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Stocks;
DROP TABLE IF EXISTS OrderItems;
DROP TABLE IF EXISTS Orders;
-- +goose StatementEnd
