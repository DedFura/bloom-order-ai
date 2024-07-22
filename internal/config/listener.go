package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type Listener struct {
	Addr string `yaml:"addr"`
}

func (l *Listener) Validate() error {
	return errors.Wrap(validation.ValidateStruct(l,
		validation.Field(&l.Addr, validation.Required),
	), "invalid listener config")
}
