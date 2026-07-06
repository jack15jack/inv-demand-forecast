package inventory

import "time"

type InventorySnapshot struct {
	ID uint `gorm:"column:id;primaryKey"`

	ItemID uint `gorm:"column:item_id"`

	SnapshotDate time.Time `gorm:"column:snapshot_date"`

	Quantity int `gorm:"column:quantity"`

	CreatedAt time.Time `gorm:"column:created_at"`
}

func (InventorySnapshot) TableName() string {
	return "inventory_snapshots"
}
