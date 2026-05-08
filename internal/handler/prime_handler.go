package handler

import (
	"net/http"
	"prime-generator/internal/models"
	"prime-generator/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler acts as the HTTP layer.
// It receives requests, validates input, and delegates work to the service layer.
type PrimeHandler struct {
	service *service.PrimeService
}


func NewPrimeHandler(s *service.PrimeService) *PrimeHandler {
	return &PrimeHandler{
		service: s,
	}
}

func (h *PrimeHandler) GeneratePrimes(c *gin.Context) {

	startStr := c.Query("start")
	endStr := c.Query("end")
	algo := c.DefaultQuery("algo", "auto")

	start, err := strconv.Atoi(startStr)
	if err != nil {

		c.JSON(http.StatusBadRequest, models.APIResponse{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid start value",
			Data:       nil,
		})
		return
	}

	end, err := strconv.Atoi(endStr)
	if err != nil {

		c.JSON(http.StatusBadRequest, models.APIResponse{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid end value",
			Data:       nil,
		})
		return
	}

	if start > end {

		c.JSON(http.StatusBadRequest, models.APIResponse{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Message:    "Start cannot be greater than end",
			Data:       nil,
		})
		return
	}

	result, err := h.service.Generate(start, end, algo)
	if err != nil {

		c.JSON(http.StatusInternalServerError, models.APIResponse{
			Status:     "error",
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	response := models.APIResponse{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Prime numbers fetched successfully",
		Data:       result,
	}

	c.JSON(http.StatusOK, response)
}

func (h *PrimeHandler) GetStats(c *gin.Context) {
	stats := h.service.GetStats()
	c.JSON(http.StatusOK, stats)
}