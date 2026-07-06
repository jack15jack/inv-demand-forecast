ALTER TABLE inventory_transactions
ADD COLUMN direction VARCHAR(10) NOT NULL DEFAULT 'IN'
CHECK (
    direction IN ('IN', 'OUT')
);