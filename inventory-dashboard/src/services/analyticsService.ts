import api from "./api"

export interface AnalyticsResponse {
    itemID: number
	analysisWindowDays: number
	currentStock:number
	averageDailyDemand:number
	averageWeeklyDemand:number
	daysOfInventoryRemaining:number
	unitsSold:number
	lastSale: string
}

export async function getAnalytics(itemID: number, days: number): Promise<AnalyticsResponse> {

    const response = await api.get(
        `/items/${itemID}/analytics`,
        {
            params: { days }
        }
    )

    return response.data
}