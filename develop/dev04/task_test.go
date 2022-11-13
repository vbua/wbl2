package main

import (
	"reflect"
	"testing"
)

func TestCheckIfAnagram(t *testing.T) {
	type test struct {
		input []string
		want  bool
	}
	tests := []test{
		{[]string{"тяпка", "пятка"}, true},
		{[]string{"тяпка", "пяток"}, false},
	}

	for _, tc := range tests {
		got := checkIfAnagram(tc.input[0], tc.input[1])
		if got != tc.want {
			t.Errorf("got %v want %v given: %v, %v", got, tc.want, tc.input[0], tc.input[1])
		}
	}
}

func TestFormSetOfAnagrams(t *testing.T) {
	type test struct {
		input []string
		want  map[string][]string
	}
	tests := []test{
		{
			[]string{"пятак", "лИсток", "слиток", "слиток", "листок", "пятка", "столик", "тяпка"},
			map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			[]string{"тяпка", "пятак", "лИсток", "слиток", "слиток", "листок", "пятка", "столик"},
			map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"тяпка":  {"тяпка", "пятак", "пятка"},
			},
		},
		{
			[]string{"слиток", "что-то"},
			map[string][]string{},
		},
	}

	for _, tc := range tests {
		got := formSetOfAnagrams(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v want %v given: %v", got, tc.want, tc.input)
		}
	}
}
