package dex

import "swapshield/internal/models"

// GetMockPool returns hardcoded pool reserves for MVP testing.
func GetMockPool() models.Pool {
	return models.Pool{
		Token0: models.Token{
			Address:  "0xTokenA",
			Symbol:   "USDT",
			Decimals: 18,
		},
		Token1: models.Token{
			Address:  "0xTokenB",
			Symbol:   "AAVE",
			Decimals: 18,
		},
		Reserve0: 100000, // USDT
		Reserve1: 500,    // AAVE
	}
}
