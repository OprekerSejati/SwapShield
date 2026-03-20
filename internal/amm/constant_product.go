package amm

func GetAmountOut(amountIn, reserveIn, reserveOut float64) float64 {
	if amountIn <= 0 || reserveIn <= 0 || reserveOut <= 0 {
		return 0
	}

	// 0.3% fee (Uniswap v2 style)
	amountInWithFee := amountIn * 0.997
	numerator := amountInWithFee * reserveOut
	denominator := reserveIn + amountInWithFee

	return numerator / denominator
}

func GetPriceImpact(amountIn, reserveIn, reserveOut float64) float64 {
	if amountIn <= 0 || reserveIn <= 0 || reserveOut <= 0 {
		return 0
	}

	amountOut := GetAmountOut(amountIn, reserveIn, reserveOut)
	if amountOut <= 0 {
		return 0
	}

	spotPrice := reserveOut / reserveIn
	if spotPrice == 0 {
		return 0
	}
	executionPrice := amountOut / amountIn

	impact := (spotPrice - executionPrice) / spotPrice * 100
	if impact < 0 {
		return 0
	}

	return impact
}
