package risk

import "swapshield/internal/models"

func EvaluateRisk(result *models.SwapResult) {
	if result.WarningLevel == "ERROR" {
		return
	}

	impact := result.PriceImpact
	liquidity := result.LiquidityUsage

	// Default values
	level := "LOW"
	message := "Safe trade"

	// Step 1: Price impact based classification
	switch {
	case impact > 50:
		level = "CRITICAL"
		message = "You may lose more than 50% due to extreme price impact"
	case impact > 20:
		level = "HIGH"
		message = "High price impact detected"
	case impact > 5:
		level = "MEDIUM"
		message = "Moderate price impact"
	}

	// Step 2: Liquidity-based overrides (stronger signal)
	if liquidity > 80 {
		level = "CRITICAL"
		message = "This trade is extremely unsafe due to high liquidity consumption"
	} else if liquidity > 30 && level != "CRITICAL" {
		level = "HIGH"
	}

	// Step 3: Append liquidity insight (UX layer)
	if liquidity > 10 {
		message += " | Trade consumes a large portion of pool liquidity"
	}

	// Final assignment
	result.WarningLevel = level
	result.Message = message
}