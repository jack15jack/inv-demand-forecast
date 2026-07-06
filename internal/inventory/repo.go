package inventory

import "gorm.io/gorm"

type Repository interface {
	Create(item *Item) error
	GetAll() ([]Item, error)
	GetByID(id uint) (*Item, error)

	// Inventory Transactions
	CreateTransaction(transaction *InventoryTransaction) error
	GetTransactionsForItem(itemID uint) ([]InventoryTransaction, error)
	GetAllTransactions() ([]InventoryTransaction, error)
	GetByItemNumber(itemNumber string) (*Item, error)

	CreateSnapshot(snapshot *InventorySnapshot) error
	GetSnapshotsForItem(itemID uint) ([]InventorySnapshot, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(item *Item) error {
	return r.db.Create(item).Error
}

func (r *repository) GetAll() ([]Item, error) {
	var items []Item

	err := r.db.Order("item_number").Find(&items).Error

	return items, err
}

func (r *repository) GetByID(id uint) (*Item, error) {
	var item Item

	err := r.db.First(&item, id).Error

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *repository) CreateTransaction(transaction *InventoryTransaction) error {
	return r.db.Create(transaction).Error
}

func (r *repository) GetTransactionsForItem(itemID uint) ([]InventoryTransaction, error) {
	var transactions []InventoryTransaction

	err := r.db.
		Where("item_id = ?", itemID).
		Order("created_at ASC").
		Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetAllTransactions() ([]InventoryTransaction, error) {
	var transactions []InventoryTransaction

	err := r.db.
		Order("created_at DESC").
		Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetByItemNumber(itemNumber string) (*Item, error) {
	var item Item

	err := r.db.
		Where("item_number = ?", itemNumber).
		First(&item).Error

	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *repository) CreateSnapshot(snapshot *InventorySnapshot) error {
	return r.db.Create(snapshot).Error
}

func (r *repository) GetSnapshotsForItem(itemID uint) ([]InventorySnapshot, error) {

	var snapshots []InventorySnapshot

	err := r.db.
		Where("item_id = ?", itemID).
		Order("snapshot_date ASC").
		Find(&snapshots).Error

	return snapshots, err
}
