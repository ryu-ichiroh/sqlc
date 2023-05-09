// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const subqueryWithHavingClause = `-- name: SubqueryWithHavingClause :many
SELECT a, (SELECT COUNT(a) FROM foo GROUP BY b HAVING COUNT(a) > 10) as "total" FROM foo
`

type SubqueryWithHavingClauseRow struct {
	A     int32
	Total sql.NullInt64
}

func (q *Queries) SubqueryWithHavingClause(ctx context.Context) ([]SubqueryWithHavingClauseRow, error) {
	rows, err := q.db.QueryContext(ctx, subqueryWithHavingClause)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SubqueryWithHavingClauseRow
	for rows.Next() {
		var i SubqueryWithHavingClauseRow
		if err := rows.Scan(&i.A, &i.Total); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const subqueryWithWhereClause = `-- name: SubqueryWithWhereClause :many
SELECT a, (SELECT COUNT(a) FROM foo WHERE a > 10) as "total" FROM foo
`

type SubqueryWithWhereClauseRow struct {
	A     int32
	Total sql.NullInt64
}

func (q *Queries) SubqueryWithWhereClause(ctx context.Context) ([]SubqueryWithWhereClauseRow, error) {
	rows, err := q.db.QueryContext(ctx, subqueryWithWhereClause)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SubqueryWithWhereClauseRow
	for rows.Next() {
		var i SubqueryWithWhereClauseRow
		if err := rows.Scan(&i.A, &i.Total); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
