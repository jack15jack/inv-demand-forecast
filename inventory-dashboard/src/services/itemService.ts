import api from "./api"

export interface CreateItemRequest {
    itemNumber: string
    description: string
    category: string
    unitCost: number
    unitPrice: number
    minimumStock: number
    safetyStock: number
    isActive: boolean
}

export interface Item {
    id: number
    itemNumber: string
    description: string
    category: string
    unitCost: number
    unitPrice: number
    minimumStock: number
    safetyStock: number
    isActive: boolean
    createdAt: string
    updatedAt: string
}

export async function createItem(item: CreateItemRequest): Promise<Item> {

    const response = await api.post("/items", item)

    return response.data
}