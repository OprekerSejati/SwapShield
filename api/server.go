package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"swapshield/internal/dex"
	"swapshield/internal/models"
	"swapshield/internal/risk"
	"swapshield/internal/simulation"
)

// SimulateSwapHandler handles POST /simulate-swap requests.
func SimulateSwapHandler(w http.ResponseWriter, r *http.Request) {
	// Optional CORS support for frontend integration.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.SwapRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	pool := dex.GetMockPool()
	result := simulation.SimulateSwap(pool, req)
	risk.EvaluateRisk(&result)

	w.Header().Set("Content-Type", "application/json")
	if result.WarningLevel == "ERROR" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}

// StartServer starts the HTTP API server on the provided address.
func StartServer(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/simulate-swap", SimulateSwapHandler)
	return http.ListenAndServe(addr, mux)
}
