package config

import (
	"github.com/pangum/pangu"
)

func init() {
	pangu.New().Get().Dependency().Puts(
		newWrapper, // 外层配置

		newConfig,    // 统一标识配置
		newSnowflake, // 配置
	).Build().Apply()
}
