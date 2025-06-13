package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"

	"menusPlanningApp/api/controller"
	"menusPlanningApp/api/repository"
	"menusPlanningApp/api/router"
	"menusPlanningApp/api/service"
	"menusPlanningApp/config"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	c := config.New()

	var logLevel gormlogger.LogLevel
	if c.DB.Debug {
		logLevel = gormlogger.Info
	} else {
		logLevel = gormlogger.Error
	}

	dbString := fmt.Sprintf(fmtDBString, c.DB.Host, c.DB.Username, c.DB.Password, c.DB.DBName, c.DB.Port)
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: gormlogger.Default.LogMode(logLevel)})
	if err != nil {
		log.Fatal("DB connection start failure")
		return
	}

	dishRepo := repository.NewDishRepository(db)
	dishService := service.NewDishService(dishRepo)
	dishController := controller.NewDishController(dishService)

	planningRepo := repository.NewPlanningRepository(db)
	planningService := service.NewPlanningService(planningRepo)
	planningController := controller.NewPlanningController(planningService)

	appRouter := &router.AppRouter{
		DishController:     dishController,
		PlanningController: planningController,
	}

	r := router.New(appRouter)
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
