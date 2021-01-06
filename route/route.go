package route

import (
	"github.com/FerdinaKusumah/fastcrud/handler"
	"github.com/FerdinaKusumah/fastcrud/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	var router = mux.NewRouter()
	// enabling cors
	router.Use(mux.CORSMethodMiddleware(router))
	// health check
	router.HandleFunc("/api/hello", handler.HealthCheck).Methods("GET", "OPTIONS")
	// load data example
	router.HandleFunc("/api/seed-example", handler.LoadExampleApi).Methods("GET", "OPTIONS")

	// middleware to set header response
	router.Use(middleware.SetRespJson)

	// crud method
	router.HandleFunc("/api/{key}", handler.ListResource).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/{key}", handler.CreateResource).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/{key}/bulk", handler.CreateBulkResource).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/{key}/{id}", handler.DetailResource).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/{key}/{id}", handler.UpdateResource).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/{key}/{id}", handler.DeleteResource).Methods("DELETE", "OPTIONS")
	return router
}
