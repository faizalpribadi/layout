package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/faizalpribadi/layout/internal/config"
)

type (
	Breaker interface{}
	breaker struct {
		cfg config.Configuration
	}
)

func NewCircuitBreaker(cfg config.Configuration) *breaker {
	return &breaker{cfg}
}

//Register the key names of hystrix command
func (b *breaker) Register(names ...string) {
	commands := make(map[string]hystrix.CommandConfig)
	for _, name := range names {
		commands[name] = hystrix.CommandConfig{
			Timeout: b.cfg.CircuitBreaker.Timeout,
		}
	}

	hystrix.Configure(commands)
}
