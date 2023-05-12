package db

import (
	"WB_L0/internal"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type DB interface {
	Get(uid uuid.UUID) (entity.Entity, error)
	Add(uid uuid.UUID, e entity.Entity) (entity.Entity, error)
	Delete(uid uuid.UUID) error
	//Find(entity.Entity) ([]entity.Entity, error)
}

type PostgresDB struct {
	ctx   context.Context
	conn  *pgx.Conn
	table string
}

func (db *PostgresDB) Get(uid uuid.UUID) (entity.Entity, error) {
	var order entity.Order
	row := db.conn.QueryRow(db.ctx, "SELECT data FROM "+db.table+" WHERE uid=$1;", uid)
	err := row.Scan(&order)
	fmt.Println(db.conn.Close(db.ctx))
	return entity.Entity(order), err
}

func (db *PostgresDB) Add(uid uuid.UUID, e entity.Entity) (entity.Entity, error) {
	//defer db.conn.Close(db.ctx)
	_, err := db.conn.Exec(db.ctx, "INSERT INTO "+db.table+" (uid, data) VALUES ($1, $2);", uid, e.(entity.Order))
	return e, err
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
