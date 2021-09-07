package model

type Wallet struct {
	Id     string `json:"id"`
	UserId int    `json:"userid"`
	Name   string `json:"name"`
}

type WalletOperation struct {
	Id          string `json:"id"`
	WalletId    string `json:"walletid"`
	OperationId string `json:"operationid"`
}

type WalletOperationDTO struct {
	Id          string  `json:"id"`
	WalletId    string  `json:"walletid"`
	OperationId string  `json:"operationid"`
	UserId      string  `json:"userid"`
	CoinSymbol  string  `json:"coinsymbol"`
	CoinAmount  float32 `json:"coinamount"`
	BuyCost     float32 `json:"buycost"`
}
