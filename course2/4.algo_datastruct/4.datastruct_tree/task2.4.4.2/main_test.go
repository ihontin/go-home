package main

import (
	"bytes"
	"github.com/google/btree"
	"os"
	"reflect"
	"testing"
)

func TestBTree_Insert(t *testing.T) {
	type fields struct {
		tree *btree.BTree
	}
	type args struct {
		user User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"ok",
			fields{btree.New(5)},
			args{
				User{1, "", 1},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &BTree{
				tree: tt.fields.tree,
			}
			bt.Insert(tt.args.user)
		})
	}
}
func TestNewBTree(t *testing.T) {
	type args struct {
		degree int
	}
	tests := []struct {
		name string
		args args
		want *BTree
	}{
		{
			"ok",
			args{5},
			&BTree{btree.New(5)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBTree(tt.args.degree); !reflect.DeepEqual(got.tree.Len(), tt.want.tree.Len()) {
				t.Errorf("NewBTree() = %v, want %v", got.tree, tt.want.tree)
			}
		})
	}
}

func TestBTree_Search(t *testing.T) {
	type fields struct {
		tree *btree.BTree
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *User
	}{
		{
			"ok",
			fields{btree.New(5)},
			args{5},
			&User{ID: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &BTree{
				tree: tt.fields.tree,
			}
			bt.Insert(User{ID: 5})
			if got := bt.Search(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Less(t *testing.T) {
	us := User{16, "rrr", 7}
	a := &BTree{
		tree: btree.New(2),
	}
	a.Insert(us)
	btree.New(16)
	type fields struct {
		ID   int
		Name string
		Age  int
	}
	type args struct {
		than btree.Item
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"ok",
			fields{16, "rrr", 7},
			args{a.tree.Get(User{ID: 16})},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
				Age:  tt.fields.Age,
			}

			if got := u.Less(tt.args.than); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	_ = w.Close()
	os.Stdout = old

	expected := "Найден пользователь: {2 Bob 25}\n"

	stdOut := bytes.Buffer{}
	_, _ = stdOut.ReadFrom(r)
	if expected != stdOut.String() {
		t.Errorf("expected = %s, got = %s", expected, stdOut.String())
	}
}
