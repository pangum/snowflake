package config

import (
	"github.com/pangum/pangu"
)

type Wrapper struct {
	Id *Config `json:"id,omitempty" validate:"required"`
}

func newWrapper(config *pangu.Config) (wrapper *Wrapper, err error) {
	wrapper = new(Wrapper)
	err = config.Build().Get(wrapper)

	return
}
