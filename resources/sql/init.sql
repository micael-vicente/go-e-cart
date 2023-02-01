CREATE TABLE IF NOT EXISTS carts (
    id          SERIAL PRIMARY KEY,
    customer_id varchar(100) NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX cart_customer_id_idx ON carts (customer_id);

CREATE TABLE IF NOT EXISTS cart_items (
    id         SERIAL PRIMARY KEY,
    cart_id    INT NOT NULL,
    sku        VARCHAR(100) NOT NULL,
    quantity   INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (cart_id) REFERENCES carts(id),
    CONSTRAINT cart_item_cart_id_sku_unique UNIQUE (cart_id, sku)
);

CREATE INDEX cart_item_cart_id_idx ON cart_items (cart_id);