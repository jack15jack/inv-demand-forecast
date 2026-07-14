import api from "./api"


export interface Item {
    id:number
    itemNumber:string
    description:string
    category:string
    unitCost:number
    unitPrice:number
    minimumStock:number
    safetyStock:number
    isActive:boolean
}



export async function getItems(){

    const response = await api.get("/items")

    return response.data

}