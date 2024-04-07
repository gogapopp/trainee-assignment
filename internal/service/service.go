package service

import "go.uber.org/zap"

type (
	Repostiry interface {
	}

	Service struct {
		logger *zap.SugaredLogger
		repo   Repostiry
	}
)

func New(logger *zap.SugaredLogger, repository Repostiry) *Service {
	return &Service{
		logger: logger,
		repo:   repository,
	}
}
