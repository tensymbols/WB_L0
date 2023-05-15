package db

import (
	"WB_L0/internal/orders"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type CacheDB struct {
	mu   sync.Mutex
	data map[uuid.UUID]orders.Order
}

func (db *CacheDB) Get(uid uuid.UUID) (orders.Order, error) {
	db.mu.Lock()
	order, ok := db.data[uid]
	db.mu.Unlock()
	if !ok {
		return order, fmt.Errorf("value with uid %s doesn't exist", uid.String())
	}
	return order, nil
}

func (db *CacheDB) Add(uid uuid.UUID, data []byte) (orders.Order, error) {
	order := orders.Order{
		UID:  uid,
		Data: data,
	}
	db.mu.Lock()
	db.data[uid] = order
	db.mu.Unlock()
	return order, nil
}

func (db *CacheDB) Delete(uid uuid.UUID) error {
	db.mu.Lock()
	delete(db.data, uid)
	db.mu.Unlock()
	return nil
}

func NewCacheDB(db *PostgresDB) (DB, error) {

	cache := CacheDB{
		data: map[uuid.UUID]orders.Order{},
	}
	rows, err := db.conn.Query(db.ctx, "SELECT * FROM "+db.table+";")
	if err != nil {
		return &cache, err
	}

	var currOrder orders.Order
	for rows.Next() {

		err = rows.Scan(&currOrder.UID, &currOrder.Data)
		if err != nil {
			return &cache, err
		}
		_, err = cache.Add(currOrder.UID, currOrder.Data)
		if err != nil {
			return &cache, err
		}
	}
	return &cache, nil
}
