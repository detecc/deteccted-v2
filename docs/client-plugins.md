# Client plugins

Client plugins must implement the `Handler` interface defined in the `plugin` package. The plugin should be registered
by calling the `GetPluginManager().Register(commandName, plugin)` or `pluginManager.Register(commandName, plugin)`.

## Plugin interface

```go
package example

import "context"

// Handler is the interface for the Plugin.
type Handler interface {
	// Execute method is called when a client receives a command from the server.
	// The arguments of the method is Payload.Data, produced by the corresponding server cmd-manager.
	// The response should be data, ready to be sent back to the server for processing and an error, if one occurred.
	Execute(ctx context.Context, args interface{}) (interface{}, error)

	// GetMetadata returns the metadata of the client cmd-manager.
	GetMetadata() Metadata
}
```

## Plugin example

```go
package example

import (
	"log"
	"github.com/detecc/deteccted-v2/internal/models/plugin"
	pluginManager "github.com/detecc/deteccted-v2/internal/plugin-manager"
)

func init() {
	examplePlugin := &ExamplePlugin{}
	pluginManager.Register(examplePlugin.GetCmdName(), examplePlugin)
}

type ExamplePlugin struct {
	plugin.Handler
}

func (e ExamplePlugin) GetCmdName() string {
	return "/exampleCmd"
}

func (e ExamplePlugin) Execute(args interface{}) (interface{}, error) {
	log.Println(args)
	return "ping", nil
}

func (e ExamplePlugin) GetMetadata() plugin.Metadata {
	return plugin.Metadata{Type: plugin.ClientServer}
}
```

## Compiling the plugin

The plugin is compiled using the following command:

```bash
go build -buildmode=plugin . 
```

It produces a file with `.so` format. The file's name (without the format) is then specified in the configuration file
under `plugins`. 