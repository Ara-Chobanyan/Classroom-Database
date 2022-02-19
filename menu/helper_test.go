package menu

import "testing"

func TestFindAverage(t *testing.T) {
	testCases := []struct {
		input []int
		want  float32
	}{
		{[]int{33, 44, 55, 66}, 49.5},
		{[]int{50, 100, 87, 66}, 75.75},
		{[]int{100, 100, 100, 100}, 100.0},
		{[]int{98, 77, 33, 78}, 71.5},
		{[]int{89, 51, 45, 100}, 71.25},
	}

	for _, tc := range testCases {
		t.Run("check average", func(t *testing.T) {
			actual := FindAverage(tc.input)
			if actual != tc.want {
				t.Errorf("got %f; want %f", actual, tc.want)
			}
		})
	}

}

func TestConvertStringToSlice(t *testing.T) {
	testCases := []struct {
		input  string
		output []int
	}{
		{"55", []int{55}},
		{"100", []int{100}},
		{"77", []int{77}},
		{"88", []int{88}},
		{"11", []int{11}},
		{"45", []int{45}},
		{"89", []int{89}},
		{"0", []int{0}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := ConvertStringToSlice(tc.input)
			if actual[0] != tc.output[0] {
				t.Errorf("got %d; want %d", tc.output, actual)
			}
		})
	}
}
