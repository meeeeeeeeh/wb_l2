package main

import (
	"testing"
)

func TestGetAnagrams(t *testing.T) {
	str := "пятак ПЯТКА тяпка слиток лиСТок листок Столик акрОбат работка бочка"

	expected := make(map[string][]string)
	expected["акробат"] = []string{"акробат", "работка"}
	expected["пятак"] = []string{"пятак", "пятка", "тяпка"}
	expected["слиток"] = []string{"листок", "слиток", "столик"}

	res := getAnagrams(str)

	for key, val := range res {
		item, ok := expected[key] 
		if !ok {
			t.Errorf("expected value: '%v' wasn't found in the result map", val)
		}

		

		if len(res[key]) != len(expected[key]) {
			t.Errorf("result value: '%v' doesn't match expected value ", val)
		}

		for i := range item {
			if res[key][i] != expected[key][i] {
				t.Errorf("result value: '%v' doesn't match expected value ", val)
			}
		}
	}
}