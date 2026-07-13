package inventory

import "time"

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

func calculateTrend(values []int) float64 {

	if len(values) < 2 {
		return 0
	}

	mid := len(values) / 2

	first := calculateAverage(values[:mid])
	second := calculateAverage(values[mid:])

	// avoid huge spikes from zero demand
	if first == 0 {
		return 0
	}

	return (second - first) / first
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

		dayAverage :=
			float64(dayTotals[i]) /
				float64(dayCounts[i])

		seasonality[i] =
			dayAverage / average
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

		monthAverage :=
			float64(monthlyTotals[i]) /
				float64(monthCounts[i])

		seasonality[i] =
			monthAverage / average
	}

	return seasonality
}
