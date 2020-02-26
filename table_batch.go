package edb

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type BatchInsertSqlBuilder = func(tableName, fields, valuePlaceholders string) string

func batchInsertWithRawSql(db *gorm.DB, objArr []interface{}, builder BatchInsertSqlBuilder) error {
	if len(objArr) == 0 {
		return nil
	}

	mainObj := objArr[0]
	mainScope := db.NewScope(mainObj)
	mainFields := mainScope.Fields()
	quoted := make([]string, 0, len(mainFields))
	for i := range mainFields {
		// If primary key has blank value (0 for int, "" for string, nil for interface ...), skip it.
		// If field is ignore field, skip it.
		if (mainFields[i].IsPrimaryKey && mainFields[i].IsBlank) || (mainFields[i].IsIgnored) || (mainFields[i].Relationship != nil) {
			continue
		}
		quoted = append(quoted, mainScope.Quote(mainFields[i].DBName))
	}

	placeholdersArr := make([]string, 0, len(objArr))

	for _, obj := range objArr {
		scope := db.NewScope(obj)
		fields := scope.Fields()
		placeholders := make([]string, 0, len(fields))
		for i := range fields {
			if (fields[i].IsPrimaryKey && fields[i].IsBlank) || (fields[i].IsIgnored) || (fields[i].Relationship != nil) {
				continue
			}
			placeholders = append(placeholders, scope.AddToVars(fields[i].Field.Interface()))
		}
		placeholdersStr := "(" + strings.Join(placeholders, ", ") + ")"
		placeholdersArr = append(placeholdersArr, placeholdersStr)
		// add real variables for the replacement of placeholders' '?' letter later.
		mainScope.SQLVars = append(mainScope.SQLVars, scope.SQLVars...)
	}

	rawSql := builder(
		mainScope.QuotedTableName(),
		strings.Join(quoted, ", "),
		strings.Join(placeholdersArr, ", "),
	)

	mainScope.Raw(rawSql)

	result, err := mainScope.SQLDB().Exec(mainScope.SQL, mainScope.SQLVars...)
	if err != nil {
		return err
	}
	row, _ := result.RowsAffected()
	if row > 0 {
		fmt.Printf("mainScope.TableName():%s,ROW:%d,quoted=%v,sql_vars:%v\n", mainScope.TableName(), row, quoted, mainScope.SQLVars)
	}
	return nil
}

/*
** 批量插入数据，此方法依赖表的table name，因此只能做单表的批量插入
 */
func BatchInsertTable(db *gorm.DB, objArr []interface{}, ignoreDuplicate bool) error {
	return batchInsertWithRawSql(db, objArr, func(tableName, fields, valuePlaceholders string) string {
		var sql string
		if ignoreDuplicate {
			sql = "INSERT IGNORE INTO %s (%s) VALUES %s"
		} else {
			sql = "INSERT INTO %s (%s) VALUES %s"
		}
		return fmt.Sprintf(sql, tableName, fields, valuePlaceholders)
	})
}

/*
** 批量更新数据，如果唯一键冲突，将先删除原有行并插入新行，性能会降低，此方法依赖表的table name，因此只能做单表的批量插入
 */
func BatchReplaceInsertTable(db *gorm.DB, objArr []interface{}) error {
	return batchInsertWithRawSql(db, objArr, func(tableName, fields, valuePlaceholders string) string {
		return fmt.Sprintf("REPLACE INTO %s (%s) VALUES %s", tableName, fields, valuePlaceholders)
	})
}
