package repositories

import (
	"context"
	"fmt"
	"log"
	"lumel/internal/repositories/models"
	"lumel/internal/utils"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (r *repository) GetTotalRevenue(ctx context.Context, startDate, endDate string) (float64, error) {
	start, err := utils.ParseDate(startDate)
	if err != nil {
		return 0, err
	}
	end, err := utils.ParseDate(endDate)
	if err != nil {
		return 0, err
	}

	orders, err := models.Orders(qm.Where("date_of_sale BETWEEN ? AND ?", start, end)).All(ctx, r.db)
	if err != nil {
		return 0, err
	}

	var totalRevenue float64
	for _, order := range orders {
		unitPrice, _ := order.UnitPrice.Float64()
		totalRevenue += float64(order.QuantitySold.Int) * float64(unitPrice)
	}

	return totalRevenue, nil
}

func (r *repository) GetTotalRevenueByProduct(ctx context.Context, startDate, endDate string) (map[string]float64, error) {
	start, err := utils.ParseDate(startDate)
	if err != nil {
		return nil, err
	}
	end, err := utils.ParseDate(endDate)
	if err != nil {
		return nil, err
	}

	orders, err := models.Orders(qm.Where("date_of_sale BETWEEN ? AND ?", start, end)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	// Calculate total revenue by product
	revenueByProduct := make(map[string]float64)
	for _, order := range orders {
		product, _ := models.Products(qm.Where("product_id = ?", order.ProductID)).One(ctx, r.db)
		unitPrice, _ := order.UnitPrice.Float64()
		revenueByProduct[product.ProductName.String] += float64(order.QuantitySold.Int) * float64(unitPrice)
	}

	return revenueByProduct, nil
}

func (r *repository) GetTotalRevenueByCategory(ctx context.Context, startDate, endDate string) (map[string]float64, error) {
	start, err := utils.ParseDate(startDate)
	if err != nil {
		return nil, err
	}
	end, err := utils.ParseDate(endDate)
	if err != nil {
		return nil, err
	}

	orders, err := models.Orders(qm.Where("date_of_sale BETWEEN ? AND ?", start, end)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	revenueByCategory := make(map[string]float64)
	for _, order := range orders {
		product, _ := models.Products(qm.Where("product_id = ?", order.ProductID)).One(ctx, r.db)
		unitPrice, _ := order.UnitPrice.Float64()
		revenueByCategory[product.Category.String] += float64(order.QuantitySold.Int) * float64(unitPrice)
	}

	return revenueByCategory, nil
}

func (r *repository) GetTotalRevenueByRegion(ctx context.Context, startDate, endDate string) (map[string]float64, error) {
	start, err := utils.ParseDate(startDate)
	if err != nil {
		return nil, err
	}
	end, err := utils.ParseDate(endDate)
	if err != nil {
		return nil, err
	}

	orders, err := models.Orders(qm.Where("date_of_sale BETWEEN ? AND ?", start, end)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	fmt.Println(orders)

	revenueByRegion := make(map[string]float64)
	for _, order := range orders {
		unitPrice, _ := order.UnitPrice.Float64()
		revenueByRegion[order.Region.String] += float64(order.QuantitySold.Int) * float64(unitPrice)

	}

	return revenueByRegion, nil
}

func (r *repository) Refresh(ctx context.Context) {
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				_, err := models.Orders().DeleteAll(ctx, r.db)
				_, err = models.Products().DeleteAll(ctx, r.db)
				_, err = models.Customers().DeleteAll(ctx, r.db)
				if err != nil {
					log.Printf("Error refreshing data: %v", err)
				} else {
					log.Println("Data refresh successful")
				}
			}
		}
	}()
}
