package storage

type Store interface {
	Users() UserStore
	Wallets() WalletStore
	WalletOps() WalletOperationStore
	RecentOps() RecentOperationStore
	ArchivedOps() ArchivedOperationStore
}
