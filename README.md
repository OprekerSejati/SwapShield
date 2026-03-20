# 🛡️ SwapShield — DeFi Slippage & Price Impact Guard

![Go](https://img.shields.io/badge/Go-1.22-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-MVP-orange)

> A safety layer that prevents catastrophic swap losses in DeFi.
SwapShield is an early-stage DeFi risk engine focused on pre-trade protection.

---

## 🚨 The Problem

In AMM-based DEXs (like Uniswap), large swaps can silently suffer from **extreme price impact** due to limited liquidity.

Users often:

* Execute swaps without understanding slippage
* Lose a significant portion of their funds
* Have no clear warning before execution

Real-world incidents have shown users losing **50–90%+ of value** in a single trade.

---

## 💡 The Solution

**SwapShield** simulates swaps *before execution* and provides real-time risk analysis:

* 📊 Accurate **price impact calculation** (execution-based)
* 💧 **Liquidity usage detection**
* ⚠️ Clear **risk classification system**
* 🧠 Human-readable warnings for better decision making

---

## ⚙️ How It Works

SwapShield uses a constant product AMM model:

* Applies standard 0.3% fee (Uniswap-style)
* Calculates expected output (`amountOut`)
* Computes **execution price vs spot price**
* Derives **true price impact**
* Evaluates **liquidity consumption**

---

## 📈 Example Output

```
----------------------------------------
=== Large Trade ===
Token In: 50000.00 USDT
Amount Out: 120.500000 AAVE
Price Impact: 65.30%
Liquidity Usage: 50.00%
Risk Level: 🔴 CRITICAL
Message: You may lose more than 50% due to extreme price impact | Trade consumes a large portion of pool liquidity
🚨 WARNING: This trade is extremely unsafe!
----------------------------------------
```

---

## 🧠 Key Features

### 🔍 Pre-Trade Simulation

Simulate swaps before execution to prevent unexpected losses.

### 📊 Execution-Based Price Impact

Uses **execution price**, not just pool price — matching real DEX behavior.

### 💧 Liquidity Awareness

Detects when a trade consumes a significant portion of pool liquidity.

### ⚠️ Risk Engine

Classifies trades into:

* 🟢 LOW
* 🟡 MEDIUM
* 🟠 HIGH
* 🔴 CRITICAL

### 🧩 Modular Architecture

Clean separation of:

* AMM logic
* Simulation engine
* Risk analysis

---

## 🏗️ Project Structure

```
swapshield/
├── cmd/            # Entry point
├── internal/
│   ├── amm/        # AMM math (constant product formula)
│   ├── simulation/ # Swap simulation engine
│   ├── risk/       # Risk evaluation logic
│   ├── models/     # Data structures
│   └── dex/        # DEX data source (mock / future real integration)
```

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/OprekerSejati/SwapShield.git
cd SwapShield
```

### 2. Run the simulation

```bash
go run cmd/main.go
```

---

## 🎯 Use Cases

* 🛡️ Protect users from high-slippage trades
* 📊 Integrate into DEX frontends
* 🤖 Power trading bots with risk awareness
* 🧪 Simulate large trades safely
* 🧱 Build DeFi safety infrastructure

---

## 🤝 Contributing

Contributions are welcome!

Feel free to:
- Open issues
- Suggest improvements
- Submit pull requests

---

## 🌐 Roadmap

* [ ] Integrate real on-chain data (Uniswap pools)
* [ ] REST API for external integrations
* [ ] Frontend dashboard
* [ ] Multi-chain support
* [ ] Trade splitting optimization

---

## 🔌 Future Integrations

SwapShield is designed to evolve into an API that can be integrated into:
- DEX frontends
- Trading bots
- Portfolio dashboards

---

## 🤝 Why This Matters

DeFi is powerful — but unsafe UX leads to user losses.

SwapShield aims to:

> Make every swap **transparent, predictable, and safe**

---

## 👤 Author

GitHub: https://github.com/OprekerSejati

---

## ⭐ Support

If you find this useful, consider giving it a star ⭐
