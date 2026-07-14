export interface Item {
    ID: number;
    ItemNumber: string;
    Description: string;
    Category: string;
    UnitCost: number;
    UnitPrice: number;
    MinimumStock: number;
    SafetyStock: number;
    IsActive: boolean;
}

export interface InventoryTransaction {
    Id: number;
    ItemId: number;
    TransactionType: string;
    Direction: string;
    Quantity: number;
    Reference: string;
    Notes: string;
    CreatedAt: string;
}