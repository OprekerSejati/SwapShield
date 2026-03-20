package models

type Token struct {
	Address  string
	Symbol   string
	Decimals int
}

type Pool struct {
	Token0   Token
	Token1   Token
	Reserve0 float64
	Reserve1 float64
}

type SwapRequest struct {
	TokenIn  Token
	TokenOut Token
	AmountIn float64
}

type SwapResult struct {
	AmountOut      float64
	PriceImpact    float64
	LiquidityUsage float64
	WarningLevel   string
	Message        string
}
