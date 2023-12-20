package db

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type SQLBuilder struct {
	db    *gorm.DB
	query string
}

func NewSQLBuilder(db *gorm.DB) *SQLBuilder {
	return &SQLBuilder{
		db: db,
	}
}

func (sb *SQLBuilder) Select(columns ...string) *SQLBuilder {
	sb.query = fmt.Sprintf("SELECT %s", join(columns, ", "))
	return sb
}

func (sb *SQLBuilder) From(table string) *SQLBuilder {
	sb.query = fmt.Sprintf("%s FROM %s", sb.query, table)
	return sb
}

func (sb *SQLBuilder) Where(condition string) *SQLBuilder {
	sb.query = fmt.Sprintf("%s WHERE %s", sb.query, condition)
	return sb
}

func (sb *SQLBuilder) Insert(table string, values map[string]interface{}) *SQLBuilder {
	columns := make([]string, 0, len(values))
	params := make([]interface{}, 0, len(values)*2)

	for column, value := range values {
		columns = append(columns, column)
		params = append(params, value)
	}

	sb.query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, join(columns, ", "), placeholders(len(columns)))
	sb.db = sb.db.Exec(sb.query, params...)
	return sb
}

func (sb *SQLBuilder) Update(table string, values map[string]interface{}) *SQLBuilder {
	columnValuePairs := make([]string, 0, len(values))
	params := make([]interface{}, 0, len(values)*2)

	for column, value := range values {
		columnValuePairs = append(columnValuePairs, fmt.Sprintf("%s = ?", column))
		params = append(params, value)
	}

	sb.query = fmt.Sprintf("UPDATE %s SET %s", table, join(columnValuePairs, ", "))
	sb.db = sb.db.Exec(sb.query, params...)
	return sb
}

func (sb *SQLBuilder) Query() *gorm.DB {
	return sb.db.Raw(sb.query)
}

func placeholders(count int) string {
	return strings.Join(strings.Split(strings.Repeat("?", count), ""), ", ")
}

func join(elements []string, delimiter string) string {
	return strings.Join(elements, delimiter)
}
