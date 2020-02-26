package tables

import (
	"github.com/shopspring/decimal"
)

type TableTestInfo struct {
	Id      uint64          `json:"id"                gorm:"column:id;primary_key;AUTO_INCREMENT"` //自增主键
	Uuid    string          `json:"uuid"              gorm:"column:uuid"`                          //用于唯一标识
	Version uint64          `json:"version"           gorm:"column:version"`                       //版本号
	Balance decimal.Decimal `json:"balance"           gorm:"column:balance"`                       //余额
}

func (t *TableTestInfo) TableName() string {
	return "t_assert_info"
}

func (t *TableTestInfo) Column() string {
	return "(`uuid`,`balance`,`version`) VALUES"
}

func (t *TableTestInfo) Values() interface{} {
	return []interface{}{t.Uuid, t.Balance, t.Version}
}
