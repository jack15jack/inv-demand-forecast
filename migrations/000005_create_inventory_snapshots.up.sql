CREATE TABLE inventory_snapshots (
    id BIGSERIAL PRIMARY KEY,

    item_id BIGINT NOT NULL
        REFERENCES items(id)
        ON DELETE RESTRICT,

    snapshot_date DATE NOT NULL,

    quantity INTEGER NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(item_id, snapshot_date)
);

CREATE INDEX idx_inventory_snapshots_item_date
ON inventory_snapshots(item_id, snapshot_date);