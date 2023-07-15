package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/WillyWilsen/Delos-Task-Assignment.git/model"
    "github.com/WillyWilsen/Delos-Task-Assignment.git/repository"
)

type PondHandler struct {
    pondRepository repository.PondRepository
	farmRepository repository.FarmRepository
	logRepository repository.LogRepository
}

func NewPondHandler(
	pondRepository repository.PondRepository, 
	farmRepository repository.FarmRepository,
	logRepository repository.LogRepository,
) *PondHandler {
    return &PondHandler{
        pondRepository: pondRepository,
		farmRepository: farmRepository,
		logRepository: logRepository,
    }
}

func (h *PondHandler) CreatePond(c *gin.Context) {
	// Create log
    log := model.Log{
		Endpoint:  "POST /pond",
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

    var pond model.Pond

    // Bind payload
    if err := c.Bind(&pond); err != nil || pond.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "status": "error",
            "message": "Invalid request payload",
        })
        return
    }

    // Exist pond name
    existPond, _ := h.pondRepository.GetByName(pond.Name)
    if existPond != nil {
        c.JSON(http.StatusConflict, gin.H{
            "code": http.StatusConflict,
            "status": "error",
            "message": "Pond name already exists",
        })
        return
    }

	// Farm data not found
	farm, _ := h.farmRepository.GetById(pond.FarmID)
    if farm == nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusNotFound,
            "status": "error",
            "message": "Farm Data Not Found",
        })
        return
    }

    // Create pond
    if err := h.pondRepository.Create(&pond); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusInternalServerError,
            "status": "error",
            "message": "Failed to create pond",
        })
        return
    }

	// Success
    c.JSON(http.StatusOK, gin.H{
        "code": http.StatusOK,
        "status": "success",
        "message": "New pond created successfully",
        "data": pond,
    })
}

func (h *PondHandler) GetPond(c *gin.Context) {
	// Create log
    log := model.Log{
		Endpoint:  "GET /pond",
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

    pond, _ := h.pondRepository.Get()

    // Empty pond
    if pond == nil {
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
        "message": "Pond fetched successfully",
        "data": pond,
    })
}

func (h *PondHandler) GetPondById(c *gin.Context) {
	// Create log
    log := model.Log{
		Endpoint:  "GET /pond/:id",
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

    pond, _ := h.pondRepository.GetById(id)

    // Empty pond
    if pond == nil {
        c.JSON(http.StatusInternalServerError, gin.H{
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
        "message": "Pond fetched successfully",
        "data": pond,
    })
}

func (h *PondHandler) UpdatePond(c *gin.Context) {
	// Create log
    log := model.Log{
		Endpoint:  "PUT /pond/:id",
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
    var pondPayload model.Pond
    if err := c.Bind(&pondPayload); err != nil || pondPayload.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "code": http.StatusBadRequest,
            "status": "error",
            "message": "Invalid request payload",
        })
        return
    }

    // Exist pond name
    existPond, _ := h.pondRepository.GetByName(pondPayload.Name)
    if existPond != nil && existPond.ID != id {
        c.JSON(http.StatusConflict, gin.H{
            "code": http.StatusConflict,
            "status": "error",
            "message": "Pond name already exists",
        })
        return
    }

	// Farm data not found
	farm, _ := h.farmRepository.GetById(pondPayload.FarmID)
    if farm == nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "code": http.StatusNotFound,
            "status": "error",
            "message": "Farm Data Not Found",
        })
        return
    }

    pond, _ := h.pondRepository.GetById(id)

    // Empty pond
    if pond == nil {
        // Create pond
        if err := h.pondRepository.Create(&pondPayload); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "code": http.StatusInternalServerError,
                "status": "error",
                "message": "Failed to create pond",
            })
            return
        }

        // Success
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusOK,
            "status": "success",
            "message": "Data Not Found. New pond created successfully",
            "data": pondPayload,
        })
    } else {
        // Update pond
        if err := h.pondRepository.Update(pond.ID, &pondPayload); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "code": http.StatusInternalServerError,
                "status": "error",
                "message": "Failed to create pond",
            })
            return
        }
        pondPayload.ID = pond.ID

        // Success
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusOK,
            "status": "success",
            "message": "Pond updated successfully",
            "data": pondPayload,
        })
    }
}

func (h *PondHandler) DeletePond(c *gin.Context) {
	// Create log
    log := model.Log{
		Endpoint:  "DELETE /pond/:id",
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

    pond, _ := h.pondRepository.GetById(id)

    // Empty pond
    if pond == nil {
        c.JSON(http.StatusNotFound, gin.H{
            "code": http.StatusNotFound,
            "status": "error",
            "message": "Data Not Found",
        })
    } else {
        // Delete pond
        if err := h.pondRepository.Delete(pond); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "code": http.StatusInternalServerError,
                "status": "error",
                "message": "Failed to delete pond",
            })
            return
        }

        // Success
        c.JSON(http.StatusOK, gin.H{
            "code": http.StatusOK,
            "status": "success",
            "message": "Pond deleted successfully",
        })
    }
}