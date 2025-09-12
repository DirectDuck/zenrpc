module github.com/vmkteam/zenrpc/v2

go 1.24.0

require (
	github.com/gorilla/websocket v1.5.3
	github.com/prometheus/client_golang v1.23.2
	github.com/smartystreets/goconvey v1.8.1
	golang.org/x/text v0.29.0
	golang.org/x/tools v0.37.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.66.1 // indirect
	github.com/prometheus/procfs v0.17.0 // indirect
	github.com/smarty/assertions v1.16.0 // indirect
	go.yaml.in/yaml/v2 v2.4.3 // indirect
	golang.org/x/mod v0.28.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	google.golang.org/protobuf v1.36.9 // indirect
)

retract (
	v2.2.10 // invalid version cached
	v2.2.5 // invalid version cached
)
