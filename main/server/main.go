package main

import (
	"fmt"
	"prime-generator/internal/handler"
	"prime-generator/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Prime Generator Server Starting...")

	// Initialize service
	svc := service.NewPrimeService()

	// Initialize handler
	h := handler.NewPrimeHandler(svc)

	// Create Gin router
	r := gin.Default()

	// Routes
	r.GET("/primes", h.GeneratePrimes)
	r.GET("/stats", h.GetStats)

	// Start server
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}