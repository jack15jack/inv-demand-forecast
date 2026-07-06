package inventory

import "errors"

type Service interface {
	CreateItem(req CreateItemRequest) (*Item, error)
	GetItems() ([]Item, error)
	GetItem(id uint) (*Item, error)

	// Transactions
	CreateTransaction(req CreateTransactionRequest) (*InventoryTransaction, error)
	GetTransactions(itemID uint) ([]InventoryTransaction, error)

	GetStock(itemID uint) (*StockResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateItem(req CreateItemRequest) (*Item, error) {

	item := &Item{
		ItemNumber:   req.ItemNumber,
		Description:  req.Description,
		Category:     req.Category,
		UnitCost:     req.UnitCost,
		UnitPrice:    req.UnitPrice,
		MinimumStock: req.MinimumStock,
		SafetyStock:  req.SafetyStock,
		IsActive:     true,
	}

	err := s.repo.Create(item)

	return item, err
}

func (s *service) GetItems() ([]Item, error) {
	return s.repo.GetAll()
}

func (s *service) GetItem(id uint) (*Item, error) {
	return s.repo.GetByID(id)
}

func (s *service) CreateTransaction(req CreateTransactionRequest) (*InventoryTransaction, error) {

	// Verify item exists
	_, err := s.repo.GetByID(req.ItemID)
	if err != nil {
		return nil, err
	}

	// Validate quantity
	if req.Quantity <= 0 {
		return nil, errors.New("quantity must be greater than zero")
	}

	//Validate direction
	if req.Direction != Inbound && req.Direction != Outbound {
		return nil, errors.New("invalid transaction direction")
	}

	transaction := &InventoryTransaction{
		ItemID:          req.ItemID,
		TransactionType: req.TransactionType,
		Direction:       req.Direction,
		Quantity:        req.Quantity,
		Reference:       req.Reference,
		Notes:           req.Notes,
	}
	err = s.repo.CreateTransaction(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *service) GetTransactions(itemID uint) ([]InventoryTransaction, error) {
	return s.repo.GetTransactionsForItem(itemID)
}

func (s *service) GetStock(itemID uint) (*StockResponse, error) {

	// Verify the item exists
	_, err := s.repo.GetByID(itemID)
	if err != nil {
		return nil, err
	}

	transactions, err := s.repo.GetTransactionsForItem(itemID)
	if err != nil {
		return nil, err
	}

	stock := &StockResponse{
		ItemID: itemID,
	}

	for _, t := range transactions {

		if t.Direction == Inbound {
			stock.CurrentStock += t.Quantity
		}

		if t.Direction == Outbound {
			stock.CurrentStock -= t.Quantity
		}
	}

	return stock, nil
}
