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
	if err := Insert(db.LogMode(true), &r1, true); err != nil {
		t.Fatal(err)
	}

	r2 := tables.TableTestInfo{Uuid: "r2", Version: 0, Balance: decimal.New(1, 3)}

	if err := BatchInsert(db.LogMode(true), []DbMessage{&r1, &r2}, (&tables.TableTestInfo{}).TableName(), true); err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	db, err := Initialize(dbAddr, dbUser, dbPass, dbName, 1)
	if err != nil {
		t.Fatal(err)
	}
	r1 := tables.TableTestInfo{Uuid: "r1", Version: 0, Balance: decimal.New(2, 2)}
	if err := Update(db.LogMode(true), &r1, "balance"); err != nil {
		t.Fatal(err)
	}

	r2 := tables.TableTestInfo{Uuid: "r2", Version: 0, Balance: decimal.New(2, 3)}

	if err := BatchUpdate(db.LogMode(true), []DbMessage{&r1, &r2}, (&tables.TableTestInfo{}).TableName(), "balance"); err != nil {
		t.Fatal(err)
	}
}

func TestReplace(t *testing.T) {
	db, err := Initialize(dbAddr, dbUser, dbPass, dbName, 1)
	if err != nil {
		t.Fatal(err)
	}
	r1 := tables.TableTestInfo{Uuid: "r1", Version: 0, Balance: decimal.New(3, 2)}
	if err := Replace(db.LogMode(true), &r1); err != nil {
		t.Fatal(err)
	}

	r2 := tables.TableTestInfo{Uuid: "r2", Version: 0, Balance: decimal.New(3, 3)}

	if err := BatchReplace(db.LogMode(true), []DbMessage{&r1, &r2}, (&tables.TableTestInfo{}).TableName()); err != nil {
		t.Fatal(err)
	}
}

func TestBatchInsertTable(t *testing.T) {
	db, err := Initialize(dbAddr, dbUser, dbPass, dbName, 1)
	if err != nil {
		t.Fatal(err)
	}
	r1 := tables.TableTestInfo{Uuid: "r3", Version: 0, Balance: decimal.New(4, 2)}
	r2 := tables.TableTestInfo{Uuid: "r4", Version: 0, Balance: decimal.New(4, 3)}

	if err := BatchInsertTable(db.LogMode(true), []interface{}{&r1, &r2}, true); err != nil {
		t.Fatal(err)
	}
}
