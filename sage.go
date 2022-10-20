package sage

import (
	"github.com/WildSage-Labs/sage/config"
	"github.com/WildSage-Labs/sage/database"
	"net/http"
	"sync"
)

type (
	Sage struct {
		wg     sync.WaitGroup
		cfg    config.Config
		db     *database.Store
		client *http.Client
	}
)
