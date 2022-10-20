package chain

type (
	Wallet struct {
		Address string
		Balance string
		Denom   string
	}
	Wallets []Wallet

	Denom struct {
		Denom   string
		Wallets Wallets
	}
	Denoms []Denom

	Chain struct {
		ChainID string   `json:"chain_id"`
		Height  string   `json:"height"`
		RPC     []string `json:"rpc"`
		LCD     []string `json:"lcd"`
		Denoms  Denoms   `json:"denoms"`
	}
)

func (c Chain) GetLatestBlockID() uint32 {
	return 0
}

func (c Chain) GetNumberOfDenoms() int {
	return len(c.Denoms)
}

func (c Chain) UpdateDenomWalletBalance(denom, wallet, balance string) error {
	return nil
}
