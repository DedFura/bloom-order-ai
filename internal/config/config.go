package config

import (
	"gopkg.in/yaml.v3"
	"os"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
)

type Config struct {
	Addr string
	Log  string
}

func NewConfigFromFile(path string) (*Config, error) {
	r, err := newRawFromFile(path)
	if err != nil {
		return nil, err
	}
	if err = r.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid config")
	}

	return r.parse()
}

type raw struct {
	Listener Listener `yaml:"listener"`
	Log      Log      `yaml:"logger"`
}

func newRawFromFile(path string) (*raw, error) {
	r := new(raw)
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config file")
	}

	err = yaml.NewDecoder(file).Decode(r)
	if err != nil {
		return nil, errors.Wrap(err, "to parse yaml")
	}

	return r, nil
}

func (r raw) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Listener),
	)
}

func (r raw) parse() (*Config, error) {
	var cfg Config

	cfg.Addr = r.Listener.Addr
	cfg.Log = r.Log.Parse()

	return &cfg, nil
}
