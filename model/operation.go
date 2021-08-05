package model

type RecentBuyOperation struct {
	Id         int     `json:"id"`
	UserId     int     `json:"userid"`
	CoinSymbol string  `json:"coinsymbol"`
	CoinAmount float32 `json:"coinamount"`
	BuyCost    float32 `json:"buycost"`
}

type ArchivedOperation struct {
	Id         int     `json:"id"`
	UserId     int     `json:"userid"`
	CoinSymbol string  `json:"coinsymbol"`
	CoinAmount float32 `json:"coinamount"`
	BuyCost    float32 `json:"buycost"`
	SellCost   float32 `json:"sellcost"`
}
