package main

import (
	"fmt"
	"log"
	"os"

	"swapshield/api"
	"swapshield/internal/dex"
	"swapshield/internal/models"
	"swapshield/internal/risk"
	"swapshield/internal/simulation"
)

func riskIndicator(level string) string {
	switch level {
	case "CRITICAL":
		return "🔴"
	case "HIGH":
		return "🟠"
	case "MEDIUM":
		return "🟡"
	case "LOW":
		return "🟢"
	default:
		return "⚪"
	}
}

func runScenario(title string, poolAmountIn float64) {
	pool := dex.GetMockPool()
	req := models.SwapRequest{
		TokenIn:  pool.Token0,
		TokenOut: pool.Token1,
		AmountIn: poolAmountIn,
	}

	result := simulation.SimulateSwap(pool, req)
	risk.EvaluateRisk(&result)

	indicator := riskIndicator(result.WarningLevel)

	fmt.Println("----------------------------------------")
	fmt.Printf("=== %s ===\n", title)
	fmt.Printf("Token In: %.2f %s\n", req.AmountIn, req.TokenIn.Symbol)
	fmt.Printf("Amount Out: %.6f %s\n", result.AmountOut, req.TokenOut.Symbol)
	fmt.Printf("Price Impact: %.2f%%\n", result.PriceImpact)
	fmt.Printf("Liquidity Usage: %.2f%%\n", result.LiquidityUsage)
	fmt.Printf("Risk Level: %s %s\n", indicator, result.WarningLevel)
	fmt.Printf("Message: %s\n", result.Message)
	if result.WarningLevel == "CRITICAL" {
		fmt.Println("🚨 WARNING: This trade is extremely unsafe!")
	}
	fmt.Println("----------------------------------------")
	fmt.Println()
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "api" {
		fmt.Println("Starting SwapShield API on :8080")
		fmt.Println("Endpoint: POST /simulate-swap")
		if err := api.StartServer(":8080"); err != nil {
			log.Fatalf("failed to start API server: %v", err)
		}
		return
	}

	runScenario("Small Trade", 1000)
	runScenario("Large Trade", 50000)
	runScenario("Extreme Trade", 90000)
}
