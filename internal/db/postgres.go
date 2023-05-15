package db

import (
	"WB_L0/internal/orders"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type PostgresDB struct {
	ctx   context.Context
	conn  *pgx.Conn
	table string
}

func (db *PostgresDB) Get(uid uuid.UUID) (orders.Order, error) {
	order := orders.Order{}
	row := db.conn.QueryRow(db.ctx, "SELECT uid, data FROM "+db.table+" WHERE uid=$1;", uid)

	err := row.Scan(&order.UID, &order.Data)

	return order, err
}

func (db *PostgresDB) Add(uid uuid.UUID, data []byte) (orders.Order, error) {
	order := orders.Order{
		UID:  uid,
		Data: data,
	}
	_, err := db.conn.Exec(db.ctx, "INSERT INTO "+db.table+" (uid, data) VALUES ($1, $2);", uid, data)
	return order, err
}

func (db *PostgresDB) Delete(uid uuid.UUID) error {
	_, err := db.conn.Exec(db.ctx, "DELETE FROM "+db.table+" uid=$1;", uid)
	return err
}

func NewPostgresDB(c context.Context, conn *pgx.Conn, table string) *PostgresDB {
	return &PostgresDB{
		ctx:   c,
		conn:  conn,
		table: table,
	}
}
