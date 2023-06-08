module data-platform-api-request-reads-cache-manager-rmq-kube

go 1.19

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/latonaio/golang-logging-library-for-data-platform v1.0.4
	github.com/latonaio/rabbitmq-golang-client-for-data-platform v1.0.4
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
	golang.org/x/net v0.6.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
)

replace (
	github.com/latonaio/golang-logging-library-for-data-platform v1.0.2 => /home/ampamman/go/src/latona/golang-logging-library-for-data-platform
	github.com/latonaio/rabbitmq-golang-client-for-data-platform v1.0.3 => /home/ampamman/go/src/latona/rabbitmq-golang-client-for-data-platform
)
