package simulation

import (
	"swapshield/internal/amm"
	"swapshield/internal/models"
)

func SimulateSwap(pool models.Pool, req models.SwapRequest) models.SwapResult {
	if req.AmountIn <= 0 {
		return models.SwapResult{
			AmountOut:      0,
			PriceImpact:    0,
			LiquidityUsage: 0,
			WarningLevel:   "ERROR",
			Message:        "Invalid swap amount",
		}
	}

	var reserveIn, reserveOut float64

	if req.TokenIn.Address == pool.Token0.Address {
		reserveIn = pool.Reserve0
		reserveOut = pool.Reserve1
	} else if req.TokenIn.Address == pool.Token1.Address {
		reserveIn = pool.Reserve1
		reserveOut = pool.Reserve0
	} else {
		return models.SwapResult{
			AmountOut:      0,
			PriceImpact:    0,
			LiquidityUsage: 0,
			WarningLevel:   "ERROR",
			Message:        "Invalid token pair",
		}
	}

	amountOut := amm.GetAmountOut(req.AmountIn, reserveIn, reserveOut)
	priceImpact := amm.GetPriceImpact(req.AmountIn, reserveIn, reserveOut)
	liquidityUsage := (req.AmountIn / reserveIn) * 100
	if liquidityUsage > 100 {
		liquidityUsage = 100
	}

	return models.SwapResult{
		AmountOut:      amountOut,
		PriceImpact:    priceImpact,
		LiquidityUsage: liquidityUsage,
	}
}
