package plugin

import "context"

const (
	ClientServer = Type("clientServer")
	ClientOnly   = Type("clientOnly")
)

type (
	Type string

	// Handler is the interface for the Plugin.
	Handler interface {
		// Execute method is called when a client receives a command from the server.
		// The arguments of the method is Payload.Data, produced by the corresponding server cmd-manager.
		// The response should be data, ready to be sent back to the server for processing and an error, if one occurred.
		Execute(ctx context.Context, args interface{}) (interface{}, error)

		// GetMetadata returns the metadata of the client cmd-manager.
		GetMetadata() Metadata
	}

	// Metadata is the metadata object for the cmd-manager and determines the behaviour of the cmd-manager interaction with the client.
	Metadata struct {
		Type Type
	}
)
