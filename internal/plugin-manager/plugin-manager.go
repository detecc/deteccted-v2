package plugin_manager

import (
	"fmt"
	deteccPlugin "github.com/detecc/deteccted-v2/internal/models/plugin"
	log "github.com/sirupsen/logrus"
	"plugin"
	"sync"
)

var pluginManager Manager

func init() {
	once := sync.Once{}
	once.Do(func() {
		GetPluginManager()
	})
}

type (
	// ManagerImpl is a manager for the plugins. It stores and maps the plugins to the command.
	ManagerImpl struct {
		plugins sync.Map
	}

	Manager interface {
		HasPlugin(name string) bool
		AddPlugin(name string, plugin deteccPlugin.Handler)
		GetPlugin(name string) (deteccPlugin.Handler, error)
		GetPlugins() []deteccPlugin.Handler
		LoadPlugins(pluginDir string, plugins []string)
	}
)

// Register a cmd to the manager.
func Register(name string, plugin deteccPlugin.Handler) {
	GetPluginManager().AddPlugin(name, plugin)
}

// GetPluginManager gets the cmd manager instance (singleton).
func GetPluginManager() Manager {
	if pluginManager == nil {
		pluginManager = &ManagerImpl{plugins: sync.Map{}}
	}

	return pluginManager
}

// HasPlugin Check if the cmd exists in the manager.
func (pm *ManagerImpl) HasPlugin(name string) bool {
	_, exists := pm.plugins.Load(name)
	return exists
}

// AddPlugin Add a cmd to the manager.
func (pm *ManagerImpl) AddPlugin(name string, plugin deteccPlugin.Handler) {
	log.Debugf("Adding cmd-manager to manager %s", name)
	if !pm.HasPlugin(name) {
		pm.plugins.Store(name, plugin)
	}
}

// GetPlugin returns the cmd if found.
func (pm *ManagerImpl) GetPlugin(name string) (deteccPlugin.Handler, error) {
	mPlugin, exists := pm.plugins.Load(name)
	if exists {
		return mPlugin.(deteccPlugin.Handler), nil
	}

	return nil, fmt.Errorf("cmd-manager doesnt exist")
}

// GetPlugins returns the plugins in the manager
func (pm *ManagerImpl) GetPlugins() []deteccPlugin.Handler {
	var plugins []deteccPlugin.Handler

	pm.plugins.Range(func(key, value interface{}) bool {
		pluginHandler, canCast := value.(deteccPlugin.Handler)
		if canCast {
			plugins = append(plugins, pluginHandler)
		}

		return true
	})

	return plugins
}

// LoadPlugins Load the plugins from the folder, specified in the configuration file.
func (pm *ManagerImpl) LoadPlugins(pluginDir string, plugins []string) {
	log.Info("Loading plugins..")

	for _, pluginFromList := range plugins {
		log.Debugf("Loading cmd: %s", pluginFromList)

		_, err := plugin.Open(fmt.Sprintf("%s/%s.so", pluginDir, pluginFromList))
		if err != nil {
			log.WithError(err).Errorf("Error loading cmd to manager")
			continue
		}
	}
}
