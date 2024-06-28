package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetTotalRevenue(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	totalRevenue, err := h.service.CalculateTotalRevenue(ctx, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]float64{"total_revenue": totalRevenue}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetTotalRevenueByProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	totalRevenueByProduct, err := h.service.CalculateTotalRevenueByProduct(ctx, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(totalRevenueByProduct)
}

func (h *Handler) GetTotalRevenueByCategory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	totalRevenueByCategory, err := h.service.CalculateTotalRevenueByCategory(ctx, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(totalRevenueByCategory)
}

func (h *Handler) GetTotalRevenueByRegion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	totalRevenueByRegion, err := h.service.CalculateTotalRevenueByRegion(ctx, startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(totalRevenueByRegion)
}
