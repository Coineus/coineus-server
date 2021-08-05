package model

type Wallet struct {
	Id     int    `json:"id"`
	UserId int    `json:"userid"`
	Name   string `json:"name"`
}

type WalletOperation struct {
	Id          int `json:"id"`
	WalletId    int `json:"walletid"`
	OperationId int `json:"operationid"`
}

type WalletOperationDTO struct {
	Id          int     `json:"id"`
	WalletId    int     `json:"walletid"`
	OperationId int     `json:"operationid"`
	UserId      int     `json:"userid"`
	CoinSymbol  string  `json:"coinsymbol"`
	CoinAmount  float32 `json:"coinamount"`
	BuyCost     float32 `json:"buycost"`
}
