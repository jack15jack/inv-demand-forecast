package inventory

type PurchaseRecommendation struct {
	ItemID              uint   `json:"itemId"`
	ItemNumber          string `json:"itemNumber"`
	Description         string `json:"description"`
	CurrentStock        int    `json:"currentStock"`
	ForecastDays        int    `json:"forecastDays"`
	ForecastedDemand    int    `json:"forecastedDemand"`
	SafetyStock         int    `json:"safetyStock"`
	ProjectedInventory  int    `json:"projectedInventory"`
	RecommendedPurchase int    `json:"recommendedPurchase"`
	Urgency             string `json:"urgency"`
	Reason              string `json:"reason"`
}
