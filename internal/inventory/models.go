package inventory

import "time"

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ItemNumber  string `gorm:"unique;not null" json:"itemNumber"`
	Description string `json:"description"`
	Category    string `json:"category"`

	UnitCost  float64 `json:"unitCost"`
	UnitPrice float64 `json:"unitPrice"`

	MinimumStock int `json:"minimumStock"`
	SafetyStock  int `json:"safetyStock"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
