package implementation

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// SQLBuilder is a struct that will hold the database connection options
type SQLBuilder struct {
	db       *gorm.DB
	query    string
	paginate Paginate
}

// Paginate represents the information for pagination
type Paginate struct {
	Page  int `query:"page" mod:"default=1"`
	Limit int `query:"limit" mod:"default=5"`
}

// NewSQLBuilder creates a new instance of SQLBuilder
func NewSQLBuilder(db *gorm.DB) *SQLBuilder {
	return &SQLBuilder{
		db: db,
	}
}

// Select sets the SELECT part of the SQL query
func (sb *SQLBuilder) Select(columns ...string) *SQLBuilder {
	sb.query = fmt.Sprintf("SELECT %s", join(columns, ", "))
	return sb
}

// From sets the FROM part of the SQL query
func (sb *SQLBuilder) From(table string) *SQLBuilder {
	sb.query = fmt.Sprintf("%s FROM %s", sb.query, table)
	return sb
}

// Where sets the WHERE part of the SQL query
func (sb *SQLBuilder) Where(condition string) *SQLBuilder {
	sb.query = fmt.Sprintf("%s WHERE %s", sb.query, condition)
	return sb
}

// Insert sets the INSERT part of the SQL query
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

// Update sets the UPDATE part of the SQL query
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

// SetPaginate sets the pagination information
func (sb *SQLBuilder) SetPaginate(p Paginate) *SQLBuilder {
	sb.paginate = p
	return sb
}

// Query generates the final SQL query and returns a *gorm.DB object
func (sb *SQLBuilder) Query() *gorm.DB {
	// Apply pagination
	if sb.paginate.Page > 0 && sb.paginate.Limit > 0 {
		offset := (sb.paginate.Page - 1) * sb.paginate.Limit
		sb.query = fmt.Sprintf("%s OFFSET %d LIMIT %d", sb.query, offset, sb.paginate.Limit)
	}
	return sb.db.Raw(sb.query)
}

// placeholders generates placeholder strings for SQL queries
func placeholders(count int) string {
	return strings.Join(strings.Split(strings.Repeat("?", count), ""), ", ")
}

// join concatenates a slice of strings with a delimiter
func join(elements []string, delimiter string) string {
	return strings.Join(elements, delimiter)
}
