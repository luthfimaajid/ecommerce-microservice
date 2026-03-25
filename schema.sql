CREATE SCHEMA IF NOT EXISTS "order";
CREATE SCHEMA IF NOT EXISTS inventory;

CREATE TABLE "order".product (
    id            BIGSERIAL       PRIMARY KEY,
    category      VARCHAR(200)    NOT NULL,
    name          VARCHAR(200)    NOT NULL,
    description   TEXT,
    price         NUMERIC(15, 2)  NOT NULL,
    stock         INT             NOT NULL DEFAULT 0,
    is_active     BOOLEAN         NOT NULL DEFAULT true,
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT now(),
    updated_at    TIMESTAMPTZ     NOT NULL DEFAULT now(),

    CONSTRAINT chk_product_price_positive CHECK (price > 0),
    CONSTRAINT chk_product_stock_non_neg  CHECK (stock >= 0)
);

CREATE TABLE "order"."order" (
    id              BIGSERIAL       PRIMARY KEY,
    user_id         BIGINT          NOT NULL,
    product_id      BIGINT          NOT NULL REFERENCES "order".product(id),
    quantity        INT             NOT NULL,
    unit_price      NUMERIC(15,2)   NOT NULL,
    total_amount    NUMERIC(15,2)   GENERATED ALWAYS AS (unit_price * quantity) STORED,
    status          TEXT            NOT NULL DEFAULT 'pending',
    note            TEXT,
    created_at      TIMESTAMPTZ     NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ     NOT NULL DEFAULT now(),

    CONSTRAINT chk_order_qty_positive CHECK (quantity > 0),
    CONSTRAINT chk_order_price_positive CHECK (unit_price > 0),
    CONSTRAINT chk_order_status CHECK (status IN ('pending', 'success', 'cancelled'))
);

CREATE INDEX idx_order_user_id   ON "order"."order"(user_id);
CREATE INDEX idx_order_product   ON "order"."order"(product_id);
CREATE INDEX idx_order_status    ON "order"."order"(status);

CREATE INDEX idx_product_active  ON "order".product(is_active);

CREATE TABLE inventory.stock_history (
    id              BIGSERIAL PRIMARY KEY,
    product_id      BIGINT NOT NULL,
    order_id        BIGINT,
    user_id         BIGINT NOT NULL,
    movement_type   TEXT NOT NULL,
    quantity        INT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT chk_movement_qty_nonzero CHECK (quantity <> 0),
    CONSTRAINT chk_movement_type CHECK (
        movement_type IN (
            'stock_in',
            'stock_out',
            'order_pending',
            'order_success',
            'order_cancelled'
        )
    )
);

CREATE INDEX idx_stock_product   ON inventory.stock_history(product_id);
CREATE INDEX idx_stock_order     ON inventory.stock_history(order_id);
CREATE INDEX idx_stock_user      ON inventory.stock_history(user_id);
CREATE INDEX idx_stock_product_created ON inventory.stock_history(product_id, created_at);

INSERT INTO "order".product (category, name, description, price, stock)
VALUES
('Electronics', 'Wireless Mouse', '2.4GHz ergonomic mouse', 15.90, 200),
('Electronics', 'Mechanical Keyboard', 'RGB mechanical keyboard', 79.90, 150),
('Electronics', 'USB-C Hub', '7-in-1 USB-C hub', 39.50, 120),
('Electronics', '27 inch Monitor', '4K IPS monitor', 329.00, 60),
('Electronics', 'Laptop Stand', 'Aluminum adjustable stand', 29.90, 180),

('Books', 'Clean Code', 'Robert C. Martin', 42.00, 80),
('Books', 'Designing Data Intensive Applications', 'Martin Kleppmann', 55.00, 60),
('Books', 'The Pragmatic Programmer', 'Programming classic', 45.50, 90),

('Coffee', 'Ethiopian Yirgacheffe', 'Specialty coffee beans 250g', 18.00, 200),
('Coffee', 'Colombian Supremo', 'Specialty coffee beans 250g', 16.50, 220);

INSERT INTO inventory.stock_history
(product_id, user_id, movement_type, quantity)
VALUES
(1, 1, 'stock_in', 200),
(2, 1, 'stock_in', 150),
(3, 1, 'stock_in', 120),
(4, 1, 'stock_in', 60),
(5, 1, 'stock_in', 180),
(6, 1, 'stock_in', 80),
(7, 1, 'stock_in', 60),
(8, 1, 'stock_in', 90),
(9, 1, 'stock_in', 200),
(10,1, 'stock_in', 220);

INSERT INTO "order"."order"
(user_id, product_id, quantity, unit_price, status, note)
VALUES
(1, 1, 2, 15.90, 'success', 'Office mouse'),
(2, 6, 1, 42.00, 'success', 'Book purchase'),
(3, 5, 1, 29.90, 'pending', 'Laptop accessory'),
(1, 9, 3, 18.00, 'success', 'Coffee beans'),
(4, 4, 1, 329.00, 'success', 'Monitor purchase'),
(2, 8, 2, 45.50, 'cancelled', 'Changed mind'),
(5, 3, 1, 39.50, 'success', 'USB hub'),
(3, 10, 4, 16.50, 'pending', 'Coffee stock');


INSERT INTO inventory.stock_history
(product_id, order_id, user_id, movement_type, quantity)
VALUES
(1, 1, 1, 'order_pending', -2),
(1, 1, 1, 'order_success', -2);

INSERT INTO inventory.stock_history
(product_id, order_id, user_id, movement_type, quantity)
VALUES
(6, 2, 2, 'order_pending', -1),
(6, 2, 2, 'order_success', -1);

INSERT INTO inventory.stock_history
(product_id, order_id, user_id, movement_type, quantity)
VALUES
(5, 3, 3, 'order_pending', -1);

INSERT INTO inventory.stock_history
(product_id, order_id, user_id, movement_type, quantity)
VALUES
(9, 4, 1, 'order_pending', -3),
(9, 4, 1, 'order_success', -3);

INSERT INTO inventory.stock_history
(product_id, order_id, user_id, movement_type, quantity)
VALUES
(4, 5, 4, 'order_pending', -1),
(4, 5, 4, 'order_success', -1);

INSERT INTO inventory.stock_history
(product_id, order_id, user_id, movement_type, quantity)
VALUES
(8, 6, 2, 'order_pending', -2),
(8, 6, 2, 'order_cancelled', 2);

INSERT INTO inventory.stock_history
(product_id, order_id, user_id, movement_type, quantity)
VALUES
(3, 7, 5, 'order_pending', -1),
(3, 7, 5, 'order_success', -1);

INSERT INTO inventory.stock_history
(product_id, order_id, user_id, movement_type, quantity)
VALUES
(10, 8, 3, 'order_pending', -4);