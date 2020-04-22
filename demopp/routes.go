package demopp

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

// Router returns the API router
func Router() http.Handler {
	mux := httprouter.New()

	// Cashier endpoints
	mux.GET("/v1/status", statusHandler)
	mux.GET("/v1/info", infoHandler)
	mux.POST("/v1/authorize/:pk/:mk", addPreauthHandler)
	mux.DELETE("/v1/authorize/:pk/:mk", delPreauthHandler)
	mux.POST("/v1/auth", authHandler)
	mux.POST("/v1/charge/auth/:pa", preauthChargeHandler)
	mux.POST("/v1/charge/onetime/:pa", onetimeChargeHandler)
	mux.POST("/v1/pay", payHandler)
	// mux.POST("/v1/error", errorHandler)
	mux.POST("/v1/success", successHandler)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	})

	handler := c.Handler(mux)

	return handler
}
