package inventory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateItem(c *gin.Context) {
	var req CreateItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.service.CreateItem(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *Handler) GetItems(c *gin.Context) {
	items, err := h.service.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) GetItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	item, err := h.service.GetItem(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) CreateTransaction(c *gin.Context) {
	var req CreateTransactionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	transaction, err := h.service.CreateTransaction(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func (h *Handler) GetTransactions(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid item id",
		})
		return
	}

	transactions, err := h.service.GetTransactions(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetStock(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid item id",
		})
		return
	}

	stock, err := h.service.GetStock(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, stock)
}

func (h *Handler) CreateSnapshot(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid item id",
		})
		return
	}

	snapshot, err := h.service.CreateSnapshot(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, snapshot)
}

func (h *Handler) GetSnapshots(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid item id",
		})
		return
	}

	snapshots, err := h.service.GetSnapshots(uint(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, snapshots)
}

func (h *Handler) GetAnalytics(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid item id",
		})
		return
	}

	days := 30

	if daysStr := c.Query("days"); daysStr != "" {

		d, err := strconv.Atoi(daysStr)

		if err != nil || d <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "days must be a positive integer",
			})
			return
		}

		days = d
	}

	analytics, err := h.service.GetAnalytics(uint(id), days)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, analytics)
}

func (h *Handler) GetForecast(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid item id",
		})
		return
	}

	historyDays := 30
	forecastDays := 7

	if value := c.Query("historyDays"); value != "" {

		historyDays, err = strconv.Atoi(value)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid historyDays",
			})
			return
		}
	}

	if value := c.Query("forecastDays"); value != "" {

		forecastDays, err = strconv.Atoi(value)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "invalid forecastDays",
			})
			return
		}
	}

	forecast, err := h.service.GetForecast(uint(id), historyDays, forecastDays)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, forecast)
}

func (h *Handler) GetPurchaseRecommendation(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {

		c.JSON(400, gin.H{
			"error": "invalid item id",
		})

		return
	}

	forecastDays := 30

	if days := c.Query("days"); days != "" {

		d, err := strconv.Atoi(days)

		if err != nil || d <= 0 {

			c.JSON(400, gin.H{
				"error": "days must be positive",
			})

			return
		}

		forecastDays = d
	}

	result, err := h.service.GetPurchaseRecommendation(uint(id), forecastDays)

	if err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, result)
}
