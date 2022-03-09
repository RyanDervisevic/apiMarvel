package sqlite

import (
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/RyanDervisevic/apiMarvel/db"
	"github.com/RyanDervisevic/apiMarvel/model"
)

type SQLite struct {
	Conn *gorm.DB
}

func New(dbName string) *db.Storage {
	conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = conn.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	return &db.Storage{
		Heroes: &SQLite{
			Conn: conn,
		},
	}
}

func (c *SQLite) GetByID(id string) (*model.Heroes, error) {
	var u model.Heroes
	err := c.Conn.First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// func (c *SQLite) GetByEmail(email string) (*model.User, error) {
// 	var u model.User
// 	err := c.Conn.First(&u, "email = ?", email).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &u, nil
// }

func (c *SQLite) GetAll() ([]model.Heroes, error) {
	var us []model.Heroes
	err := c.Conn.Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (c *SQLite) DeleteByID(id string) error {
	return c.Conn.Where("id = ?", id).Delete(&model.Heroes{}).Error
}

func (c *SQLite) Create(u *model.Heroes) (*model.Heroes, error) {
	u.ID = uuid.NewString()
	err := c.Conn.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (c *SQLite) Update(id string, data map[string]interface{}) (*model.Heroes, error) {
	u := model.Heroes{ID: id}
	err := c.Conn.Model(&u).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return c.GetByID(id)
}
