package services

import (
	"context"
	"lumel/internal/repositories"

	"github.com/sirupsen/logrus"
)

type Service interface {
	CalculateTotalRevenue(ctx context.Context, startDate string, endDate string) (float64, error)
	CalculateTotalRevenueByProduct(ctx context.Context, startDate, endDate string) (map[string]float64, error)
	CalculateTotalRevenueByCategory(ctx context.Context, startDate, endDate string) (map[string]float64, error)
	CalculateTotalRevenueByRegion(ctx context.Context, startDate, endDate string) (map[string]float64, error)
}

type service struct {
	logger *logrus.Logger
	repo   repositories.Repository
}

func NewService(logger *logrus.Logger, repo repositories.Repository) Service {
	return &service{logger: logger, repo: repo}
}
