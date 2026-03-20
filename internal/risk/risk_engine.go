package risk

import "swapshield/internal/models"

func EvaluateRisk(result *models.SwapResult) {
	if result.WarningLevel == "ERROR" {
		return
	}

	impact := result.PriceImpact
	baseMessage := ""

	switch {
	case impact > 50:
		result.WarningLevel = "CRITICAL"
		baseMessage = "You may lose more than 50% due to extreme price impact"
	case impact > 20:
		result.WarningLevel = "HIGH"
		baseMessage = "High price impact detected"
	case impact > 5:
		result.WarningLevel = "MEDIUM"
		baseMessage = "Moderate price impact"
	default:
		result.WarningLevel = "LOW"
		baseMessage = "Safe trade"
	}

	if result.LiquidityUsage > 30 && result.WarningLevel != "CRITICAL" {
		result.WarningLevel = "HIGH"
	}

	result.Message = baseMessage

	if result.LiquidityUsage > 10 {
		result.Message += " | Trade consumes a large portion of pool liquidity"
	}
}
