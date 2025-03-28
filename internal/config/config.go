package config

type Config struct {
	// 雪花算法配置
	Snowflake *Snowflake `json:"snowflake,omitempty" validate:"required"`
}

func newConfig(wrapper *Wrapper) (config *Config) {
	return wrapper.Id
}
