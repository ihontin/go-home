package main

import "testing"

type checkMap struct {
	received []string
	expected map[string]int
}

func TestCountWordsInText(t *testing.T) {
	var expected1 = map[string]int{"ipsum": 2, "sit": 3}
	var expected2 = map[string]int{".": 0}
	var expected3 = map[string]int{}
	newCheck := []checkMap{
		{[]string{"ipsum", "sit"}, expected1},
		{[]string{"."}, expected2},
		{[]string{}, expected3},
	}

	newTxt := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris.
       Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor.
       Praesent et diam eget libero egestas mattis sit amet vitae augue.`
	for _, s := range newCheck {
		got := CountWordsInText(newTxt, s.received)
		for key, val := range got {
			if _, ok := s.expected[key]; !ok || s.expected[key] != val {
				t.Errorf("expected key = %v, got key = %v", s.expected, got)
			}
		}

	}
}
