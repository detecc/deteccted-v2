module github.com/detecc/deteccted-v2

go 1.17

require (
	github.com/agrison/go-commons-lang v0.0.0
	github.com/eclipse/paho.mqtt.golang v1.3.5
	github.com/kkyr/fig v0.3.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
)

//replace github.com/detecc/detecctor-v2 => ../detecctor-v2

require (
	github.com/detecc/detecctor-v2 v0.0.0
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/rifflock/lfshook v0.0.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/kr/text v0.1.0 // indirect
	github.com/lestrrat-go/strftime v1.0.5 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cobra v1.3.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20211205182925-97ca703d548d // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
