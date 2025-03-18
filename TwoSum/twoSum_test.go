package main

import "testing"

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		text     string
		inpuList []int
		inputInt int
		output   []int
	}{
		{"first scenario", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"second scenario", []int{3, 2, 4}, 6, []int{1, 2}},
		{"third scenario", []int{3, 2, 3}, 6, []int{0, 2}},
	}

	for _, tt := range testCases {
		t.Run(tt.text, func(t *testing.T) {
			result := twoSum(tt.inpuList, tt.inputInt)
			if !compare(result, tt.output) {
				t.Errorf("have: %v, want: %v", result, tt.output)
			}
		})
	}
}

func compare(a, b []int) bool {
	result := false
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			result = true
		} else {
			result = false
		}
	}
	return result
}
