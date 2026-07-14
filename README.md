# Inventory & Demand Forecasting System

A full-stack inventory intelligence platform designed to help businesses track inventory, analyze demand patterns, forecast future requirements, and make data-driven purchasing decisions.

The system combines:

- Inventory management
- Transaction-based stock tracking
- Demand analytics
- Statistical forecasting
- Inventory depletion prediction
- Automated purchasing recommendations

The goal is to transform raw inventory data into actionable decisions:

> "What do we have?"
>
> "What will we need?"
>
> "What should we buy?"

---

# Overview

Inventory management requires balancing two competing problems:

- Overstocking ties up capital and increases storage costs
- Understocking causes missed sales and production delays

This system helps purchasing teams answer:

- How much inventory is currently available?
- How quickly are products being consumed?
- Are products seasonal?
- When will inventory run out?
- How much should be purchased?

The platform uses a modular backend architecture so forecasting algorithms can evolve independently from the inventory system.

---

# Current Features

## Inventory Management

Each inventory item contains:

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

# Transaction-Based Inventory

Inventory quantities are not directly modified.

Instead, all inventory movement is recorded as transactions.

Supported transaction types:

- Sale
- Purchase
- Return
- Adjustment
- Transfer

Each transaction stores:

- Item
- Quantity
- Direction
- Transaction type
- Reference
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

Current inventory:

```
Inventory =
Inbound Transactions -
Outbound Transactions
```

This provides a complete audit history of inventory movement.

---

# Analytics

The analytics engine provides historical demand insights.

Current metrics:

- Current stock
- Units sold
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
  "currentStock":425,
  "averageDailyDemand":4.2,
  "averageWeeklyDemand":29.4,
  "daysOfInventoryRemaining":101
}
```

---

# Demand Forecasting

The forecasting engine converts transaction history into future demand predictions.

Pipeline:

```
Transactions
      |
      v
Daily Demand History
      |
      v
Holt Linear Smoothing
      |
      v
Seasonality Adjustment
      |
      v
Forecast Generation
```

---

# Forecasting Methods

## Holt Linear Exponential Smoothing

The forecasting engine uses Holt's Linear Trend Method, an exponential smoothing technique that models both the current demand level and the underlying demand trend.

Unlike simple exponential smoothing, Holt's method can detect increasing or decreasing demand patterns.

The model maintains:

- Level component
- Trend component

Benefits:

- Recent demand has higher influence
- Adapts to changing demand patterns
- Handles gradual growth or decline
- Provides a stronger baseline than moving averages

---

## Weekly Seasonality

The system analyzes demand differences between weekdays.

Example:

```
Monday:
1.2x average demand

Friday:
1.8x average demand

Sunday:
0.4x average demand
```

Seasonality is only applied when enough historical data exists to avoid overfitting.

---

## Monthly Seasonality

Support for yearly demand cycles:

Examples:

- Holiday products
- Seasonal equipment
- School supplies
- Clothing

Example:

```
January:
0.5x demand

November:
2.5x demand

December:
4.0x demand
```

---

# Forecast Confidence

Forecasts include a confidence estimate based on:

- Amount of historical data
- Demand consistency
- Available seasonality information

Example:

```json
{
  "confidence":{
    "score":82,
    "level":"HIGH"
  }
}
```

This allows users to understand when predictions should be trusted.

---

# Purchasing Recommendations

Forecasts are converted into actionable purchasing decisions.

Instead of only predicting demand:

```
Expected Demand:
300 units
```

the system provides:

```
Recommended Purchase:
250 units
```

Calculation:

```
Required Inventory =
Forecasted Demand
+
Safety Stock

Recommended Purchase =
Required Inventory
-
Current Stock
```

Example:

```json
{
  "itemID":1,
  "currentStock":120,
  "forecastedDemand":300,
  "safetyStock":100,
  "recommendedPurchase":280,
  "urgency":"HIGH"
}
```

---

# API Endpoints

## Items

```
POST   /items
GET    /items
GET    /items/:id
```

## Transactions

```
POST   /transactions
GET    /items/:id/transactions
```

## Stock
```
GET    /items/:id/stock
```

## Snapshots

```
POST   /items/:id/snapshots
GET    /items/:id/snapshots
```

## Analytics

```
GET /items/:id/analytics
```

Parameters:

```
days
```

## Forecasting

```
GET /items/:id/forecast
```

Parameters:

```
historyDays
forecastDays
```

## Purchasing Recommendation

```
GET /items/:id/recommendation
```

---

# Technology Stack

## Frontend

- React
- Vite
- Material UI
- React Router
- TanStack Query
- Axios
- Recharts
- ESLint
- Prettier

## Backend

- Go
- Gin Web Framework
- GORM ORM

## Database

- PostgreSQL
- SQL migrations

## Development Tools

- Git
- Docker support planned
- Data generation tooling

---

# Architecture

```
                 REST API
                    |
                    |
                Controllers
                    |
                    |
                Services
                    |
                    |
              Repositories
                    |
                    |
               PostgreSQL


Forecasting System:

Transactions
      |
      v
Demand History
      |
      v
Forecast Engine
      |
      v
Purchase Recommendation
```

---

# Database Design

## Items

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

# Future Development

## Dashboard Application

Frontend dashboard for interacting with the inventory system.

Planned features:

- Inventory overview
- Demand charts
- Forecast visualization
- Purchasing recommendations
- Stock alerts
- Item management

Technology:

- React
- Charting library
- REST API integration

---

## Machine Learning Forecasting

Future forecasting models may include:

- Regression models
- ARIMA
- Prophet-style forecasting
- Gradient boosted forecasting
- Neural forecasting models

The plan is to maintain statistical forecasting methods as a baseline before introducing machine learning models.

---

## Batch Forecasting

Support forecasting across all inventory items.

Example:

```
Generate recommendations for all items

↓

Rank by urgency

↓

Prioritize purchasing decisions
```

---

## Inventory Alerts

Future support:

- Stockout warnings
- Overstock detection
- Slow-moving inventory
- Dead inventory identification

---

# Data Generator

The project includes tools for generating inventory histories.

This allows forecasting algorithms to be tested with realistic scenarios.

---

# Long-Term Goals

The goal is to create an intelligent inventory assistant capable of:

- Tracking inventory automatically
- Predicting future demand
- Identifying seasonal behavior
- Preventing stockouts
- Reducing excess inventory
- Recommending purchasing decisions

---

# License

This project is currently developed for educational and portfolio purposes.