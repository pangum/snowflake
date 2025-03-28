package config

import (
	"time"
)

type Snowflake struct {
	// 开始时间
	Start time.Time `json:"start,omitempty"`
	// 集群标识
	Node uint64 `default:"1" json:"node,omitempty" validate:"required,min=1,max=32"`
	// 工作节点标识
	Worker uint64 `default:"1" json:"worker,omitempty" validate:"required,min=1,max=32"`
	// 顺序号
	Sequence uint64 `default:"1" json:"sequence,omitempty" validate:"required,min=0,max=4096"`
}

func newSnowflake(config *Config) *Snowflake {
	return config.Snowflake
}
