package main

import (
	"testing"
)

func TestCountWordOccurrences(t *testing.T) {
	checkList := []string{"rammstein", "rammstein", "rammstein", "rammstein", "rammstein"}
	var checkWord string
	var checkMap = make(map[string]int)
	for i, _ := range checkList {
		checkWord += " " + checkList[i]
		checkMap = countWordOccurrences(checkWord)
		expected := i + 1
		if expected != checkMap[checkList[i]] {
			t.Errorf("expected = %d, got = %d", expected, checkMap[checkList[i]])
		}
	}
}

//func TestMainFunc(t *testing.T) {
//	old := os.Stdout
//	r, w, _ := os.Pipe()
//	os.Stdout = w
//
//	main()
//	w.Close()
//	os.Stdout = old
//	expected := "amet: 1\nconsectetur: 1\nadipiscing: 1\nelit: 1\nLorem: 1\nipsum: 2\ndolor: 1\nsit: 1\n"
//
//	var stdOsOut = bytes.Buffer{}
//	stdOsOut.ReadFrom(r)
//	if expected != stdOsOut.String() {
//		t.Errorf("expected = %s, got = %s", expected, stdOsOut.String())
//	}
//}
