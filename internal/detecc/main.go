package detecc

import (
	"context"
	"github.com/detecc/deteccted-v2/internal/models/config"
	"github.com/detecc/deteccted-v2/internal/pkg/logging"
	"github.com/detecc/deteccted-v2/internal/pkg/mqtt"
	pluginManager "github.com/detecc/deteccted-v2/internal/plugin-manager"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(isDebug bool, cfg *config.Configuration) {
	var (
		quitChannel  = make(chan os.Signal, 1)
		clientConfig = cfg.Client
		manager      = pluginManager.GetPluginManager()
		ctx, cancel  = context.WithCancel(context.Background())
	)

	// Capture the terminate signal
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)

	// Setup logger
	logging.Setup(log.StandardLogger(), cfg.Logging, isDebug)

	// Load plugins
	manager.LoadPlugins(clientConfig.PluginDir, clientConfig.Plugins)

	// Create new MQTT client (message bus)
	mqttClient := mqtt.NewMqttClient(clientConfig.ServiceNodeIdentifier, cfg.MqttBroker)

	// Create a client
	client := NewClient(manager, mqttClient, clientConfig.ServiceNodeIdentifier)

	// Register the client with management service
	client.Register(clientConfig.AuthPassword)

	// Send heartbeat periodically
	go client.SendHeartbeat(ctx, time.Minute)

ExitListener:
	for {
		select {
		case <-quitChannel:
			log.Info("Exiting..")
			cancel()
		case <-ctx.Done():
			cancel()
			break ExitListener
		}
	}
}
