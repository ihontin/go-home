package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	a1 := []User{{5, "1", 9}, {7, "3", 30}}
	a2 := []User{{6, "2", 9}, {8, "4", 30}}

	type args struct {
		arr1 []User
		arr2 []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{"arr1 empty", args{[]User{}, a1}, a1},
		{"arr2 empty", args{a2, []User{}}, a2},
		{"arrays empty", args{[]User{}, []User{}}, []User{}},
		{"arrays filed", args{a2, a1}, []User{
			{5, "1", 9},
			{6, "2", 9},
			{7, "3", 30},
			{8, "4", 30}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
