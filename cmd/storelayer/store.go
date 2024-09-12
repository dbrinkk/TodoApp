package storelayer

import (
	"gorm.io/gorm"
)

type Store interface {
}

type store struct {
	db *gorm.DB
}

func New() *store {
	return &store{}
}
