package implementation

import (
	"fmt"
	"strings"
	"time"

	"github.com/Edwinfpirajan/server.git/internal/domain/dto"
	"gorm.io/gorm"
)

// SQLBuilder is a struct that will hold the database connection options
type SQLBuilder struct {
	db       *gorm.DB
	query    string
	paginate dto.Paginate
	alias    string
}

type FilterCondition struct {
	Field    string
	Operator string
	Value    interface{}
}

// Paginate represents the information for pagination

// NewSQLBuilder creates a new instance of SQLBuilder
func NewSQLBuilder(db *gorm.DB) *SQLBuilder {
	return &SQLBuilder{
		db: db,
	}
}

// Select sets the SELECT part of the SQL query
func (sb *SQLBuilder) Select(columns ...string) *SQLBuilder {
	// Utilizar el alias almacenado en la cláusula SELECT si está disponible
	if sb.alias != "" {
		columns = prefixColumns(columns, sb.alias)
	}
	sb.query = fmt.Sprintf("SELECT %s", join(columns, ", "))
	return sb
}

// prefixColumns agrega el alias a cada columna en la lista
func prefixColumns(columns []string, alias string) []string {
	prefixedColumns := make([]string, len(columns))
	for i, col := range columns {
		prefixedColumns[i] = fmt.Sprintf("%s.%s", alias, col)
	}
	return prefixedColumns
}

// From sets the FROM part of the SQL query
func (sb *SQLBuilder) From(table string, alias ...string) *SQLBuilder {
	if len(alias) > 0 {
		sb.query = fmt.Sprintf("%s FROM %s %s", sb.query, table, alias[0])
	} else {
		sb.query = fmt.Sprintf("%s FROM %s", sb.query, table)
	}
	return sb
}

// Where sets the WHERE part of the SQL query
func (sb *SQLBuilder) Where(condition string) *SQLBuilder {
	sb.query = fmt.Sprintf("%s WHERE %s", sb.query, condition)
	return sb
}

// Where sets the WHERE part of the SQL query
func BuildFilters(field string, value interface{}, op string, where *string) {
	switch v := value.(type) {
	case string:
		buildStringFilter(field, v, op, where)
	case time.Time:
		buildTimeFilter(field, v, op, where)
	case *bool:
		buildBoolFilter(field, v, op, where)
	default:
		// Puedes manejar otros tipos de valores según sea necesario
	}
}

// buildStringFilter construye una cláusula WHERE para una columna de tipo string
func buildStringFilter(field, value, op string, where *string) {
	buildFilter(field, fmt.Sprintf("%%%s%%", value), op, "ILIKE", where)
}

// buildTimeFilter construye una cláusula WHERE para una columna de tipo time.Time
func buildTimeFilter(field string, value time.Time, op string, where *string) {
	if !value.IsZero() {
		buildFilter(field, value.Format("2006-01-02"), op, "=", where)
	}
}

// buildBoolFilter construye una cláusula WHERE para una columna de tipo *bool
func buildBoolFilter(field string, value *bool, op string, where *string) {
	if value != nil {
		buildFilter(field, fmt.Sprintf("%v", *value), op, "=", where)
	}
}

// buildFilter construye una cláusula WHERE para una columna de tipo string
func buildFilter(field, value, op, comparison string, where *string) {
	if len(*where) > 0 {
		*where += fmt.Sprintf(" %s %s %s", op, field, comparison, value)
	} else {
		*where = fmt.Sprintf("%s %s %s", field, comparison, value)
	}
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

func (sb *SQLBuilder) Join(joinType, table, on, alias string) *SQLBuilder {
	sb.query = fmt.Sprintf("%s %s %s ON %s", sb.query, joinType, table, on)
	// Almacenar el alias para su uso posterior
	sb.alias = alias
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
func (sb *SQLBuilder) SetPaginate(p dto.Paginate) *SQLBuilder {
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
