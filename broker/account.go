package broker

import "google.golang.org/genproto/googleapis/type/decimal"

type Account struct {
	ID uint64
	CustomerID uint64
	AssetUUID string
	Balance decimal.Decimal
	LockedBalance decimal.Decimal
}