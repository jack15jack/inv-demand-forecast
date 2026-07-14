import api from "./api"

export interface PurchaseRecommendation {
    itemID: number
    itemNumber: string
    description: string
    currentStock: number
    forecastDays: number
    forecastedDemand: number
    safetyStock: number
    projectedInventory: number
    recommendedPurchase: number
    urgency: string
    reason: string
}

export async function getRecommendations(){

    const response = await api.get("/recommendations")

    return response.data

}