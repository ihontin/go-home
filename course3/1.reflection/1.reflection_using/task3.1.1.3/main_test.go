package main

import (
	sq "github.com/Masterminds/squirrel"
	"reflect"
	"testing"
)

func TestDAO_BuildSelect(t *testing.T) {
	type fields struct {
		sqlBuilder sq.StatementBuilderType
	}
	type args struct {
		tableName string
		condition Condition
		fields    []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   []interface{}
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				sqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
			},
			args: args{tableName: "users", condition: Condition{
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
			}, fields: []string{"id", "username"}},
			want:    "SELECT id, username FROM users WHERE username = $1 ORDER BY id ASC LIMIT 3 OFFSET 5",
			want1:   []interface{}{"test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DAO{
				sqlBuilder: tt.fields.sqlBuilder,
			}
			got, got1, err := d.BuildSelect(tt.args.tableName, tt.args.condition, tt.args.fields...)
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildSelect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BuildSelect() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("BuildSelect() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
