package inventory

import (
	"errors"
	"math"
	"sort"
	"time"
)

type Service interface {
	// Items
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

	// Recommendation
	GetPurchaseRecommendation(itemID uint, forecastDays int) (*PurchaseRecommendation, error)
	GetBatchPurchaseRecommendations(forecastDays int) ([]PurchaseRecommendation, error)
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

	// Get transactions
	transactions, err := s.repo.GetTransactionsForItem(itemID)
	if err != nil {
		return nil, err
	}

	currentStock, err := s.CalculateCurrentStock(itemID)
	if err != nil {
		return nil, err
	}

	// Calculate Analytics
	history := buildDemandHistory(transactions, historyDays)

	holt := holtLinearForecast(history, 0.3, 0.1)

	dailyDemand := holt.Level

	trend := holt.Trend

	weeklySeasonality := make([]float64, 7)

	for i := range weeklySeasonality {
		weeklySeasonality[i] = 1
	}

	if historyDays >= 60 {
		weeklySeasonality = calculateWeeklySeasonality(history)
	}

	monthlySeasonality := make([]float64, 12)

	for i := range monthlySeasonality {
		monthlySeasonality[i] = 1
	}

	if historyDays >= 365 {
		monthlySeasonality = calculateMonthlySeasonality(history)
	}

	confidence := calculateForecastConfidence(history, historyDays, weeklySeasonality, monthlySeasonality)

	response := &ForecastResponse{
		ItemID:             itemID,
		HistoricalDays:     historyDays,
		ForecastDays:       forecastDays,
		CurrentStock:       currentStock,
		DailyDemand:        dailyDemand,
		DailyDemandTrend:   trend,
		WeeklySeasonality:  weeklySeasonality,
		MonthlySeasonality: monthlySeasonality,
		HistoricalDemand:   history,
		Confidence:         confidence,
	}

	predictedInventory := currentStock

	forecast := dailyDemand

	today := time.Now()

	for i := 0; i < forecastDays; i++ {

		date := today.AddDate(0, 0, i+1)

		weekdayFactor := 1.0
		monthFactor := 1.0

		if len(weeklySeasonality) == 7 {
			weekdayFactor =
				weeklySeasonality[int(date.Weekday())]
		}

		if len(monthlySeasonality) == 12 {

			monthFactor =
				monthlySeasonality[int(date.Month())-1]
		}

		demand := forecast * weekdayFactor * monthFactor

		if demand < 0 {
			demand = 0
		}

		response.DailyForecast = append(response.DailyForecast, demand)

		rounded := int(math.Round(demand))

		response.ForecastedDemand += rounded

		predictedInventory -= rounded

		if predictedInventory < 0 {
			predictedInventory = 0
		}

		forecast += forecast * trend / float64(historyDays)
	}

	response.PredictedEndingInventory = predictedInventory

	return response, nil
}

func (s *service) GetPurchaseRecommendation(itemID uint, forecastDays int) (*PurchaseRecommendation, error) {

	item, err := s.repo.GetByID(itemID)

	if err != nil {
		return nil, err
	}

	return s.generatePurchaseRecommendation(*item, forecastDays)
}

func (s *service) GetBatchPurchaseRecommendations(forecastDays int) ([]PurchaseRecommendation, error) {

	items, err := s.repo.GetActiveItems()

	if err != nil {
		return nil, err
	}

	var recommendations []PurchaseRecommendation

	for _, item := range items {

		recommendation, err := s.generatePurchaseRecommendation(item, forecastDays)

		if err != nil {
			return nil, err
		}

		recommendations = append(recommendations, *recommendation)
	}

	sort.Slice(recommendations,
		func(i, j int) bool {
			return urgencyRank(recommendations[i].Urgency) > urgencyRank(recommendations[j].Urgency)
		},
	)

	return recommendations, nil
}

func (s *service) generatePurchaseRecommendation(item Item, forecastDays int) (*PurchaseRecommendation, error) {

	forecast, err := s.GetForecast(item.ID, 365, forecastDays)

	if err != nil {
		return nil, err
	}

	currentStock, err := s.CalculateCurrentStock(item.ID)

	if err != nil {
		return nil, err
	}

	projectedInventory := currentStock - forecast.ForecastedDemand

	requiredInventory := forecast.ForecastedDemand + item.SafetyStock

	purchase := requiredInventory - currentStock

	if purchase < 0 {
		purchase = 0
	}

	urgency := "LOW"

	reason := "Inventory level sufficient"

	if projectedInventory < 0 {

		urgency = "HIGH"

		reason = "Projected stockout during forecast period"

	} else if projectedInventory < item.MinimumStock {

		urgency = "MEDIUM"

		reason = "Inventory below minimum threshold"
	}

	return &PurchaseRecommendation{
		ItemID:              item.ID,
		ItemNumber:          item.ItemNumber,
		Description:         item.Description,
		CurrentStock:        currentStock,
		ForecastDays:        forecastDays,
		ForecastedDemand:    forecast.ForecastedDemand,
		SafetyStock:         item.SafetyStock,
		ProjectedInventory:  projectedInventory,
		RecommendedPurchase: purchase,
		Urgency:             urgency,
		Reason:              reason,
	}, nil
}

func urgencyRank(level string) int {

	switch level {

	case "HIGH":
		return 3

	case "MEDIUM":
		return 2

	default:
		return 1
	}
}
