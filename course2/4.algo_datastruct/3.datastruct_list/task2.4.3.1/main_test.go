package main

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestDoubleLinkedList_Current(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		{
			name:   "ok",
			fields: fields{curr: &Node{}},
			want:   &Node{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Current(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Current() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Delete(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{curr: &Node{}},
			args:    args{23},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			newCom := GenerateData()
			for _, com := range newCom {
				_ = d.Push(com)
			}
			if err := d.Delete(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_DeleteCurrent(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{curr: &Node{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			newCom := GenerateData()
			for _, com := range newCom {
				_ = d.Push(com)
			}
			_ = d.SetCurrent(11)
			if err := d.DeleteCurrent(); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCurrent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_GetByIndex(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{curr: &Node{}},
			args:    args{30},
			wantErr: false,
		},
		{
			name:    "!ok",
			fields:  fields{curr: &Node{}},
			args:    args{31},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			newCom := GenerateData()
			for _, com := range newCom {
				_ = d.Push(com)
			}
			commits := Commit{}
			commits = Commit{
				Message: "111",
				UUID:    "222",
			}
			_ = d.Push(commits)
			got, err := d.GetByIndex(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.name == "!ok" {
				return
			}
			if !reflect.DeepEqual(*got.data, commits) {
				t.Errorf("GetByIndex() got = %v, want %v", got, commits)
			}
		})
	}
}

func TestDoubleLinkedList_Index(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{curr: &Node{}},
			want:    23,
			wantErr: false,
		},
		{
			name:    "!ok",
			fields:  fields{curr: &Node{}},
			want:    33,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			newCom := GenerateData()
			for _, com := range newCom {
				_ = d.Push(com)
			}
			_ = d.SetCurrent(tt.want)
			got, err := d.Index()
			if (err != nil) != tt.wantErr {
				t.Errorf("Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.name == "!ok" {
				return
			}
			if got != tt.want {
				t.Errorf("Index() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Init(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		c []Commit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "ok",
			fields: fields{curr: &Node{}},
			args:   args{make([]Commit, 0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			d.Init(tt.args.c)
		})
	}
}

func TestDoubleLinkedList_Insert(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		n int
		c Commit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{curr: &Node{}},
			args:    args{1, Commit{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			newCom := GenerateData()
			for _, com := range newCom {
				_ = d.Push(com)
			}
			_ = d.SetCurrent(11)
			if err := d.Insert(tt.args.n, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Len(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "ok",
			fields: fields{curr: &Node{}},
			want:   30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			newCom := GenerateData()
			for _, com := range newCom {
				_ = d.Push(com)
			}

			if got := d.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_LoadData(t *testing.T) {
	_, exe, _, _ := runtime.Caller(0)
	dir := filepath.Dir(exe) + "/commits.json"
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{curr: &Node{}},
			args:    args{dir},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.LoadData(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("LoadData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Next(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		{
			name:   "ok",
			fields: fields{curr: &Node{}},
			want:   &Node{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			newCom := GenerateData()
			for _, com := range newCom {
				_ = d.Push(com)
			}
			commits := Commit{
				Message: "111",
				UUID:    "222",
			}
			_ = d.Push(commits)
			_ = d.SetCurrent(31)
			if got := d.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Pop(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		{
			name:   "ok",
			fields: fields{curr: &Node{}},
			want: &Node{data: &Commit{
				Message: "111",
				UUID:    "222",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			commits := Commit{
				Message: "111",
				UUID:    "222",
			}
			_ = d.Push(commits)
			if got := d.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Prev(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		{
			name:   "ok",
			fields: fields{curr: &Node{}},
			want:   &Node{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			newCom := GenerateData()
			for _, com := range newCom {
				_ = d.Push(com)
			}
			commits := Commit{
				Message: "111",
				UUID:    "222",
			}
			_ = d.Push(commits)
			_ = d.SetCurrent(33)
			if got := d.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Push(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		c Commit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{curr: &Node{}},
			args:    args{Commit{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.Push(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Reverse(t *testing.T) {
	commits2 := Commit{
		Message: "33",
		UUID:    "44",
	}
	commits := Commit{
		Message: "111",
		UUID:    "222",
	}
	forTest := &DoubleLinkedList{}
	_ = forTest.Push(commits2)
	_ = forTest.Push(commits)

	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *DoubleLinkedList
	}{
		{
			name:   "ok",
			fields: fields{curr: &Node{}},
			want:   forTest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}

			_ = d.Push(commits)
			_ = d.Push(commits2)

			if got := d.Reverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Search(t *testing.T) {
	commits2 := Commit{
		Message: "33",
		UUID:    "44",
	}
	commits := Commit{
		Message: "test",
		UUID:    "222",
	}
	forTest := &DoubleLinkedList{}
	_ = forTest.Push(commits)
	_ = forTest.Push(commits2)
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			name:   "ok",
			fields: fields{curr: &Node{}},
			args:   args{"test"},
			want:   forTest.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			_ = d.Push(commits)
			_ = d.Push(commits2)
			if got := d.Search(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_SearchUUID(t *testing.T) {
	commits2 := Commit{
		Message: "33",
		UUID:    "44",
	}
	commits := Commit{
		Message: "test",
		UUID:    "test",
	}
	forTest := &DoubleLinkedList{}
	_ = forTest.Push(commits)
	_ = forTest.Push(commits2)
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		uuID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			name:   "ok",
			fields: fields{curr: &Node{}},
			args:   args{"test"},
			want:   forTest.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			_ = d.Push(commits)
			_ = d.Push(commits2)
			if got := d.SearchUUID(tt.args.uuID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_SetCurrent(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "ok",
			fields:  fields{curr: &Node{}},
			args:    args{24},
			wantErr: false,
		},
		{
			name:    "ok",
			fields:  fields{curr: &Node{}},
			args:    args{33},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			newCom := GenerateData()
			for _, com := range newCom {
				_ = d.Push(com)
			}

			if err := d.SetCurrent(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("SetCurrent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Shift(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		{
			name:   "ok",
			fields: fields{curr: &Node{}},
			want: &Node{data: &Commit{
				Message: "111",
				UUID:    "222",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			commits := Commit{
				Message: "111",
				UUID:    "222",
			}
			_ = d.Push(commits)
			if got := d.Shift(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateData(t *testing.T) {
	tests := []struct {
		name string
		want []Commit
	}{
		{
			name: "ok",
			want: make([]Commit, 0, 30),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateData()
			got = got[:0]
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	ti := time.Now().Add(time.Hour * 1)
	ti1 := time.Now().Add(time.Hour * 4)
	ti2 := time.Now().Add(time.Hour * 2)
	ti3 := time.Now().Add(time.Hour * 3)
	type args struct {
		arr  []Commit
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ok", args{[]Commit{{Date: ti}, {Date: ti1}, {Date: ti2}, {Date: ti3}}, 2, 3}, 3},
		{"!ok", args{[]Commit{}, 0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.arr, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quicksort(t *testing.T) {
	type args struct {
		arr  []Commit
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
	}{
		{"ok", args{[]Commit{{UUID: "1"}, {UUID: "2"}, {UUID: "4"}, {UUID: "3"}}, 2, 3}},
		{"!ok", args{[]Commit{}, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quicksort(tt.args.arr, tt.args.low, tt.args.high)
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

	expected := "/home/i/work/src/studentgit.kata.academy/Alkolex/go-kata/course2/4.algo_datastruct/3.datastruct_list/task2.4.3.1/commits.json\n" +
		"We need to program the bluetooth ADP protocol!\n" +
		"Try to calculate the SSL protocol, maybe it will quantify the primary system!\n"

	stdOut := bytes.Buffer{}
	_, _ = stdOut.ReadFrom(r)
	if expected != stdOut.String() {
		t.Errorf("expected = %s, got = %s", expected, stdOut.String())
	}
}
