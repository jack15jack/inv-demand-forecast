import api from "./api"

export interface ForecastConfidence {
    score:number
    level:string
    factors:{
        consistency:number
        history:number
        seasonality:number
    }
}

export interface ForecastResponse {
    itemId:number
    historicalDays:number
    forecastDays:number
    currentStock:number
    predictedEndingInventory:number
    dailyDemand:number
    dailyDemandTrend:number
    weeklySeasonality:number[]
    monthlySeasonality:number[]
    historicalDemand:number[]
    dailyForecast:number[]
    forecastedDemand:number
    confidence:ForecastConfidence
}

export async function getForecast(itemId:number, historyDays:number, forecastDays:number){

    const response = await api.get(`/items/${itemId}/forecast`,
        {
            params:{
                historyDays,
                forecastDays
            }
        }
    )

    return response.data

}