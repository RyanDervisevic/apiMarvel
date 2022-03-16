package db

import "github.com/RyanDervisevic/apiMarvel/model"

type Storage struct {
	Heroes StorageHeroes
}

type StorageHeroes interface {
	GetByID(id string) (*model.Heroes, error)
	GetAll() ([]model.Heroes, error)
	DeleteByID(id string) error
	Create(u *model.Heroes) (*model.Heroes, error)
	Update(id string, data map[string]interface{}) (*model.Heroes, error)
}
