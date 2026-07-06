package inventory

import "gorm.io/gorm"

type Repository interface {
	Create(item *Item) error
	GetAll() ([]Item, error)
	GetByID(id uint) (*Item, error)
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
	err := r.db.Find(&items).Error
	return items, err
}

func (r *repository) GetByID(id uint) (*Item, error) {
	var item Item
	err := r.db.First(&item, id).Error
	return &item, err
}
