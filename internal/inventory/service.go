package inventory

import (
	"errors"
	"time"
)

type Service interface {
	CreateItem(req CreateItemRequest) (*Item, error)
	GetItems() ([]Item, error)
	GetItem(id uint) (*Item, error)

	// Transactions
	CreateTransaction(req CreateTransactionRequest) (*InventoryTransaction, error)
	GetTransactions(itemID uint) ([]InventoryTransaction, error)

	//Stock
	GetStock(itemID uint) (*StockResponse, error)
	CalculateCurrentStock(itemID uint) (int, error)

	// Snapshots
	CreateSnapshot(itemID uint) (*InventorySnapshot, error)
	GetSnapshots(itemID uint) ([]InventorySnapshot, error)

	// Analytics
	GetAnalytics(itemID uint, days int) (*AnalyticsResponse, error)
	GetForecast(itemID uint, historyDays int, forecastDays int) (*ForecastResponse, error)
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

	createdAt := time.Now()

	if req.Timestamp != nil {
		createdAt = *req.Timestamp
	}

	transaction := &InventoryTransaction{
		ItemID:          req.ItemID,
		TransactionType: req.TransactionType,
		Direction:       req.Direction,
		Quantity:        req.Quantity,
		Reference:       req.Reference,
		Notes:           req.Notes,
		CreatedAt:       createdAt,
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

func (s *service) CalculateCurrentStock(itemID uint) (int, error) {

	transactions, err := s.repo.GetTransactionsForItem(itemID)

	if err != nil {
		return 0, err
	}

	stock := 0

	for _, t := range transactions {

		if t.Direction == Inbound {
			stock += t.Quantity
		}

		if t.Direction == Outbound {
			stock -= t.Quantity
		}
	}

	return stock, nil
}

func (s *service) CreateSnapshot(itemID uint) (*InventorySnapshot, error) {

	// Verify item exists
	_, err := s.repo.GetByID(itemID)
	if err != nil {
		return nil, err
	}

	// Calculate current inventory
	stock, err := s.CalculateCurrentStock(itemID)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	snapshotDate := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0,
		0,
		0,
		0,
		now.Location(),
	)

	snapshot := &InventorySnapshot{
		ItemID:       itemID,
		Quantity:     stock,
		SnapshotDate: snapshotDate,
	}

	err = s.repo.CreateSnapshot(snapshot)

	if err != nil {
		return nil, err
	}

	return snapshot, nil
}

func (s *service) GetSnapshots(itemID uint) ([]InventorySnapshot, error) {

	_, err := s.repo.GetByID(itemID)

	if err != nil {
		return nil, err
	}

	return s.repo.GetSnapshotsForItem(itemID)
}

func (s *service) GetAnalytics(itemID uint, days int) (*AnalyticsResponse, error) {

	// Verify item exists
	_, err := s.repo.GetByID(itemID)
	if err != nil {
		return nil, err
	}

	// Current stock
	currentStock, err := s.CalculateCurrentStock(itemID)
	if err != nil {
		return nil, err
	}

	transactions, err := s.repo.GetTransactionsForItem(itemID)
	if err != nil {
		return nil, err
	}

	if days <= 0 {
		days = 30
	}

	analytics := &AnalyticsResponse{
		ItemID:             itemID,
		CurrentStock:       currentStock,
		AnalysisWindowDays: days,
	}

	cutoff := time.Now().AddDate(0, 0, -days)

	unitsSold := 0
	var lastSale *time.Time

	for _, t := range transactions {

		// Only count outbound sales
		if t.TransactionType != Sale {
			continue
		}

		if t.Direction != Outbound {
			continue
		}

		// Ignore sales older than 30 days
		if t.CreatedAt.Before(cutoff) {
			continue
		}

		unitsSold += t.Quantity

		if lastSale == nil || t.CreatedAt.After(*lastSale) {
			temp := t.CreatedAt
			lastSale = &temp
		}
	}

	analytics.UnitsSold = unitsSold
	analytics.LastSale = lastSale

	analytics.AverageDailyDemand = float64(unitsSold) / float64(days)
	analytics.AverageWeeklyDemand = analytics.AverageDailyDemand * 7.0

	if analytics.AverageDailyDemand > 0 {
		analytics.DaysOfInventoryRemaining =
			float64(currentStock) / analytics.AverageDailyDemand
	}

	return analytics, nil
}

func (s *service) GetForecast(itemID uint, historyDays int, forecastDays int) (*ForecastResponse, error) {

	// Verify item exists
	_, err := s.repo.GetByID(itemID)

	if err != nil {
		return nil, err
	}

	if historyDays <= 0 {
		historyDays = 30
	}

	if forecastDays <= 0 {
		forecastDays = 7
	}

	transactions, err := s.repo.GetTransactionsForItem(itemID)

	if err != nil {
		return nil, err
	}

	cutoff := time.Now().AddDate(0, 0, -historyDays)

	unitsSold := 0

	for _, t := range transactions {

		if t.TransactionType != Sale {
			continue
		}

		if t.Direction != Outbound {
			continue
		}

		if t.CreatedAt.Before(cutoff) {
			continue
		}

		unitsSold += t.Quantity
	}

	averageDailyDemand := float64(unitsSold) / float64(historyDays)

	response := &ForecastResponse{
		ItemID:             itemID,
		ForecastDays:       forecastDays,
		HistoricalDays:     historyDays,
		AverageDailyDemand: averageDailyDemand,
	}

	for i := 0; i < forecastDays; i++ {

		response.DailyForecast =
			append(
				response.DailyForecast,
				averageDailyDemand,
			)

		response.ForecastedDemand += int(averageDailyDemand)
	}

	return response, nil
}
