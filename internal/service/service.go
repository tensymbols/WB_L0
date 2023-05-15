package service

import (
	"WB_L0/internal/db"
	"WB_L0/internal/orders"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log"
)

type Service interface {
	AddOrder(order orders.Order) (orders.Order, error)
	DeleteOrder(uid uuid.UUID) error
	GetOrder(uid uuid.UUID) (orders.Order, error)
}

type WBService struct {
	storage db.DB
	cache   db.DB
}

func (a *WBService) AddOrder(order orders.Order) (orders.Order, error) {

	o, err := a.storage.Add(order.UID, order.Data)
	if err != nil {
		return o, err
	}
	_, err = a.cache.Add(order.UID, order.Data)
	return o, err
}

func (a *WBService) DeleteOrder(uid uuid.UUID) error {
	err := a.cache.Delete(uid)
	if err != nil {
		return err
	}
	return a.storage.Delete(uid)
}
func (a *WBService) GetOrder(uid uuid.UUID) (orders.Order, error) {
	o, err := a.storage.Get(uid)
	return o, err
}

func NewService(ctx context.Context, conn *pgx.Conn, table string) *WBService {
	pgDB := db.NewPostgresDB(ctx, conn, table)
	cacheDB, err := db.NewCacheDB(pgDB)
	if err != nil {
		log.Fatalf("could not initialize cache: %v", err)
	}
	return &WBService{
		storage: pgDB,
		cache:   cacheDB,
	}
}
