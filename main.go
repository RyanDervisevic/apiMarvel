package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := moke.New()
	s := service.New(db)
	r.GET("/users/:id", s.GetHeroe)
	r.POST("/users", s.CreateHeroe)
	r.GET("/users", s.GetAllHeroes)
	r.DELETE("/users/:id", s.DeleteHeroe)
	r.Run(":8081")
}
