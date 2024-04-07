package config

import "github.com/joho/godotenv"

type Config struct {
	HTTPConfig *httpConfig
	PGConfig   *pgConfig
}

func New(path string) (*Config, error) {
	err := load(path)
	if err != nil {
		return nil, err
	}
	p, err := newPGConfig()
	if err != nil {
		return nil, err
	}
	h, err := newHTTPConfig()
	if err != nil {
		return nil, err
	}
	return &Config{
		HTTPConfig: h,
		PGConfig:   p,
	}, nil
}

func load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
