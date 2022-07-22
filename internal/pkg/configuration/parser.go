package configuration

import (
	"github.com/detecc/deteccted-v2/internal/cache"
	"github.com/detecc/deteccted-v2/internal/models/config"
	"github.com/kkyr/fig"
	goCache "github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"path/filepath"
)

func getConfiguration(config interface{}, cache *goCache.Cache, filePath, cacheKey string) {
	var (
		err error
	)

	// Load the config
	err = fig.Load(config,
		fig.File(filepath.Base(filePath)),
		fig.Dirs(filepath.Dir(filePath), "."),
	)
	if err != nil {
		log.WithError(err).Fatal("Unable to load configuration")
	}

	if cache != nil {
		// Cache the config
		cache.Set(cacheKey, &config, goCache.NoExpiration)
	}
}

// GetClientConfiguration get the configuration from the configuration file and store the configuration in the cache
func GetClientConfiguration(configurationFilePath string) *config.Configuration {
	var (
		cfg                 config.Configuration
		memory              = cache.Memory()
		isFound             bool
		cachedConfiguration interface{}
	)

	cachedConfiguration, isFound = memory.Get(config.ConfigurationKey)
	if isFound {
		return cachedConfiguration.(*config.Configuration)
	}

	getConfiguration(&cfg, memory, configurationFilePath, config.ConfigurationKey)

	return &cfg
}
