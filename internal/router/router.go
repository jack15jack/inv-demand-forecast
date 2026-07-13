package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jack15jack/inv-demand-forecast/internal/inventory"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Inventory wiring
	repo := inventory.NewRepository(db)
	service := inventory.NewService(repo)
	handler := inventory.NewHandler(service)

	items := r.Group("/items")
	{
		items.POST("", handler.CreateItem)
		items.GET("", handler.GetItems)
		items.GET("/:id", handler.GetItem)

		items.GET("/:id/transactions", handler.GetTransactions)
		items.GET("/:id/stock", handler.GetStock)

		items.POST("/:id/snapshots", handler.CreateSnapshot)
		items.GET("/:id/snapshots", handler.GetSnapshots)

		items.GET("/:id/analytics", handler.GetAnalytics)
		items.GET("/:id/forecast", handler.GetForecast)

		items.GET("/:id/recommendation", handler.GetPurchaseRecommendation)
	}

	transactions := r.Group("/transactions")
	{
		transactions.POST("", handler.CreateTransaction)
	}

	return r
}
