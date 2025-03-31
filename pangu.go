package snowflake

import (
	"github.com/pangum/pangu"
	"github.com/pangum/snowflake/internal/plugin"
)

func init() {
	pangu.New().Get().Dependency().Put(
		plugin.New,
	).Build().Build().Apply()
}
