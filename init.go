package sage

import (
	"fmt"
	"github.com/Entrio/subenv"
	"github.com/WildSage-Labs/sage/chain"
	"github.com/WildSage-Labs/sage/config"
	"github.com/WildSage-Labs/sage/database"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"time"
)

func NewSage() (error, *Sage) {
	s := &Sage{}
	e := s.init()
	return e, s
}

func (s *Sage) init() error {
	var chains []chain.Chain
	if err, cfg := config.ParseConfig(subenv.Env("CONFIG_PATH", "./chains.yml")); err != nil {
		return err
	} else {
		log.Debug().Int("Chains", cfg.NumberOfChains()).Msg("Config file loaded")
		s.cfg = cfg
		for k := range cfg.Chains {
			c := cfg.Chains[k]
			currentChain := chain.Chain{
				ChainID: c.ID,
				Height:  "0",
				RPC:     c.RPC,
				LCD:     c.LCD,
			}

			for d := range c.Denoms {
				var w chain.Wallets
				denom := c.Denoms[d]

				// Add wallet to current denom list
				for wi := range c.Wallets {
					w = append(w, chain.Wallet{
						Address: c.Wallets[wi],
						Balance: "0",
						Denom:   denom,
					})
				}

				// Add denom (with wallets) to current chain
				currentChain.Denoms = append(currentChain.Denoms, chain.Denom{
					Denom:   denom,
					Wallets: w,
				})
			}
			chains = append(chains, currentChain)
			log.Trace().Int("wallets", c.NumberOfWallets()).Str("name", c.Name).Str("chain-id", c.ID).Msg("\tLoaded")
		}
	}
	s.db = &database.Store{
		Chains: chains,
	}

	s.client = &http.Client{
		Timeout: time.Second * 3,
	}

	return nil
}

func (s *Sage) Start() {
	for ci := range s.db.Chains {
		chain := &s.db.Chains[ci]
		if chain.GetNumberOfDenoms() == 0 {
			log.Warn().Str("chain-id", chain.ChainID).Msg("Chain has no denoms to monitor, skipping...")
			continue
		}
		s.wg.Add(1)
		go startChainMonitoring(chain, s)
	}
	s.wg.Wait()
}

func startChainMonitoring(chain *chain.Chain, s *Sage) {
	log.Trace().Str("chain-id", chain.ChainID).Msg("Starting chain monitoring")

	defer s.wg.Done()
	for {
		for di := range chain.Denoms {
			d := chain.Denoms[di]
			for wi := range d.Wallets {
				w := d.Wallets[wi]
				balanceUrl := fmt.Sprintf("%s/cosmos/bank/v1beta1/balances/%s/by_denom?denom=%s", chain.LCD[1], w.Address, d.Denom)
				req, err := http.NewRequest("GET", balanceUrl, nil)
				if err != nil {
					log.Warn().Str("url", balanceUrl).Msg("Failed to create request")
					continue
				}

				resp, err := s.client.Do(req)
				if err != nil {
					log.Warn().Err(err).Str("url", balanceUrl).Msg("Failed to do request")
					continue
				}

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					resp.Body.Close()
					log.Warn().Str("err", err.Error()).Msg("Failed to read body data")
					continue
				}

				fmt.Println(string(body))

				// Dont forget to close the body
				resp.Body.Close()
			}

		}
		time.Sleep(time.Second * 5)
	}
}
