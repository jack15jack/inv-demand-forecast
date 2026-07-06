CREATE TABLE inventory_transactions (
    id BIGSERIAL PRIMARY KEY,

    item_id BIGINT NOT NULL
        REFERENCES items(id)
        ON DELETE RESTRICT,

    transaction_type VARCHAR(20) NOT NULL
        CHECK (
            transaction_type IN (
                'PURCHASE',
                'SALE',
                'RETURN',
                'ADJUSTMENT',
                'DAMAGE'
            )
        ),

    quantity INTEGER NOT NULL
        CHECK (quantity > 0),

    reference VARCHAR(100),

    notes TEXT,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_inventory_transactions_item
ON inventory_transactions(item_id);

CREATE INDEX idx_inventory_transactions_created_at
ON inventory_transactions(created_at);