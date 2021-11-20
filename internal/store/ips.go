package store

import (
	"github.com/maxheckel/censys-assessment/internal/domain"
	"gorm.io/gorm"
)

type IPStore interface {
	GetIP(address string) (domain.IP, error)
}

type ipStore struct {
	db *gorm.DB
}

func NewIPStore(db *gorm.DB) IPStore{
	return ipStore{db}
}

func (i ipStore) GetIP(address string) (domain.IP, error) {
	panic("implement me")
}
