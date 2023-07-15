package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/WillyWilsen/Delos-Task-Assignment.git/database"
	"github.com/WillyWilsen/Delos-Task-Assignment.git/utility"
	"github.com/WillyWilsen/Delos-Task-Assignment.git/repository"
	"github.com/WillyWilsen/Delos-Task-Assignment.git/handler"
)

func main() {
	utility.PrintConsole("API started", "info")
	utility.PrintConsole("Loading application configuration", "info")
	configuration, errConfig := utility.LoadApplicationConfiguration("")
	if errConfig != nil {
		log.WithFields(log.Fields{"error": errConfig}).Fatal("Failed to load app configuration")
	}
	utility.PrintConsole("Application configuration loaded successfully", "info")

	db, gormDB, err := database.Open(configuration)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Fatal("Failed to open database")
	}
	defer db.Close()

	// Repository
	farmRepository := repository.NewFarmRepository(gormDB)
	pondRepository := repository.NewPondRepository(gormDB)
	logRepository := repository.NewLogRepository(gormDB)

	// Handler
	farmHandler := handler.NewFarmHandler(farmRepository, logRepository)
	pondHandler := handler.NewPondHandler(pondRepository, farmRepository, logRepository)
	statisticsHandler := handler.NewStatisticsHandler(logRepository)

	// Router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(utility.CORSMiddleware())
	router.Use(gin.Recovery())
	router.NoRoute(func(c *gin.Context) {
		c.JSON(503, gin.H{"status": "error", "message": "Endpoint not found!"})
	})

	farmRouter := router.Group("/api/farm")
	farmRouter.POST("/", farmHandler.CreateFarm)
	farmRouter.GET("/", farmHandler.GetFarm)
	farmRouter.GET("/:id", farmHandler.GetFarmById)
	farmRouter.PUT("/:id", farmHandler.UpdateFarm)
	farmRouter.DELETE("/:id", farmHandler.DeleteFarm)

	pondRouter := router.Group("/api/pond")
	pondRouter.POST("/", pondHandler.CreatePond)
	pondRouter.GET("/", pondHandler.GetPond)
	pondRouter.GET("/:id", pondHandler.GetPondById)
	pondRouter.PUT("/:id", pondHandler.UpdatePond)
	pondRouter.DELETE("/:id", pondHandler.DeletePond)

	statisticsRouter := router.Group("/api/statistics")
	statisticsRouter.GET("/", statisticsHandler.GetStatistics)

	errServer := router.Run(":" + configuration.Http.HttpPort)
	if errServer != nil {
		utility.PrintConsole(fmt.Sprintf("%v", errServer.Error()), "error")
	}
}