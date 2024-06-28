package repositories

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

type Repository interface {
	GetTotalRevenue(ctx context.Context, startDate, endDate string) (float64, error)
	GetTotalRevenueByProduct(ctx context.Context, startDate, endDate string) (map[string]float64, error)
	GetTotalRevenueByCategory(ctx context.Context, startDate, endDate string) (map[string]float64, error)
	GetTotalRevenueByRegion(ctx context.Context, startDate, endDate string) (map[string]float64, error)
	Refresh(ctx context.Context)
}

type repository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewRepository(db *sql.DB, logger *logrus.Logger) Repository {
	return &repository{db: db, logger: logger}
}
