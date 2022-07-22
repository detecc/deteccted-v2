package detecc

import (
	"context"
	"fmt"
	"github.com/detecc/deteccted-v2/internal/pkg/mqtt"
	pluginManager "github.com/detecc/deteccted-v2/internal/plugin-manager"
	"github.com/detecc/detecctor-v2/pkg/payload"
	log "github.com/sirupsen/logrus"
	"time"
)

type Client struct {
	id            string
	pluginManager pluginManager.Manager
	mqttClient    mqtt.Client
}

func NewClient(pluginManager pluginManager.Manager, mqttClient mqtt.Client, id string) *Client {
	// Setup topics
	var (
		clientListenTopic        = mqtt.Topic(fmt.Sprintf("client/%s/cmd-manager/#", id))
		clientRegisterReplyTopic = mqtt.Topic(fmt.Sprintf("client/%s/register/reply", id))
		client                   = &Client{
			pluginManager: pluginManager,
			id:            id,
			mqttClient:    mqttClient,
		}
	)

	// Subscribe to the topics
	mqttClient.Subscribe(clientListenTopic, client.PluginHandler())
	mqttClient.Subscribe(clientRegisterReplyTopic, client.AuthHandler())

	return client
}

func (c *Client) PluginHandler() mqtt.MessageHandler {
	return func(client mqtt.Client, topicIds []string, payloadId uint16, payload interface{}, err error) {

	}
}

func (c *Client) AuthHandler() mqtt.MessageHandler {
	return func(client mqtt.Client, topicIds []string, payloadId uint16, payload interface{}, err error) {

	}
}

func (c *Client) Register(authPassword string) {
	var (
		clientRegisterTopic = mqtt.Topic(fmt.Sprintf("client/%s/register", c.id))

		authMessage = payload.NewPayload(
			payload.WithData(authPassword),
			payload.ForCommand("/auth"),
		)

		// Send a message
		err = c.mqttClient.Publish(clientRegisterTopic, authMessage)
	)

	if err != nil {
		log.WithError(err).Error("Cannot send authorization request")
	}
}

func (c *Client) SendHeartbeat(ctx context.Context, timeout time.Duration) {
	// Every minute, send a heartbeat
	heartBeatTopic := mqtt.Topic(fmt.Sprintf("client/%s/heartbeat", c.id))

heartbeatLoop:
	for {
		select {
		case <-ctx.Done():
			break heartbeatLoop
		default:
			_ = c.mqttClient.Publish(heartBeatTopic, nil)
			time.Sleep(timeout)
		}
	}
}
