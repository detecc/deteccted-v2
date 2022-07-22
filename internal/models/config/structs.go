package config

const (
	ConfigurationKey         = "configuration"
	ConfigurationFilePathKey = "configurationFilePath"
)

type (
	MqttBroker struct {
		Host     string `fig:"host" validate:"required"`
		Port     int    `fig:"port" validate:"required"`
		Username string `fig:"username" validate:"required"`
		Password string `fig:"password" validate:"required"`
		TLS      TLS    `fig:"tls"`
	}

	TLS struct {
		IsEnabled bool   `fig:"isEnabled"`
		CertPath  string `fig:"certPath"`
		KeyPath   string `fig:"keyPath"`
	}

	Logging struct {
		Type    []string `fig:"type" validate:"required"` // file, remote, console
		Format  string   `fig:"format" default:"syslog"`  // syslog, json, etc
		Address string   `fig:"address" default:"localhost:514"`
	}

	Client struct {
		ServiceNodeIdentifier string   `fig:"serviceNodeId" validate:"required"`
		AuthPassword          string   `fig:"authPassword" validate:"required"`
		PluginDir             string   `fig:"pluginDir"`
		Plugins               []string `fig:"plugins"`
	}

	Configuration struct {
		Client     Client     `fig:"client" validate:"required"`
		MqttBroker MqttBroker `fig:"mqtt" validate:"required"`
		Logging    Logging    `fig:"logging" validate:"required"`
	}
)
