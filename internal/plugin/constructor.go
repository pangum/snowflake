package plugin

import (
	"github.com/goexl/id"
	"github.com/goexl/snowflake"
	"github.com/pangum/snowflake/internal/config"
)

func New(config *config.Snowflake) id.Generator {
	return snowflake.New().
		Node(config.Node).
		Worker(config.Worker).
		Sequence(config.Sequence).
		Start(config.Start).
		Build()
}
