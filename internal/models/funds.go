package models

import (
	"time"

	"github.com/shopspring/decimal" // Библиотека для точных вычислений, т.к. float64 имеет погрешности
)

type Funds struct {
	Id               int    `json:"id"`
	Name             string    `json:"name"`
	Ticker           string    `json:"ticker"`
	Amount           decimal.Decimal    `json:"amount"`
	PricePerItem     decimal.Decimal   `json:"pricePerItem"`
	PurchasePrice    decimal.Decimal   `json:"purchasePrice"`
	PriceCurrent     decimal.Decimal   `json:"priceCurrent"`
	PercentChanges   decimal.Decimal   `json:"percentChanges"`
	YearlyInvestment decimal.Decimal   `json:"yearlyInvestment"`
	ClearMoney       decimal.Decimal   `json:"clearMoney"`
	DatePurchase     time.Time `json:"datePurchase"`
	DateLastUpdate time.Time `json:"dateLastUpdate"`
	Type string     `json:"type"`
}