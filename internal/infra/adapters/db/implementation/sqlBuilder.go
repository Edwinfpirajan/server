package implementation

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// SQLBuilder is a struct that will hold the database connection options
type SQLBuilder struct {
	db    *gorm.DB
	query string
}

// This is a method that will create a single instance of the connection
func NewSQLBuilder(db *gorm.DB) *SQLBuilder {
	return &SQLBuilder{
		db: db,
	}
}

// This is a method that will create a single instance of the connection
func (sb *SQLBuilder) Select(columns ...string) *SQLBuilder {
	sb.query = fmt.Sprintf("SELECT %s", join(columns, ", "))
	return sb
}

// This is a method that will create a single instance of the connection
func (sb *SQLBuilder) From(table string) *SQLBuilder {
	sb.query = fmt.Sprintf("%s FROM %s", sb.query, table)
	return sb
}

// This is a method that will create a single instance of the connection
func (sb *SQLBuilder) Where(condition string) *SQLBuilder {
	sb.query = fmt.Sprintf("%s WHERE %s", sb.query, condition)
	return sb
}

// This is a method that will create a single instance of the connection
func (sb *SQLBuilder) Insert(table string, values map[string]interface{}) *SQLBuilder {
	columns := make([]string, 0, len(values))
	params := make([]interface{}, 0, len(values)*2)

	// This is a method that will create a single instance of the connection
	for column, value := range values {
		columns = append(columns, column)
		params = append(params, value)
	}

	// This is a method that will create a single instance of the connection
	sb.query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, join(columns, ", "), placeholders(len(columns)))
	sb.db = sb.db.Exec(sb.query, params...)
	return sb
}

// This is a method that will create a single instance of the connection
func (sb *SQLBuilder) Update(table string, values map[string]interface{}) *SQLBuilder {
	columnValuePairs := make([]string, 0, len(values))
	params := make([]interface{}, 0, len(values)*2)

	// This is a method that will create a single instance of the connection
	for column, value := range values {
		columnValuePairs = append(columnValuePairs, fmt.Sprintf("%s = ?", column))
		params = append(params, value)
	}

	// This is a method that will create a single instance of the connection
	sb.query = fmt.Sprintf("UPDATE %s SET %s", table, join(columnValuePairs, ", "))
	sb.db = sb.db.Exec(sb.query, params...)
	return sb
}

// This is a method that will create a single instance of the connection
func (sb *SQLBuilder) Query() *gorm.DB {
	return sb.db.Raw(sb.query)
}

// This is a method that will create a single instance of the connection
func placeholders(count int) string {
	return strings.Join(strings.Split(strings.Repeat("?", count), ""), ", ")
}

// This is a method that will create a single instance of the connection
func join(elements []string, delimiter string) string {
	return strings.Join(elements, delimiter)
}
