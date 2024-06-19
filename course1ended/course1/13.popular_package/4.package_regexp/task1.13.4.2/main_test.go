package main

import (
	"reflect"
	"testing"
)

func TestCensorAds(t *testing.T) {
	ads := []Ad{
		{
			Title:    "Куплю велосипед MeRiDa",
			Описание: "Куплю велосипед meriDA в хорошем состоянии.",
		},
		{
			Title:    "Продам ВаЗ 2101",
			Описание: "Продам ваз 2101 в хорошем состоянии.",
		},
		{
			Title:    "Продам БМВ",
			Описание: "Продам бМв в хорошем состоянии.",
		},
		{
			Title:    "Продам macBook pro",
			Описание: "Продам macBook PRO в хорошем состоянии.",
		},
	}

	expected := []Ad{
		{
			Title:    "Куплю телефон Apple",
			Описание: "Куплю телефон Apple в хорошем состоянии.",
		},
		{
			Title:    "Продам ВАЗ 2101",
			Описание: "Продам ВАЗ 2101 в хорошем состоянии.",
		},
		{
			Title:    "Продам BMW",
			Описание: "Продам BMW в хорошем состоянии.",
		},
		{
			Title:    "Продам Macbook Pro",
			Описание: "Продам Macbook Pro в хорошем состоянии.",
		},
	}
	var ch = map[string]string{
		"велосипед merida": "телефон Apple",
		"ваз":              "ВАЗ",
		"бмв":              "BMW",
		"macbook pro":      "Macbook Pro",
	}
	if got := censorAds(ads, ch); !reflect.DeepEqual(got, expected) {
		t.Errorf("expected = %v, got = %v", expected, got)
	}

}
