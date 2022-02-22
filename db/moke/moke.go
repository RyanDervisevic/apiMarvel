package moke

import (
	"errors"

	"github.com/google/uuid"

	"github.com/RyanDervisevic/apiMarvel/db"
	"github.com/RyanDervisevic/apiMarvel/model"
)

var _ db.StorageHeroes = &Moke{}

type Moke struct {
	listHeroes map[string]*model.Heroes
}

func New() *db.Storage {
	return &db.Storage{
		Heroes: &Moke{
			listHeroes: make(map[string]*model.Heroes),
		},
	}
}

func (m *Moke) GetByID(id string) (*model.Heroes, error) {
	h, ok := m.listHeroes[id]
	if !ok {
		return nil, errors.New("db Heroes: not found")
	}
	return h, nil
}

func (m *Moke) DeleteByID(id string) error {
	_, ok := m.listHeroes[id]
	if !ok {
		return errors.New("db Heroes: not found")
	}
	delete(m.listHeroes, id)
	return nil
}

func (m *Moke) Create(h *model.Heroes) (*model.Heroes, error) {
	h.ID = uuid.New().String()
	m.listHeroes[h.ID] = h
	return h, nil
}

func (m *Moke) Update(id string, data map[string]interface{}) (*model.Heroes, error) {
	h, ok := m.listHeroes[id]
	if !ok {
		return nil, errors.New("db Heroes: not found")
	}
	if value, ok := data["first_name"]; ok {
		h.FirstName = value.(string)
	}
	if value, ok := data["last_name"]; ok {
		h.FirstName = value.(string)
	}
	return h, nil
}

func (m *Moke) GetAll() ([]model.Heroes, error) {
	us := make([]model.Heroes, len(m.listHeroes))
	var i int
	for k := range m.listHeroes {
		if m.listHeroes[k] != nil {
			us[i] = *m.listHeroes[k]
		}
		i++
	}
	return us, nil
}
