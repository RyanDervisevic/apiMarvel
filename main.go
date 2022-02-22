package main

import (
	"github.com/RyanDervisevic/apiMarvel/db/moke"
	"github.com/RyanDervisevic/apiMarvel/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := moke.New()
	s := service.New(db)
	r.GET("/heroes/:id", s.GetHeroes)
	r.POST("/heroes", s.CreateHeroes)
	r.GET("/heroes", s.GetAllHeroes)
	// r.PUT("/heroes/:id", s.UpdateHeroes)
	r.DELETE("/heroes/:id", s.DeleteHeroes)
	r.Run(":8081")
}
