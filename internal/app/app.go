package app

import (
	entity "WB_L0/internal"
	"WB_L0/internal/db"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type App interface {
	AddOrder(order entity.Order) (entity.Order, error)
	DeleteOrder(uid uuid.UUID) error
	GetOrder(uid uuid.UUID) (entity.Order, error)
}

type WBApp struct {
	storage db.DB
}

func (a *WBApp) AddOrder(orderJSON entity.Order) (entity.Order, error) {
	var order db.OrderModel
	err := json.Unmarshal(orderJSON, &order)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling order json")
	}
	o, err := a.storage.Add(uuid.MustParse(order.OrderUID), orderJSON)
	return o.(entity.Order), err
}

func (a *WBApp) DeleteOrder(uid uuid.UUID) error {
	return a.storage.Delete(uid)
}
func (a *WBApp) GetOrder(uid uuid.UUID) (entity.Order, error) {
	o, err := a.storage.Get(uid)
	return o.(entity.Order), err
}

func NewApp(ctx context.Context, conn *pgx.Conn, table string) *WBApp {
	return &WBApp{
		storage: db.NewPostgresDB(ctx, conn, table),
	}
}
