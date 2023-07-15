package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
	"github.com/WillyWilsen/Delos-Task-Assignment.git/model"
    "github.com/WillyWilsen/Delos-Task-Assignment.git/repository"
)

type StatisticsHandler struct {
    logRepository repository.LogRepository
}

func NewStatisticsHandler(
    logRepository repository.LogRepository,
) *StatisticsHandler {
    return &StatisticsHandler{
        logRepository: logRepository,
    }
}

func (h *StatisticsHandler) GetStatistics(c *gin.Context) {
	// Create log
    log := model.Log{
		Endpoint:  "GET /log",
		UserAgent: c.GetHeader("User-Agent"),
	}
    if err := h.logRepository.Create(&log); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusInternalServerError,
            "status": "error",
            "message": "Failed to create log",
        })
        return
    }

    // Get Endpoints List
    endpoints, _ := h.logRepository.GetDistinctEndpoints()

    // Make statistics
    statistics := make(model.Statistics)
	for _, endpoint := range endpoints {
		endpointStatistics, _ := h.logRepository.GetEndpointStatistics(endpoint)
		statistics[endpoint] = endpointStatistics
	}

	// Success
    c.JSON(http.StatusOK, gin.H{
        "code": http.StatusOK,
        "status": "success",
        "message": "Statistic fetched successfully",
        "data": statistics,
    })
}