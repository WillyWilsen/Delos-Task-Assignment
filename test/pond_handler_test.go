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

func TestCreatePond(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.POST("/pond", pondHandler.CreatePond)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a test request
	requestBody := strings.NewReader(`{"name": "Pond 1", "farm_id": 1}`)
	request, _ := http.NewRequest("POST", "/pond", requestBody)
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
		"message": "New pond created successfully",
		"data": map[string]interface {}{
			"id":   1,
			"name": "Pond 1",
			"farm_id": 1,
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
		if farmIdFloat, ok := data["farm_id"].(float64); ok {
			farmId := int(farmIdFloat)
			data["farm_id"] = farmId
		}
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	// Check if the pond was created in the repository
	ponds, _ := pondRepo.Get()
	assert.Len(t, ponds, 1)
	assert.Equal(t, "Pond 1", ponds[0].Name)
}

func TestGetPond(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.GET("/pond", pondHandler.GetPond)

	// Create a test request
	request, _ := http.NewRequest("GET", "/pond", nil)
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
		"message": "Pond fetched successfully",
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

func TestGetPondById(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.GET("/pond/:id", pondHandler.GetPondById)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a pond for testing
	pondRepo.Create(&model.Pond{
		ID:   1,
		Name: "Pond 1",
		FarmID: 1,
	})

	// Create a test request
	request, _ := http.NewRequest("GET", "/pond/1", nil)
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
		"message": "Pond fetched successfully",
		"data": map[string]interface {}{
			"id":   1,
			"name": "Pond 1",
			"farm_id": 1,
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
		if farmIdFloat, ok := data["farm_id"].(float64); ok {
			farmId := int(farmIdFloat)
			data["farm_id"] = farmId
		}
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestUpdatePond(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.PUT("/pond/:id", pondHandler.UpdatePond)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a pond for testing
	pondRepo.Create(&model.Pond{
		ID:   1,
		Name: "Pond 1",
		FarmID: 1,
	})

	// Create a test request
	requestBody := strings.NewReader(`{"name": "Updated Pond", "farm_id": 1}`)
	request, _ := http.NewRequest("PUT", "/pond/1", requestBody)
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
		"message": "Pond updated successfully",
		"data": map[string]interface {}{
			"id":   1,
			"name": "Updated Pond",
			"farm_id": 1,
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
		if farmIdFloat, ok := data["farm_id"].(float64); ok {
			farmId := int(farmIdFloat)
			data["farm_id"] = farmId
		}
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	// Check if the pond was updated in the repository
	ponds, _ := pondRepo.Get()
	assert.Len(t, ponds, 1)
	assert.Equal(t, "Updated Pond", ponds[0].Name)
}

func TestDeletePond(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.DELETE("/pond/:id", pondHandler.DeletePond)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a pond for testing
	pondRepo.Create(&model.Pond{
		ID:   1,
		Name: "Pond 1",
		FarmID: 1,
	})

	// Create a test request
	request, _ := http.NewRequest("DELETE", "/pond/1", nil)
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
		"message": "Pond deleted successfully",
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

	// Check if the pond was deleted from the repository
	ponds, _ := pondRepo.Get()
	assert.Len(t, ponds, 0)
}

func TestCreatePond_NameExists(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.POST("/pond", pondHandler.CreatePond)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a pond for testing
	pondRepo.Create(&model.Pond{
		ID:   1,
		Name: "Pond 1",
		FarmID: 1,
	})

	// Create a test request
	requestBody := strings.NewReader(`{"name": "Pond 1", "farm_id": 1}`)
	request, _ := http.NewRequest("POST", "/pond", requestBody)
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
		"message": "Pond name already exists",
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

func TestGetPondById_InvalidParam(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.GET("/pond/:id", pondHandler.GetPondById)

	// Create a test request with an invalid ID param
	request, _ := http.NewRequest("GET", "/pond/invalid", nil)
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

func TestUpdatePond_NotFound(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.PUT("/pond/:id", pondHandler.UpdatePond)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a test request with a non-existing ID
	requestBody := strings.NewReader(`{"name": "Updated Pond", "farm_id": 1}`)
	request, _ := http.NewRequest("PUT", "/pond/1", requestBody)
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
		"message": "Data Not Found. New pond created successfully",
		"data": map[string]interface {}{
			"id":   1,
			"name": "Updated Pond",
			"farm_id": 1,
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
		if farmIdFloat, ok := data["farm_id"].(float64); ok {
			farmId := int(farmIdFloat)
			data["farm_id"] = farmId
		}
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	// Check if the pond was created in the repository
	ponds, _ := pondRepo.Get()
	assert.Len(t, ponds, 1)
	assert.Equal(t, "Updated Pond", ponds[0].Name)
}

func TestDeletePond_NotFound(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.DELETE("/pond/:id", pondHandler.DeletePond)

	// Create a farm for testing
	farmRepo.Create(&model.Farm{
		ID:   1,
		Name: "Farm 1",
	})

	// Create a pond for testing
	pondRepo.Create(&model.Pond{
		ID:   1,
		Name: "Pond 1",
		FarmID: 1,
	})

	// Create a test request with a non-existing ID
	request, _ := http.NewRequest("DELETE", "/pond/2", nil)
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

	// Check if the pond was not deleted from the repository
	ponds, _ := pondRepo.Get()
	assert.Len(t, ponds, 1)
}

func TestCreatePond_InvalidPayload(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.POST("/pond", pondHandler.CreatePond)

	// Create a test request with an invalid payload (missing name field)
	requestBody := strings.NewReader(`{}`)
	request, _ := http.NewRequest("POST", "/pond", requestBody)
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

func TestGetPondById_NotFound(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.GET("/pond/:id", pondHandler.GetPondById)

	// Create a test request with a non-existing ID
	request, _ := http.NewRequest("GET", "/pond/1", nil)
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

func TestUpdatePond_InvalidPayload(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.PUT("/pond/:id", pondHandler.UpdatePond)

	// Create a test request with an invalid payload (missing name field)
	requestBody := strings.NewReader(`{}`)
	request, _ := http.NewRequest("PUT", "/pond/1", requestBody)
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

func TestDeletePond_InvalidParam(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.DELETE("/pond/:id", pondHandler.DeletePond)

	// Create a test request with an invalid ID param
	request, _ := http.NewRequest("DELETE", "/pond/invalid", nil)
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

func TestCreatePond_FarmDataNotFound(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.POST("/pond", pondHandler.CreatePond)

	// Create a test request with a non-existing Farm ID
	requestBody := strings.NewReader(`{"name": "Pond 1", "farm_id": 1}`)
	request, _ := http.NewRequest("POST", "/pond", requestBody)
	request.Header.Set("User-Agent", "Test Agent")
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusNotFound,
		"status":  "error",
		"message": "Farm Data Not Found",
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

func TestUpdatePond_FarmDataNotFound(t *testing.T) {
	// Create mock repositories
	farmRepo := repository.NewMockFarmRepository()
	pondRepo := repository.NewMockPondRepository()
	logRepo := repository.NewMockLogRepository()

	// Create handler with mock repositories
	pondHandler := handler.NewPondHandler(pondRepo, farmRepo, logRepo)

	// Create a Gin router and set up the handler route
	router := gin.Default()
	router.PUT("/pond/:id", pondHandler.UpdatePond)

	// Create a test request with a non-existing Farm ID
	requestBody := strings.NewReader(`{"name": "Updated Pond", "farm_id": 1}`)
	request, _ := http.NewRequest("PUT", "/pond/1", requestBody)
	request.Header.Set("User-Agent", "Test Agent")
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(responseRecorder, request)

	// Check the response status code
	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)

	// Check the response body
	expectedResponse := gin.H{
		"code":    http.StatusNotFound,
		"status":  "error",
		"message": "Farm Data Not Found",
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