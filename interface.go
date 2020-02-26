package edb

const (
	InsertPrefix  = "INSERT INTO "
	InsertIgnore  = "INSERT IGNORE INTO "
	InsertReplace = "REPLACE INTO "
	OrmValue = "(?)"
)

type DbMessage interface {
	TableName() string
	Column() string
	Values() interface{}
}
