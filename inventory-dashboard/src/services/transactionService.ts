import api from "./api"


export interface CreateTransactionRequest {
    ItemID:number
    TransactionType:string
    Direction:string
    Quantity:number
    Reference:string
    Notes:string
}

export async function createTransaction(transaction:CreateTransactionRequest){

    const response = await api.post("/transactions", transaction)

    return response.data
}