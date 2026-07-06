package inventory

import "time"

type Item struct {
	ID uint `gorm:"column:id;primaryKey" json:"id"`

	ItemNumber  string `gorm:"column:item_number" json:"itemNumber"`
	Description string `gorm:"column:description" json:"description"`
	Category    string `gorm:"column:category" json:"category"`

	UnitCost  float64 `gorm:"column:unit_cost" json:"unitCost"`
	UnitPrice float64 `gorm:"column:unit_price" json:"unitPrice"`

	MinimumStock int `gorm:"column:minimum_stock" json:"minimumStock"`
	SafetyStock  int `gorm:"column:safety_stock" json:"safetyStock"`

	IsActive bool `gorm:"column:is_active" json:"isActive"`

	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Item) TableName() string {
	return "items"
}
