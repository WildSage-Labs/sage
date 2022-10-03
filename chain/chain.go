package chain

type (
	Chain struct {
		ChainID string `json:"chain_id"`
		Height  string `json:"height"`
	}
)

func (c Chain) GetLatestBlockID() uint32 {
	return 0
}
