package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type Log struct {
	Level string `yaml:"level"`
}

func (l *Log) Validate() error {
	return errors.Wrap(validation.ValidateStruct(l,
		validation.Field(&l.Level, validation.Required),
	), "invalid logger config")
}

func (l *Log) Parse() string {
	return l.Level
}
