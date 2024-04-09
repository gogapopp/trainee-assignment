package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPConfig  *httpConfig
	PGConfig    *pgConfig
	JWT_SECRET  string
	PASS_SECRET string
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
		HTTPConfig:  h,
		PGConfig:    p,
		JWT_SECRET:  os.Getenv("JWT_SECRET_KEY"),
		PASS_SECRET: os.Getenv("PASS_HASH_SECRET"),
	}, nil
}

func load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
