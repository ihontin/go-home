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

	expected := "Original:  [64 34 25 12 22 11 90]\nSorted by Merge Sort:  [11 12 22 25 34 64 90]\n" +
		"Sorted by Insertion Sort:  [11 12 22 25 34 64 90]\nSorted by Selection Sort:  [11 12 22 25 34 64 90]\n" +
		"Sorted by Quicksort:  [11 12 22 25 34 64 90]\nSorted by GeneralSort:  [11 12 22 25 34 64 90]\n"

	stdOut := bytes.Buffer{}
	_, _ = stdOut.ReadFrom(r)
	if expected != stdOut.String() {
		t.Errorf("expected = %s, got = %s", expected, stdOut.String())
	}
}

func TestGeneralSort(t *testing.T) {
	var bigList = make([]int, 1200)
	var bigListsorted = make([]int, 1200)
	for i := 1200 - 1; i >= 0; i-- {
		bigList[i] = (1199 - i) * 10
		bigListsorted[i] = i * 10
	}
	type args struct {
		arr []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{"ok", args{[]int{99, 6, 6, -1}}, []int{-1, 6, 6, 99}},
		{"!ok", args{[]int{}}, []int{}},
		{"o2k", args{bigList}, bigListsorted},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GeneralSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.expected) {
				t.Errorf("expected = %v, got =  %v", tt.expected, tt.args.arr)
			}
		})
	}
}

func Test_insertionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{"ok", args{[]int{99, 6, 6, -1}}, []int{-1, 6, 6, 99}},
		{"!ok", args{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertionSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.expected) {
				t.Errorf("expected = %v, got =  %v", tt.expected, tt.args.arr)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		left  []int
		right []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"ok", args{[]int{6, 99}, []int{-1, 6}}, []int{-1, 6, 6, 99}},
		{"!ok", args{[]int{}, []int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"ok", args{[]int{99, 6, 6, -1}}, []int{-1, 6, 6, 99}},
		{"!ok", args{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ok", args{[]int{99, 6, 6, -1}, 2, 3}, 2},
		{"!ok", args{[]int{}, 0, 0}, 0},
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
		arr  []int
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"ok", args{[]int{99, 6, 6, -1}, 0, 3}, []int{-1, 6, 6, 99}},
		{"!ok", args{[]int{}, 0, 0}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quicksort(tt.args.arr, tt.args.low, tt.args.high)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("expected = %v, got =  %v", tt.want, tt.args.arr)
			}
		})
	}
}

func Test_selectionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"ok", args{[]int{99, 6, 6, -1}}, []int{-1, 6, 6, 99}},
		{"!ok", args{[]int{}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectionSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("expected = %v, got =  %v", tt.want, tt.args.arr)
			}
		})
	}
}
