package db

import (
	"context"

	"github.com/go-pg/pg/v10"
)

// Logger logs queries when they are run
type Logger struct{}

// BeforeQuery runs before a query is executed
func (d Logger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

// AfterQuery runs after a query is executed
func (d Logger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	// query, _ := q.FormattedQuery()

	// fmt.Println(string(query))
	return nil
}
