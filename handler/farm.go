package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/WillyWilsen/Delos-Task-Assignment.git/model"
    "github.com/WillyWilsen/Delos-Task-Assignment.git/repository"
)

type FarmHandler struct {
    farmRepository repository.FarmRepository
    logRepository repository.LogRepository
}

func NewFarmHandler(
    farmRepository repository.FarmRepository, 
    logRepository repository.LogRepository,
) *FarmHandler {
    return &FarmHandler{
        farmRepository: farmRepository,
        logRepository: logRepository,
    }
}

func (h *FarmHandler) CreateFarm(c *gin.Context) {
    // Create log
    log := model.Log{
		Endpoint:  "POST /farm",
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

    var farm model.Farm

    // Bind payload
    if err := c.Bind(&farm); err != nil || farm.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "status": "error",
            "message": "Invalid request payload",
        })
        return
    }

    // Exist farm name
    existFarm, _ := h.farmRepository.GetByName(farm.Name)
    if existFarm != nil {
        c.JSON(http.StatusConflict, gin.H{
            "code": http.StatusConflict,
            "status": "error",
            "message": "Farm name already exists",
        })
        return
    }

    // Create farm
    if err := h.farmRepository.Create(&farm); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusInternalServerError,
            "status": "error",
            "message": "Failed to create farm",
        })
        return
    }

	// Success
    c.JSON(http.StatusOK, gin.H{
        "code": http.StatusOK,
        "status": "success",
        "message": "New farm created successfully",
        "data": farm,
    })
}

func (h *FarmHandler) GetFarm(c *gin.Context) {
    // Create log
    log := model.Log{
		Endpoint:  "GET /farm",
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

    farm, _ := h.farmRepository.Get()

    // Empty farm
    if farm == nil {
        c.JSON(http.StatusNotFound, gin.H{
            "code": http.StatusNotFound,
            "status": "error",
            "message": "Data Not Found",
        })
        return
    }

	// Success
    c.JSON(http.StatusOK, gin.H{
        "code": http.StatusOK,
        "status": "success",
        "message": "Farm fetched successfully",
        "data": farm,
    })
}

func (h *FarmHandler) GetFarmById(c *gin.Context) {
    // Create log
    log := model.Log{
		Endpoint:  "GET /farm/:id",
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

    // Get param id
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "status": "error",
            "message": "Invalid request param",
        })
        return
    }

    farm, _ := h.farmRepository.GetById(id)

    // Empty farm
    if farm == nil {
        c.JSON(http.StatusNotFound, gin.H{
            "code": http.StatusNotFound,
            "status": "error",
            "message": "Data Not Found",
        })
        return
    }

	// Success
    c.JSON(http.StatusOK, gin.H{
        "code": http.StatusOK,
        "status": "success",
        "message": "Farm fetched successfully",
        "data": farm,
    })
}

func (h *FarmHandler) UpdateFarm(c *gin.Context) {
    // Create log
    log := model.Log{
		Endpoint:  "PUT /farm/:id",
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

    // Get param id
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "status": "error",
            "message": "Invalid request param",
        })
        return
    }

    // Bind payload
    var farmPayload model.Farm
    if err := c.Bind(&farmPayload); err != nil || farmPayload.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "status": "error",
            "message": "Invalid request payload",
        })
        return
    }

    // Exist farm name
    existFarm, _ := h.farmRepository.GetByName(farmPayload.Name)
    if existFarm != nil && existFarm.ID != id {
        c.JSON(http.StatusConflict, gin.H{
            "code": http.StatusConflict,
            "status": "error",
            "message": "Farm name already exists",
        })
        return
    }

    farm, _ := h.farmRepository.GetById(id)

    // Empty farm
    if farm == nil {
        // Create farm
        if err := h.farmRepository.Create(&farmPayload); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "code": http.StatusInternalServerError,
                "status": "error",
                "message": "Failed to create farm",
            })
            return
        }

        // Success
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusOK,
            "status": "success",
            "message": "Data Not Found. New farm created successfully",
            "data": farmPayload,
        })
    } else {
        // Update farm
        if err := h.farmRepository.Update(farm.ID, &farmPayload); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "code": http.StatusInternalServerError,
                "status": "error",
                "message": "Failed to create farm",
            })
            return
        }
        farmPayload.ID = farm.ID

        // Success
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusOK,
            "status": "success",
            "message": "Farm updated successfully",
            "data": farmPayload,
        })
    }
}

func (h *FarmHandler) DeleteFarm(c *gin.Context) {
    // Create log
    log := model.Log{
		Endpoint:  "DELETE /farm/:id",
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

    // Get param id
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "status": "error",
            "message": "Invalid request param",
        })
        return
    }

    farm, _ := h.farmRepository.GetById(id)

    // Empty farm
    if farm == nil {
        c.JSON(http.StatusNotFound, gin.H{
            "code": http.StatusNotFound,
            "status": "error",
            "message": "Data Not Found",
        })
    } else {
        // Delete farm
        if err := h.farmRepository.Delete(farm); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "code": http.StatusInternalServerError,
                "status": "error",
                "message": "Failed to delete farm",
            })
            return
        }

        // Success
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusOK,
            "status": "success",
            "message": "Farm deleted successfully",
        })
    }
}