package dao

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"studentgit.kata.academy/Alkolex/go-kata/course3/1.reflection/1.reflection_using/task3.1.1.4/tabler"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=./dao.go -destination=../mock/mock.go -package=mock
type IfaceDAO interface {
	BuildSelect(tableName string, condition Condition, fields ...string) (string, []interface{}, error)
	Create(ctx context.Context, entity tabler.Tabler, opts ...interface{}) error
	List(ctx context.Context, dest interface{}, table tabler.Tabler, condition Condition, opts ...interface{}) error
	Update(ctx context.Context, entity tabler.Tabler, condition Condition, opts ...interface{}) error
}

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

// DAO data access object
type DAO struct {
	db         *sqlx.DB
	sqlBuilder sq.StatementBuilderType
}

func NewDAO(db *sqlx.DB) IfaceDAO {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &DAO{db: db, sqlBuilder: builder}
}

func (d *DAO) BuildSelect(tableName string, condition Condition, fields ...string) (string, []interface{}, error) {
	var queryRaw sq.SelectBuilder
	if len(fields) < 1 {
		queryRaw = d.sqlBuilder.Select().From(tableName)
	} else {
		queryRaw = d.sqlBuilder.Select(fields...).From(tableName)
	}

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

func filterByTag(tag string, tvalue string) func(fields *[]reflect.StructField) {
	return tabler.FilterByTags(map[string]func(value string) bool{
		tag: func(value string) bool {
			return strings.Contains(value, tvalue)
		},
	})
}

func (d *DAO) Create(ctx context.Context, entity tabler.Tabler, opts ...interface{}) error {
	val := reflect.ValueOf(entity).Elem()
	var columns = make([]string, 0, val.NumField())
	var values = make([]interface{}, 0, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		//if val.Field(i).IsZero() { // если значение отсутствует заполнит поля db - NILL
		//	continue
		//}
		column := val.Type().Field(i).Tag.Get("db")
		value := val.Field(i).Interface()
		columns = append(columns, column)
		values = append(values, value)
	}
	query := d.sqlBuilder.Insert(entity.TableName()).Columns(columns...).Values(values...)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = d.db.ExecContext(ctx, sql, args...)
	return err
}

func (d *DAO) List(ctx context.Context, dest interface{}, table tabler.Tabler, condition Condition, opts ...interface{}) error {
	valType := reflect.TypeOf(table).Elem()
	var ops = make([]string, 0, valType.NumField())

	for i := 0; i < valType.NumField(); i++ {
		ops = append(ops, valType.Field(i).Tag.Get("db"))
	}
	sql, args, err := d.BuildSelect(table.TableName(), condition, ops...)
	if err != nil {
		return err
	}
	err = d.db.SelectContext(ctx, dest, sql, args...)
	return err
}

func (d *DAO) Update(ctx context.Context, entity tabler.Tabler, condition Condition, opts ...interface{}) error {
	var idTest int
	val := reflect.ValueOf(entity).Elem()
	var setMap = make(map[string]interface{}, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		//if val.Field(i).IsZero() { // если значение отсутствует заполнит поля db - NILL
		//	continue
		//}
		column := val.Type().Field(i).Tag.Get("db")
		value := val.Field(i).Interface()
		if column == "id" {
			idTest = value.(int)
		}
		setMap[column] = value
	}

	// Проверка наличия записи в базе данных
	var exists bool
	var selects = "SELECT EXISTS"
	err := d.db.QueryRowxContext(ctx, selects+"(SELECT 1 FROM "+entity.TableName()+" WHERE id = ?)", idTest).Scan(&exists)
	if err != nil {
		return err
	}
	// Если запись не существует, выполняем вставку
	if !exists {
		sql, args, err2 := d.sqlBuilder.Insert(entity.TableName()).SetMap(setMap).ToSql()
		if err2 != nil {
			return err2
		}
		_, err2 = d.db.ExecContext(ctx, sql, args...)
		return err2
	}
	queryRaw := d.sqlBuilder.Update(entity.TableName()).SetMap(setMap)

	for k, v := range condition.Equal {
		queryRaw = queryRaw.Where(sq.Eq{k: v})
	}

	sql, args, err := queryRaw.ToSql()
	if err != nil {
		return err
	}
	_, err = d.db.ExecContext(ctx, sql, args...)
	return err
}
