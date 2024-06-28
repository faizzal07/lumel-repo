package services

import "context"

func (s *service) CalculateTotalRevenue(ctx context.Context, startDate, endDate string) (float64, error) {
	return s.repo.GetTotalRevenue(ctx, startDate, endDate)
}

func (s *service) CalculateTotalRevenueByProduct(ctx context.Context, startDate, endDate string) (map[string]float64, error) {
	return s.repo.GetTotalRevenueByProduct(ctx, startDate, endDate)
}

func (s *service) CalculateTotalRevenueByCategory(ctx context.Context, startDate, endDate string) (map[string]float64, error) {
	return s.repo.GetTotalRevenueByCategory(ctx, startDate, endDate)
}

func (s *service) CalculateTotalRevenueByRegion(ctx context.Context, startDate, endDate string) (map[string]float64, error) {
	return s.repo.GetTotalRevenueByRegion(ctx, startDate, endDate)
}
