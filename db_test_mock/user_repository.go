package dbtestmock

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID   uuid.UUID `gorm:"column:id;primary_key"`
	Name string    `gorm:"column:name"`
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(id uuid.UUID, name string) error {
	p := &User{
		ID:   id,
		Name: name,
	}

	return r.DB.Table("users").Create(p).Error
}

func (r *UserRepository) Get(id uuid.UUID) (*User, error) {
	user := new(User)
	err := r.DB.Where("id = ?", id).Find(user).Error
	return user, err
}
