package main

import (
	"testing"
)

// func TestSumOfInterval(t *testing.T) {
// 	a := [][2]int{{1, 2}, {6, 10}, {11, 15}} //9
// 	b := [][2]int{{1, 5}, {10, 20}, {15, 18}}
// 	c := [][2]int{{1, 4}, {7, 10}, {3, 5}}
// 	result1 := SumOfIntervals(a)
// 	if result1 != 9 {
// 		t.Errorf("for \"a\" scenario expected 9, got %d", result1)
// 	}
// 	t.Log("test1")
// 	result2 := SumOfIntervals(b)
// 	if result2 != 14 {
// 		t.Errorf("for \"b\" scenario expected 14, got %d", result2)
// 	}
// 	t.Log("test2")

// 	result3 := SumOfIntervals(c)
// 	if result3 != 7 {
// 		t.Errorf("for \"c\" scenario expected 7, got %d", result3)
// 	}
// 	t.Log("test3")

// }

func TestSumOfInterval(t *testing.T) {

	var testCases = []struct {
		text     string
		input    [][2]int
		expected int
	}{
		{"First scenario", [][2]int{{1, 2}, {6, 10}, {11, 15}}, 9},
		{"Second scenario", [][2]int{{1, 5}, {10, 20}, {15, 18}}, 14},
		{"Third scenario", [][2]int{{1, 4}, {7, 10}, {3, 5}}, 7},
	}

	for _, tc := range testCases {
		t.Run(tc.text, func(t *testing.T) {
			result := SumOfIntervals(tc.input)
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}

}
