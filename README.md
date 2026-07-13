# Inventory & Demand Forecasting System

A backend inventory management and demand forecasting platform designed to help businesses track inventory levels, analyze product demand patterns, and make data-driven purchasing decisions.

The system combines inventory tracking, transaction history, demand analytics, and forecasting algorithms to provide insights into:

- Current inventory levels
- Historical demand trends
- Future product demand
- Inventory depletion estimates
- Future purchasing recommendations

---

# Overview

Managing inventory effectively requires balancing two competing problems:

- Overstocking products ties up capital and increases storage costs
- Understocking products causes missed sales and production delays

This system aims to help purchasing teams answer questions such as:

- How much inventory do we currently have?
- How quickly are products selling?
- Which products are seasonal?
- When should we reorder?
- How much product should we purchase?

The platform is designed with a modular architecture so forecasting algorithms can evolve over time without changing the underlying inventory system.

---

# Inventory Management

The system maintains detailed information about each inventory item.

Supported item data includes:

- Item number
- Description
- Category
- Unit cost
- Unit price
- Minimum stock level
- Safety stock level
- Active/inactive status

Example:

```json
{
  "itemNumber": "ABC-123",
  "description": "Product Example",
  "category": "Hardware",
  "unitCost": 15.50,
  "unitPrice": 29.99,
  "minimumStock": 100,
  "safetyStock": 50
}
```

---

# Inventory Transactions

Inventory changes are tracked through transactions rather than directly modifying stock quantities.

This provides a complete audit history of inventory movement.

Supported transaction types:

- Sale
- Purchase
- Return
- Adjustment
- Transfer

Each transaction records:

- Item
- Quantity
- Direction
- Transaction type
- Reference information
- Notes
- Timestamp

Example:

```json
{
  "itemID": 1,
  "transactionType": "SALE",
  "direction": "OUTBOUND",
  "quantity": 25,
  "reference": "ORDER-12345"
}
```

Current inventory is calculated from transaction history:

```
Current Inventory =
Inbound Transactions -
Outbound Transactions
```

---

# Inventory Analytics

The analytics engine provides historical demand insights for individual items.

Current metrics include:

- Current inventory quantity
- Units sold within a selected timeframe
- Average daily demand
- Average weekly demand
- Days of inventory remaining
- Last sale date

Example:

```
GET /items/1/analytics?days=90
```

Response:

```json
{
  "currentStock": 425,
  "averageDailyDemand": 4.2,
  "averageWeeklyDemand": 29.4,
  "daysOfInventoryRemaining": 101
}
```

---

# Demand Forecasting

The forecasting engine predicts future inventory demand using historical transaction data.

Current forecasting pipeline:

```
Transaction History
        |
        v
Daily Demand Generation
        |
        v
Average Demand Calculation
        |
        v
Trend Detection
        |
        v
Weekly Seasonality
        |
        v
Forecast Generation
```

---

# Forecasting Features

## Historical Demand Generation

Transactions are converted into daily demand data.

Example:

```
30 Day History:

[
 5,
 0,
 3,
 8,
 0,
 10,
 4
]
```

Missing sales days are included as zero-demand days to prevent inflated forecasts.

---

## Demand Trend Detection

The forecasting system analyzes whether demand is increasing or decreasing.

Example:

```
Previous Period Average:
10 units/day

Recent Period Average:
14 units/day

Trend:
+40%
```

---

## Weekly Seasonality

The system analyzes demand differences between days of the week.

Example:

```
Monday:
1.2x average demand

Friday:
1.8x average demand

Sunday:
0.4x average demand
```

This allows forecasts to adjust for weekly purchasing patterns.

---

## Monthly / Annual Seasonality

Planned support for products with yearly demand cycles.

Examples:

- Holiday products
- Seasonal clothing
- School supplies
- Outdoor equipment

Example:

```
January:
0.5x demand

November:
2.5x demand

December:
4.0x demand
```

Monthly seasonality requires sufficient historical data, typically one or more years.

---

# Forecast Output

Example:

```
GET /items/1/forecast?historyDays=365&forecastDays=30
```

Response:

```json
{
  "itemID": 1,
  "historicalDays": 365,
  "forecastDays": 30,
  "averageDailyDemand": 8.4,
  "dailyDemandTrend": 0.12,
  "forecastedDemand": 270,
  "predictedEndingInventory": 850,
  "dailyForecast": [
    8.5,
    8.8,
    9.1
  ]
}
```

---

# Technology Stack

## Backend

- Go
- Gin Web Framework
- GORM ORM

## Database

- PostgreSQL

## Development Tools

- Git
- Docker (planned deployment support)
- SQL migrations

---

# Architecture

```
                    +----------------+
                    |    REST API    |
                    |      Gin       |
                    +-------+--------+
                            |
                            |
                    +-------v--------+
                    |    Services    |
                    +-------+--------+
                            |
                            |
                    +-------v--------+
                    |  Repository    |
                    +-------+--------+
                            |
                            |
                    +-------v--------+
                    |  PostgreSQL    |
                    +----------------+


Forecasting Pipeline:

Transactions
     |
     v
Demand History
     |
     v
Forecast Engine
     |
     v
Inventory Recommendation
```

---

# Database Design

## Items

Stores product information.

```
items
-----
id
item_number
description
category
unit_cost
unit_price
minimum_stock
safety_stock
is_active
created_at
updated_at
```

---

## Inventory Transactions

Stores inventory movement.

```
inventory_transactions
----------------------
id
item_id
transaction_type
direction
quantity
reference
notes
created_at
```

---

## Inventory Snapshots

Stores historical inventory states.

```
inventory_snapshots
-------------------
id
item_id
quantity
created_at
```

---

# Future Development

## Forecast Improvements

### Exponential Smoothing

Replace simple averages with weighted forecasting where recent demand has higher influence.

Benefits:

- Faster response to changing demand
- Better handling of trends
- Industry-standard forecasting method

---

## Forecast Accuracy Metrics

Add model evaluation:

- Mean Absolute Error (MAE)
- Mean Absolute Percentage Error (MAPE)
- Forecast bias
- Prediction confidence intervals

Example:

```
Forecast:
100 units

Expected Range:
85-120 units

Confidence:
82%
```

---

## Advanced Seasonality

Expand seasonal modeling:

- Weekly seasonality
- Monthly seasonality
- Holiday effects
- Promotions
- Product lifecycle trends

---

## Reorder Recommendations

Convert forecasts into purchasing decisions.

Planned calculation:

```
Reorder Point =
Demand During Lead Time
+
Safety Stock
```

Future recommendation example:

```
Item:
ABC-123

Current Stock:
250

Expected Demand:
300

Recommendation:
Purchase 100 units
```

---

## Dashboard

Planned frontend to make the tool more easy to use.

Planned support for:

- Charts
- Forecast visualization
- Inventory management UI

## Supplier Management

Planned support for:

- Suppliers
- Lead times
- Purchase history
- Supplier pricing
- Minimum order quantities

---

## Machine Learning Forecasting

Future models may include:

- Regression forecasting
- ARIMA
- Prophet-style models
- Neural forecasting models

Machine learning models will only be introduced after establishing strong statistical baselines.

---

## Batch Forecasting

Allow the user to get forecasts for a bunch of different items.

This could also be used to help the buyer prioritize items that are out of stock or almost out of stock.

# Project Goals

The long-term goal is to create an intelligent inventory assistant capable of:

- Tracking inventory automatically
- Predicting future demand
- Identifying seasonal products
- Preventing stockouts
- Reducing excess inventory
- Helping purchasing teams make better decisions

---

# License

This project is currently for educational and portfolio development purposes.



Phase 5 – Forecasting
- Moving average
- Exponential smoothing
- Seasonality detection
- Reorder point
- Recommended purchase quantity

Phase 6 – Dashboard
- React frontend
- Charts
- Forecast visualization
- Inventory management UI