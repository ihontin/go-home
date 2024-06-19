package main

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
)

type Condition struct {
	Equal       map[string]interface{}
	NotEqual    map[string]interface{}
	Order       []*Order
	LimitOffset *LimitOffset
	ForUpdate   bool
	Upsert      bool
}

type Order struct {
	Field string
	Asc   bool
}

type LimitOffset struct {
	Offset int64
	Limit  int64
}

type DAO struct {
	sqlBuilder sq.StatementBuilderType
}

func NewDAO() *DAO {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &DAO{sqlBuilder: builder}
}

func (d *DAO) BuildSelect(tableName string, condition Condition, fields ...string) (string, []interface{}, error) {
	queryRaw := d.sqlBuilder.Select(fields...).From(tableName)

	for k, v := range condition.Equal {
		queryRaw = queryRaw.Where(sq.Eq{k: v})
	}

	for k, v := range condition.NotEqual {
		queryRaw = queryRaw.Where(sq.NotEq{k: v})
	}

	if condition.LimitOffset != nil {
		queryRaw = queryRaw.Limit(uint64(condition.LimitOffset.Limit)).Offset(uint64(condition.LimitOffset.Offset))
	}

	for _, order := range condition.Order {
		direction := "ASC"
		if !order.Asc {
			direction = "DESC"
		}
		queryRaw = queryRaw.OrderBy(fmt.Sprintf("%s %s", order.Field, direction))
	}

	return queryRaw.ToSql()
}

func main() {
	d := NewDAO()
	s, _, err := d.BuildSelect("users", Condition{
		Equal: map[string]interface{}{
			"username": "test",
		},
		LimitOffset: &LimitOffset{
			Offset: 5,
			Limit:  3,
		},
		Order: []*Order{
			{
				Field: "id",
				Asc:   true,
			},
		},
	}, "id", "username")

	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}

//Output:
//SELECT id, username FROM users WHERE username = $1 ORDER BY id ASC LIMIT 3 OFFSET 5
