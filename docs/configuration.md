# Client configuration

Below is an example for client configuration. It requires a MQTT connection and basic information about the client, such
as nodeId, authentication password and a list of plugins that the client supports.

```yaml
mqtt:
  host: localhost
  port: 1883
  username: "testuser"
  password: "testpassword"
  tls:
    isEnabled: false
    certPath: ""
    keyPath: ""

logging:
  type:
    - "file"
    - "remote"
  address: "localhost:514"
  format: "syslog"

client:
  serviceNodeId: "exampleId"
  authPassword: "examplepass"
  pluginDir: ../detecc-core/compiled/client
  plugins:
    - "hw-monitor"
```