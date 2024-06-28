package routers

import (
	"lumel/internal/handlers"

	"github.com/gorilla/mux"
)

func InitRouter(handler *handlers.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/revenue", handler.GetTotalRevenue).Methods("GET")
	r.HandleFunc("/revenueByProduct", handler.GetTotalRevenueByProduct).Methods("GET")
	r.HandleFunc("/revenueByCategory", handler.GetTotalRevenueByCategory).Methods("GET")
	r.HandleFunc("/revenueByRegion", handler.GetTotalRevenueByRegion).Methods("GET")

	return r
}
