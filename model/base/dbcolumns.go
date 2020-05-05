package base

import (
	"errors"
)

type DBColumns struct {
	TableSchema            string `gorm:"column:TABLE_SCHEMA" json:"tableSchema"`
	TableName              string `gorm:"column:TABLE_NAME" json:"tableName"`
	ColumnName             string `gorm:"column:COLUMN_NAME" json:"columnName"`
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT" json:"columnDefault"`
	IsNullable             string `gorm:"column:IS_NULLABLE" json:"isNullable"`
	DataType               string `gorm:"column:DATA_TYPE" json:"dataType"`
	CharacterMaximumLength string `gorm:"column:CHARACTER_MAXIMUM_LENGTH" json:"characterMaximumLength"`
	CharacterSetName       string `gorm:"column:CHARACTER_SET_NAME" json:"characterSetName"`
	ColumnType             string `gorm:"column:COLUMN_TYPE" json:"columnType"`
	ColumnKey              string `gorm:"column:COLUMN_KEY" json:"columnKey"`
	Extra                  string `gorm:"column:EXTRA" json:"extra"`
	ColumnComment          string `gorm:"column:COLUMN_COMMENT" json:"columnComment"`
}

func (e *DBColumns) GetPage(pageSize int, pageIndex int) (doc []DBColumns, count int, err error) {

	db := DB.TestDB.Select("*").Table("information_schema.`COLUMNS`")
	db = db.Where("table_schema= ? ", "testdb")

	if e.TableName != "" {
		return nil, 0, errors.New("table name cannot be empty！")
	}

	db = db.Where("TABLE_NAME = ?", e.TableName)
	err = db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Count(&count).Error
	return
}

func (e *DBColumns) GetList() ([]DBColumns, error) {
	var doc []DBColumns

	table := DB.TestDB.Select("*").Table("information_schema.columns")
	table = table.Where("table_schema= ? ", "testdb")

	if e.TableName == "" {
		return nil, errors.New("table name cannot be empty！")
	}

	table = table.Where("TABLE_NAME = ?", e.TableName)

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
