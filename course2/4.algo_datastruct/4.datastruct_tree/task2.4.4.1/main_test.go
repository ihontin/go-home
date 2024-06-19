package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestBinaryTree_insert(t1 *testing.T) {
	type fields struct {
		root *Node
	}
	type args struct {
		user *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *BinaryTree
	}{
		{
			name: "ok",
			fields: fields{
				root: nil,
			},
			args: args{
				user: &User{18, "Galing", 13},
			},
			want: &BinaryTree{
				&Node{18, &User{18, "Galing", 13}, nil, nil},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &BinaryTree{
				root: tt.fields.root,
			}
			if got := t.insert(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryTree_search(t1 *testing.T) {
	type fields struct {
		root *Node
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *User
	}{
		{
			name:   "ok",
			fields: fields{&Node{18, &User{18, "Galing", 13}, nil, nil}},
			args:   args{18},
			want:   &User{18, "Galing", 13},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &BinaryTree{
				root: tt.fields.root,
			}
			if got := t.search(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_insert(t *testing.T) {
	type fields struct {
		index int
		data  *User
		left  *Node
		right *Node
	}
	type args struct {
		user *User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "ok",
			fields: fields{},
			args: args{
				user: &User{18, "Galing", 13},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				index: tt.fields.index,
				data:  tt.fields.data,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			n.insert(tt.args.user)
		})
	}
}

func TestNode_search(t *testing.T) {
	type fields struct {
		index int
		data  *User
		left  *Node
		right *Node
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *User
	}{
		{
			name:   "ok",
			fields: fields{18, &User{18, "Galing", 13}, nil, nil},
			args:   args{18},
			want:   &User{18, "Galing", 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				index: tt.fields.index,
				data:  tt.fields.data,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := n.search(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateData(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name      string
		args      args
		searshBin *BinaryTree
	}{
		{
			name: "ok",
			args: args{13},
			searshBin: &BinaryTree{
				&Node{18, &User{18, "Galing", 13}, nil, nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateData(tt.args.n)
			got.insert(&User{200, "Galing", 13})
			expected := tt.searshBin.search(tt.args.n)
			if !reflect.DeepEqual(got.search(tt.args.n), expected) {
				t.Errorf("generateData() = %v, want %v", got, expected)
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

	expected := "30\n"

	stdOut := bytes.Buffer{}
	_, _ = stdOut.ReadFrom(r)
	if expected != stdOut.String() {
		t.Errorf("expected = %s, got = %s", expected, stdOut.String())
	}
}
