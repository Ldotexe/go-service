package db

import (
	"context"
	"sync"
	"testing"

	"homework-6/internal/service/db"
)

type TDB struct {
	DB db.DBops
	sync.Mutex
}

func NewFromEnv() *TDB {
	//запрашиваем тестовые креды для бд из енв
	//cfg, err := config.FromEnv()

	db, err := db.NewDB(context.Background())
	if err != nil {
		panic(err)
	}
	return &TDB{DB: db}
}

func (d *TDB) SetUp(t *testing.T) {
	t.Helper()
	d.Lock()
	d.Truncate(context.Background())
}

func (d *TDB) TearDown() {
	defer d.Unlock()
	d.Truncate(context.Background())
}

func (d *TDB) Truncate(ctx context.Context) {
	if _, err := d.DB.Exec(ctx, "Truncate table students"); err != nil {
		panic(err)
	}

}
