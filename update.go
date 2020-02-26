package edb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

func BatchUpdate(db *gorm.DB, batch []DbMessage, tName string, values ...string) error {
	if len(batch) == 0 {
		return nil
	}
	var vStr []string
	for _, v := range values {
		vStr = append(vStr, fmt.Sprintf("%s=VALUES(%s)", v, v))
	}
	sql, vs := BatchInsertSql(batch, tName, false)
	sql += "ON DUPLICATE KEY UPDATE " + strings.Join(vStr, ",")
	if err := db.Exec(sql, vs).Error; err != nil {
		return err
	}
	return nil
}

func Update(db *gorm.DB, msg DbMessage, values ...string) error {
	var vStr []string
	for _, v := range values {
		vStr = append(vStr, fmt.Sprintf("%s=VALUES(%s)", v, v))
	}
	sql := InsertPrefix + msg.TableName() + msg.Column() + "ON DUPLICATE KEY UPDATE " + strings.Join(vStr, ",")
	return db.Exec(sql, msg.Values()).Error
}
