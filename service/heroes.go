package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RyanDervisevic/apiMarvel/db"
	"github.com/RyanDervisevic/apiMarvel/model"
)

type Service struct {
	db      *db.Storage
	signKey []byte
}

//func New(db *db.Storage, signKey []byte) *Service {
func New(db *db.Storage) *Service {
	return &Service{
		db: db,
	}
}

// go to Service folder.
func (s *Service) GetHeroes(c *gin.Context) {
	id := c.Param("id")
	h, err := s.db.Heroes.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"heroes": h,
	})
}

func (s *Service) GetAllHeroes(c *gin.Context) {
	h, err := s.db.Heroes.GetAll()
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Heroes": h,
	})
}

func (s *Service) CreateHeroes(c *gin.Context) {
	var h model.Heroes
	err := c.BindJSON(&h)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	_, err = s.db.Heroes.Create(&h)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Heroes": h,
	})
}

// func (s *Service) UpdateHeroes(c *gin.Context) {
// 	id := c.Param("id")
// 	var h model.Heroes
// 	err := c.BindJSON(&h)
// 	if len(id) == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error id": id,
// 		})
// 		return
// 	}
// 	_, err = s.db.Heroes.Update(id,)
// 	if err != nil {
// 		log.Println("service:", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "error internal",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusAccepted, gin.H{
// 		"updated": id,
// 	})
// }

func (s *Service) DeleteHeroes(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error id": id,
		})
		return
	}
	err := s.db.Heroes.DeleteByID(id)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"delete": id,
	})
}

// func (s *Service) Login(c *gin.Context) {

// 	var l model.LoginUser
// 	err := c.BindJSON(&l)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"err": err,
// 		})
// 		return
// 	}

// 	u, err := s.db.User.GetByEmail(l.Email)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"email": l.Email,
// 		})
// 		return
// 	}

// 	log.Printf("receive %v - got %v", *l.Password, *u.Password)

// 	if *u.Password != *l.Password {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "not authorized",
// 		})
// 		return
// 	}

// 	jwtVal, err := util.CreateJWT(s.signKey, u.ID, u.FirstName)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, nil)
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"jwt": jwtVal,
// 	})
// }
