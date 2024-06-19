package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	_ = w.Close()
	os.Stdout = old

	expected := "username\nemail\naddress\nstatus\n"
	var bufr bytes.Buffer
	_, _ = bufr.ReadFrom(r)
	if bufr.String() != expected {
		t.Errorf("expected = %s, got = %s", expected, bufr.String())
	}
}

func TestSimpleGetFieldsPointers(t *testing.T) {
	user := User{
		ID:       1,
		Username: "JohnDoe",
		Email:    "johndoe@example.com",
		Address:  "123 Main St",
		Status:   1,
		Delete:   "yes",
	}
	type args struct {
		u interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "ok",
			args: args{user},
			want: make([]interface{}, 0, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			testU := reflect.TypeOf(tt.args.u)
			for i := 0; i < testU.NumField(); i++ {
				field := testU.Field(i)
				if tagNfield, _ := field.Tag.Lookup("db"); tagNfield == "delete" || tagNfield == "id" {
					continue
				}
				tt.want = append(tt.want, &field)
			}

			if got := SimpleGetFieldsPointers(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleGetFieldsPointers() = %v, want %v", got, tt.want)
			}
		})
	}
}
