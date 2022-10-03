package config

type (
	chains map[string][]ChainConfig
	Config struct {
		Chains map[string]chains `yaml:"chains"`
	}

	ChainConfig struct {
		RPC string `yaml:"rpc"`
		LCD string `yaml:"lcd"`
	}
)
