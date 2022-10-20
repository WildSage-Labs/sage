package config

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	chains []ChainConfig
	Config struct {
		Chains chains `yaml:"chains"`
	}

	ChainConfig struct {
		Name    string   `yaml:"name"`
		ID      string   `yaml:"chain-id"`
		RPC     []string `yaml:"rpc"`
		LCD     []string `yaml:"lcd"`
		Wallets []string `yaml:"wallets"`
		Denoms  []string `yaml:"denoms"`
	}
)

func ParseConfig(filePath string) (error, Config) {
	log.Debug().Str("filepath", filePath).Msg("Opening config file")
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err, Config{}
	}
	conf := Config{}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return err, Config{}
	}

	// Do some basic config validation
	for k := range conf.Chains {
		c := conf.Chains[k]
		if len(c.Denoms) == 0 {
			log.Warn().Str("chain-id", c.ID).Msg("Chain has no denoms to query. Have you forgot to add some?")
		}
	}

	return nil, conf
}

func (cfg Config) NumberOfChains() int {
	return len(cfg.Chains)
}

func (c ChainConfig) NumberOfWallets() int {
	return len(c.Wallets)
}
