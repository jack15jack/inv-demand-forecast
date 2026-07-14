import api from "./api"

export async function getItems(){

    const response = await api.get(
        "/items"
    )

    return response.data

}

export async function getPurchaseRecommendations(){

    const response = await api.get(
        "/recommendations"
    )

    return response.data

}

export async function getTransactions(){

    const response = await api.get(
        "/transactions"
    )

    return response.data

}