package edb

import (
	"github.com/eager7/edb/tables"
	"github.com/shopspring/decimal"
	"testing"
)

const (
	dbAddr, dbUser, dbPass, dbName = "127.0.0.1:3306", "root", "plainchant", "wallet_funds_database"
)

func TestInsert(t *testing.T) {
	db, err := Initialize(dbAddr, dbUser, dbPass, dbName, 1)
	if err != nil {
		t.Fatal(err)
	}
	r1 := tables.TableTestInfo{Uuid: "r1", Version: 0, Balance: decimal.New(1, 2)}
	r2 := tables.TableTestInfo{Uuid: "r2", Version: 0, Balance: decimal.New(1, 3)}

	if err := BatchInsert(db.LogMode(true), []DbMessage{&r1, &r2}, true); err != nil {
		t.Fatal(err)
	}
}
