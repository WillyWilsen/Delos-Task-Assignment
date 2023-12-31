package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/WillyWilsen/Delos-Task-Assignment.git/handler"
	"github.com/WillyWilsen/Delos-Task-Assignment.git/model"
	"github.com/WillyWilsen/Delos-Task-Assignment.git/test/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateFarm(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.POST("/farm", farmHandler.CreateFarm)

	// Create a test request
	requestBody := strings.NewReader(`{"name": "Farm 1"}`)
	request, _ := http.NewRequest("POST", "/farm", requestBody)
	request.Header.Set("User-Agent", "Test Agent")
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "New farm created successfully",
		"data": map[string]interface {}{
			"id":   1,
			"name": "Farm 1",
		},
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)
	
	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}

	// Convert data.id to int if it exists
	if data, ok := actualResponse["data"].(map[string]interface{}); ok {
		if idFloat, ok := data["id"].(float64); ok {
			id := int(idFloat)
			data["id"] = id
		}
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	// Check if the farm was created in the repository
	farms, _ := farmRepo.Get()
	assert.Len(t, farms, 1)
	assert.Equal(t, "Farm 1", farms[0].Name)
}

func TestGetFarm(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.GET("/farm", farmHandler.GetFarm)

	// Create a test request
	request, _ := http.NewRequest("GET", "/farm", nil)
	request.Header.Set("User-Agent", "Test Agent")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "Farm fetched successfully",
		"data":    []interface{}{},
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}
	
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetFarmById(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.GET("/farm/:id", farmHandler.GetFarmById)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a test request
	request, _ := http.NewRequest("GET", "/farm/1", nil)
	request.Header.Set("User-Agent", "Test Agent")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "Farm fetched successfully",
		"data": map[string]interface {}{
			"id":   1,
			"name": "Farm 1",
		},
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}

	// Convert data.id to int if it exists
	if data, ok := actualResponse["data"].(map[string]interface{}); ok {
		if idFloat, ok := data["id"].(float64); ok {
			id := int(idFloat)
			data["id"] = id
		}
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestUpdateFarm(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.PUT("/farm/:id", farmHandler.UpdateFarm)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a test request
	requestBody := strings.NewReader(`{"name": "Updated Farm"}`)
	request, _ := http.NewRequest("PUT", "/farm/1", requestBody)
	request.Header.Set("User-Agent", "Test Agent")
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "Farm updated successfully",
		"data": map[string]interface {}{
			"id":   1,
			"name": "Updated Farm",
		},
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}

	// Convert data.id to int if it exists
	if data, ok := actualResponse["data"].(map[string]interface{}); ok {
		if idFloat, ok := data["id"].(float64); ok {
			id := int(idFloat)
			data["id"] = id
		}
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	// Check if the farm was updated in the repository
	farms, _ := farmRepo.Get()
	assert.Len(t, farms, 1)
	assert.Equal(t, "Updated Farm", farms[0].Name)
}

func TestDeleteFarm(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.DELETE("/farm/:id", farmHandler.DeleteFarm)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a test request
	request, _ := http.NewRequest("DELETE", "/farm/1", nil)
	request.Header.Set("User-Agent", "Test Agent")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "Farm deleted successfully",
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	// Check if the farm was deleted from the repository
	farms, _ := farmRepo.Get()
	assert.Len(t, farms, 0)
}

func TestCreateFarm_NameExists(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.POST("/farm", farmHandler.CreateFarm)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a test request
	requestBody := strings.NewReader(`{"name": "Farm 1"}`)
	request, _ := http.NewRequest("POST", "/farm", requestBody)
	request.Header.Set("User-Agent", "Test Agent")
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusConflict, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusConflict,
		"status":  "error",
		"message": "Farm name already exists",
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetFarmById_InvalidParam(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.GET("/farm/:id", farmHandler.GetFarmById)

	// Create a test request with an invalid ID param
	request, _ := http.NewRequest("GET", "/farm/invalid", nil)
	request.Header.Set("User-Agent", "Test Agent")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusBadRequest,
		"status":  "error",
		"message": "Invalid request param",
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}
	
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestUpdateFarm_NotFound(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.PUT("/farm/:id", farmHandler.UpdateFarm)

	// Create a test request with a non-existing ID
	requestBody := strings.NewReader(`{"name": "Updated Farm"}`)
	request, _ := http.NewRequest("PUT", "/farm/1", requestBody)
	request.Header.Set("User-Agent", "Test Agent")
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusOK,
		"status":  "success",
		"message": "Data Not Found. New farm created successfully",
		"data": map[string]interface {}{
			"id":   1,
			"name": "Updated Farm",
		},
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}

	// Convert data.id to int if it exists
	if data, ok := actualResponse["data"].(map[string]interface{}); ok {
		if idFloat, ok := data["id"].(float64); ok {
			id := int(idFloat)
			data["id"] = id
		}
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	// Check if the farm was created in the repository
	farms, _ := farmRepo.Get()
	assert.Len(t, farms, 1)
	assert.Equal(t, "Updated Farm", farms[0].Name)
}

func TestDeleteFarm_NotFound(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.DELETE("/farm/:id", farmHandler.DeleteFarm)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a test request with a non-existing ID
	request, _ := http.NewRequest("DELETE", "/farm/2", nil)
	request.Header.Set("User-Agent", "Test Agent")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusNotFound,
		"status":  "error",
		"message": "Data Not Found",
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	// Check if the farm was not deleted from the repository
	farms, _ := farmRepo.Get()
	assert.Len(t, farms, 1)
}

func TestCreateFarm_InvalidPayload(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.POST("/farm", farmHandler.CreateFarm)

	// Create a test request with an invalid payload (missing name field)
	requestBody := strings.NewReader(`{}`)
	request, _ := http.NewRequest("POST", "/farm", requestBody)
	request.Header.Set("User-Agent", "Test Agent")
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusBadRequest,
		"status":  "error",
		"message": "Invalid request payload",
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetFarmById_NotFound(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.GET("/farm/:id", farmHandler.GetFarmById)

	// Create a test request with a non-existing ID
	request, _ := http.NewRequest("GET", "/farm/1", nil)
	request.Header.Set("User-Agent", "Test Agent")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusNotFound,
		"status":  "error",
		"message": "Data Not Found",
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestUpdateFarm_InvalidPayload(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.PUT("/farm/:id", farmHandler.UpdateFarm)

	// Create a test request with an invalid payload (missing name field)
	requestBody := strings.NewReader(`{}`)
	request, _ := http.NewRequest("PUT", "/farm/1", requestBody)
	request.Header.Set("User-Agent", "Test Agent")
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusBadRequest,
		"status":  "error",
		"message": "Invalid request payload",
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestDeleteFarm_InvalidParam(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	farmHandler := handler.NewFarmHandler(farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.DELETE("/farm/:id", farmHandler.DeleteFarm)

	// Create a test request with an invalid ID param
	request, _ := http.NewRequest("DELETE", "/farm/invalid", nil)
	request.Header.Set("User-Agent", "Test Agent")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusBadRequest,
		"status":  "error",
		"message": "Invalid request param",
	}
	actualResponse := gin.H{}
	err := json.Unmarshal(responseRecorder.Body.Bytes(), &actualResponse)

	// Convert code to int if it exists
	if codeFloat, ok := actualResponse["code"].(float64); ok {
		code := int(codeFloat)
		actualResponse["code"] = code
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
