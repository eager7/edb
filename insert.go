package edb

import (
	"github.com/jinzhu/gorm"
	"strings"
)

/*
** 批量插入，可以设定表单名称，从而实现分表的策略
 */
func BatchInsert(db *gorm.DB, batch []DbMessage, tName string, ignore bool) error {
	if len(batch) == 0 {
		return nil
	}
	sql, values := BatchInsertSql(batch, tName, ignore)
	if len(values) == 0 {
		return nil
	}
	if err := db.Exec(sql, values...).Error; err != nil {
		return err
	}
	return nil
}

func Insert(db *gorm.DB, msg DbMessage, ignore bool) error {
	if ignore {
		return db.Exec(InsertIgnore+msg.TableName()+msg.Column(), msg.Values()).Error
	} else {
		return db.Exec(InsertPrefix+msg.TableName()+msg.Column(), msg.Values()).Error
	}
}

func BatchInsertSql(messages []DbMessage, tName string, ignore bool) (sql string, values []interface{}) {
	if len(messages) == 0 {
		return "", nil
	}
	var str []string
	for _, msg := range messages {
		values = append(values, msg.Values())
		str = append(str, OrmValue)
	}
	if tName == "" {
		tName = messages[0].TableName()
	}
	if ignore {
		return InsertIgnore + tName + messages[0].Column() + strings.Join(str, ","), values
	} else {
		return InsertPrefix + tName + messages[0].Column() + strings.Join(str, ","), values
	}
}
