package main

import (
	"log"
	"time"

	"github.com/jack15jack/inv-demand-forecast/internal/inventory"
	"gorm.io/gorm"
)

type Generator struct {
	db *gorm.DB
}

func NewGenerator(db *gorm.DB) *Generator {
	return &Generator{
		db: db,
	}
}

func (g *Generator) Run() error {

	log.Println("Clearing existing data...")

	if err := g.clearDatabase(); err != nil {
		return err
	}

	log.Println("Creating items...")

	items, err := g.seedItems()
	if err != nil {
		return err
	}

	log.Println("Generating transactions...")

	for _, item := range items {

		err := g.generateHistory(item)

		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) clearDatabase() error {

	if err := g.db.Exec(
		"TRUNCATE TABLE inventory_transactions RESTART IDENTITY CASCADE",
	).Error; err != nil {
		return err
	}

	if err := g.db.Exec(
		"TRUNCATE TABLE inventory_snapshots RESTART IDENTITY CASCADE",
	).Error; err != nil {
		return err
	}

	if err := g.db.Exec(
		"TRUNCATE TABLE items RESTART IDENTITY CASCADE",
	).Error; err != nil {
		return err
	}

	return nil
}

func (g *Generator) seedItems() ([]inventory.Item, error) {

	items := []inventory.Item{
		{
			ItemNumber:   "BOLT-001",
			Description:  "Standard Bolt",
			Category:     "Hardware",
			UnitCost:     0.25,
			UnitPrice:    0.75,
			MinimumStock: 500,
			SafetyStock:  200,
			IsActive:     true,
		},
		{
			ItemNumber:   "LIGHT-001",
			Description:  "Holiday Lights",
			Category:     "Seasonal",
			UnitCost:     8.00,
			UnitPrice:    19.99,
			MinimumStock: 100,
			SafetyStock:  50,
			IsActive:     true,
		},
		{
			ItemNumber:   "FILTER-001",
			Description:  "Replacement Filter",
			Category:     "Maintenance",
			UnitCost:     12.00,
			UnitPrice:    29.99,
			MinimumStock: 150,
			SafetyStock:  75,
			IsActive:     true,
		},
	}

	for i := range items {

		if err := g.db.Create(&items[i]).Error; err != nil {
			return nil, err
		}
	}

	return items, nil
}

func (g *Generator) generateHistory(item inventory.Item) error {

	startDate := time.Now().AddDate(-1, 0, 0)

	for d := startDate; d.Before(time.Now()); d = d.AddDate(0, 0, 1) {

		demand := baseDemand(item.ItemNumber, d)

		if demand > 0 {

			tx := inventory.InventoryTransaction{
				ItemID:          item.ID,
				TransactionType: inventory.Sale,
				Direction:       inventory.Outbound,
				Quantity:        demand,
				Reference:       "SIM-SALE",
				Notes:           "Generated sale",
			}

			if err := g.db.Create(&tx).Error; err != nil {
				return err
			}

			g.db.Model(&tx).Update("created_at", d)
		}
	}

	return nil
}
