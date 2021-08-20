package model

import "time"

type RecentOperation struct {
	Id         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UserId     string    `json:"userid"`
	CoinSymbol string    `json:"coinsymbol"`
	CoinAmount float32   `json:"coinamount"`
	BuyCost    float32   `json:"buycost"`
}

type ArchivedOperation struct {
	Id         string    `json:"id"`
	UserId     string    `json:"userid"`
	CreatedAt  time.Time `json:"created_at"`
	ArchivedAt time.Time `json:"archived_at"`
	CoinSymbol string    `json:"coinsymbol"`
	CoinAmount float32   `json:"coinamount"`
	BuyCost    float32   `json:"buycost"`
	SellCost   float32   `json:"sellcost"`
}
