package inventory

import "time"

type TransactionType string

const (
	Purchase   TransactionType = "PURCHASE"
	Sale       TransactionType = "SALE"
	Return     TransactionType = "RETURN"
	Adjustment TransactionType = "ADJUSTMENT"
	Damage     TransactionType = "DAMAGE"
)

type TransactionDirection string

const (
	Inbound  TransactionDirection = "IN"
	Outbound TransactionDirection = "OUT"
)

type InventoryTransaction struct {
	ID              uint                 `gorm:"column:id;primaryKey"`
	ItemID          uint                 `gorm:"column:item_id"`
	TransactionType TransactionType      `gorm:"column:transaction_type"`
	Direction       TransactionDirection `gorm:"column:direction"`
	Quantity        int                  `gorm:"column:quantity"`
	Reference       string               `gorm:"column:reference"`
	Notes           string               `gorm:"column:notes"`
	CreatedAt       time.Time            `gorm:"column:created_at"`
}

func (InventoryTransaction) TableName() string {
	return "inventory_transactions"
}
