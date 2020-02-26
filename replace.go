package edb

import (
	"github.com/jinzhu/gorm"
	"strings"
)

/*
** 批量替换插入，如果唯一键冲突，将删除旧的数据，插入新数据，可以设定表单名称，从而实现分表的策略
 */
func BatchReplace(db *gorm.DB, batch []DbMessage, tName string) error {
	if len(batch) == 0 {
		return nil
	}
	sql, values := BatchReplaceSql(batch, tName)
	if len(values) == 0 {
		return nil
	}
	if err := db.Exec(sql, values...).Error; err != nil {
		return err
	}
	return nil
}

func Replace(db *gorm.DB, msg DbMessage) error {
	return db.Exec(InsertReplace+msg.TableName()+msg.Column(), msg.Values()).Error
}

func BatchReplaceSql(messages []DbMessage, tName string) (sql string, values []interface{}) {
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
	return InsertReplace + tName + messages[0].Column() + strings.Join(str, ","), values
}
