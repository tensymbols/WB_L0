package db

import (
	"WB_L0/internal/orders"
	"github.com/google/uuid"
)

type DB interface {
	Get(uid uuid.UUID) (orders.Order, error)
	Add(uid uuid.UUID, data []byte) (orders.Order, error)
	Delete(uid uuid.UUID) error
	//Find(entity.Entity) ([]entity.Entity, error)
}
