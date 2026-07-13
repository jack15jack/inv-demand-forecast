package inventory

import (
	"math"
	"time"
)

func buildDemandHistory(transactions []InventoryTransaction, historyDays int) []int {

	dailyDemand := make(map[time.Time]int)

	now := time.Now()

	cutoff := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	).AddDate(0, 0, -historyDays+1)

	// Aggregate sales by day
	for _, t := range transactions {

		if t.TransactionType != Sale {
			continue
		}

		if t.Direction != Outbound {
			continue
		}

		if t.CreatedAt.Before(cutoff) {
			continue
		}

		day := time.Date(
			t.CreatedAt.Year(),
			t.CreatedAt.Month(),
			t.CreatedAt.Day(),
			0, 0, 0, 0,
			now.Location(),
		)

		dailyDemand[day] += t.Quantity
	}

	// Fill missing days with zeros
	history := make([]int, 0, historyDays)

	for i := historyDays - 1; i >= 0; i-- {

		day := time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			0, 0, 0, 0,
			now.Location(),
		).AddDate(0, 0, -i)

		history = append(history, dailyDemand[day])
	}

	return history
}

func calculateAverage(values []int) float64 {

	if len(values) == 0 {
		return 0
	}

	total := 0

	for _, v := range values {
		total += v
	}

	return float64(total) / float64(len(values))
}

func calculateWeeklySeasonality(history []int) []float64 {

	dayTotals := make([]int, 7)
	dayCounts := make([]int, 7)

	average := calculateAverage(history)

	if average == 0 {
		return make([]float64, 7)
	}

	now := time.Now()

	for i, demand := range history {

		date := now.AddDate(
			0,
			0,
			-(len(history) - 1 - i),
		)

		weekday := int(date.Weekday())

		dayTotals[weekday] += demand
		dayCounts[weekday]++
	}

	seasonality := make([]float64, 7)

	for i := 0; i < 7; i++ {

		if dayCounts[i] == 0 {
			seasonality[i] = 1
			continue
		}

		dayAverage := float64(dayTotals[i]) / float64(dayCounts[i])

		seasonality[i] = dayAverage / average
	}

	return seasonality
}

func calculateMonthlySeasonality(history []int) []float64 {

	monthlyTotals := make([]int, 12)
	monthCounts := make([]int, 12)

	average := calculateAverage(history)

	if average == 0 {
		return make([]float64, 12)
	}

	now := time.Now()

	for i, demand := range history {

		date := time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			0,
			0,
			0,
			0,
			now.Location(),
		).AddDate(
			0,
			0,
			-(len(history) - 1 - i),
		)

		month := int(date.Month()) - 1

		monthlyTotals[month] += demand
		monthCounts[month]++
	}

	seasonality := make([]float64, 12)

	for i := 0; i < 12; i++ {

		if monthCounts[i] == 0 {
			seasonality[i] = 1
			continue
		}

		monthAverage := float64(monthlyTotals[i]) / float64(monthCounts[i])

		seasonality[i] = monthAverage / average
	}

	return seasonality
}

func calculateDemandVariance(history []int) float64 {

	if len(history) == 0 {
		return 1
	}

	mean := calculateAverage(history)

	total := 0.0

	for _, value := range history {

		diff :=
			float64(value) - mean

		total += diff * diff
	}

	variance := total / float64(len(history))

	return math.Sqrt(variance)
}

func calculateForecastConfidence(history []int, historyDays int, weeklySeasonality []float64, monthlySeasonality []float64) ForecastConfidence {

	score := 0.0

	factors := map[string]float64{}

	// History score
	historyScore := math.Min(float64(historyDays)/365.0*100, 100)

	factors["history"] = historyScore

	score += historyScore * 0.4

	// Consistency score
	variance := calculateDemandVariance(history)

	consistency := 100 - math.Min(variance*5, 100)

	factors["consistency"] = consistency

	score += consistency * 0.4

	// Seasonality score
	seasonalityScore := 0.0

	if historyDays >= 60 && len(weeklySeasonality) == 7 {
		seasonalityScore += 50
	}

	if historyDays >= 365 && len(monthlySeasonality) == 12 {
		seasonalityScore += 50
	}

	factors["seasonality"] =
		seasonalityScore

	score += seasonalityScore * 0.2

	level := "LOW"

	if score >= 75 {
		level = "HIGH"
	} else if score >= 50 {
		level = "MEDIUM"
	}

	return ForecastConfidence{
		Score:   math.Round(score),
		Level:   level,
		Factors: factors,
	}
}

func holtLinearForecast(history []int, alpha float64, beta float64) HoltForecast {

	if len(history) < 2 {
		return HoltForecast{}
	}

	level := float64(history[0])

	trend := float64(history[1]) - float64(history[0])

	for i := 1; i < len(history); i++ {

		value := float64(history[i])

		previousLevel := level

		level = alpha*value + (1-alpha)*(level+trend)

		trend = beta*(level-previousLevel) + (1-beta)*trend

	}

	return HoltForecast{
		Level: level,
		Trend: trend,
	}
}
