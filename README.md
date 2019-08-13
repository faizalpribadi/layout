# LAYOUT

<p align="center"><img src="images/layout.png" width="220"></p>

## MICRO-SERVICE-PROJECT-LAYOUT

Description is `ambigue` for this repository

### REQUIREMENTS

- `Go 1.9+`

- `Docker and Docker Compose`

### DEPENDENCIES

- `github.com/google/wire/cmd/wire`

- `github.com/satori/go.uuid`

- `github.com/sirupsen/logrus`

- `github.com/spf13/cobra`

- `github.com/spf13/viper`

- `github.com/urfave/negroni`

### FEATURES

- [x] `OpenAPI 3 Approach`

- [x] `Docker Deployments`

- [x] `Dependency Injection Based on Interface with Google Wire`

- [x] `HTTP Multiplexer`

- [x] `Graceful Shutdown`

- [x] `Hystrix circuit breaker`

- [x] `Domain Driven Design Approach`

- [x] `Panic recovery and Logger Middleware support by Negroni`

- [x] `Centralized Configuration`

- [x] `Support DRY for Code Standard`

- [ ] `Support Opentracing or DataDog for distributed tracing`

- [ ] `Support Business metrics for DataDog`

- [ ] `Support message queue or worker job processing with kafka or nats`

- [ ] `gRPC or WebSocket protocol`

## INSTALLATION

- `$ git clone https://github.com/faizalpribadi/layout.git`

- `$ cd layout && dep ensure`

- `$ go run main.go serve-http`
