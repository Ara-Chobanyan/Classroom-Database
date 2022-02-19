package menu

import (
	"bytes"
	"testing"
)

func TestFindAverage(t *testing.T) {
	testCases := []struct {
		input []int
		want  float64
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

func TestGiveAString(t *testing.T) {
	testCases := []struct {
		input *bytes.Buffer
		want  string
	}{
		{bytes.NewBufferString("Jon"), "Jon"},
		{bytes.NewBufferString("Ara"), "Ara"},
		{bytes.NewBufferString("Steve"), "Steve"},
		{bytes.NewBufferString("Mark"), "Mark"},
		{bytes.NewBufferString("Leon"), "Leon"},
		{bytes.NewBufferString("Aramo"), "Aramo"},
		{bytes.NewBufferString("Lucy"), "Lucy"},
		{bytes.NewBufferString("Ani"), "Ani"},
		{bytes.NewBufferString("Lia"), "Lia"},
		{bytes.NewBufferString("Lyia"), "Lyia"},
		{bytes.NewBufferString("Chrissy"), "Chrissy"},
	}
	for _, tc := range testCases {
		t.Run("TestGiveAString", func(t *testing.T) {
			actual := GiveAString("test", tc.input)
			if actual != tc.want {
				t.Errorf("got %s; want %s", tc.input, actual)
			}
		})
	}

}

func TestCreateAStudentGrade(t *testing.T) {
	testCases := []struct {
		input *bytes.Buffer
		want  float64
	}{
		{bytes.NewBufferString("88 99 100"), 96},
		{bytes.NewBufferString("77 55 45"), 59},
		{bytes.NewBufferString("44 66 77 77 88"), 70},
		{bytes.NewBufferString("77 65 88 89 100"), 84},
		{bytes.NewBufferString("33 55 88 0"), 44},
	}
	for _, tc := range testCases {
		t.Run("TestCreateAStudentGrade", func(t *testing.T) {
			actual := CreateAStudentGrade(tc.input)
			if actual != tc.want {
				t.Errorf("got %s; want %f", tc.input, actual)
			}
		})
	}
}

func TestGiveAInt(t *testing.T) {
	testCases := []struct {
		input *bytes.Buffer
		want  int
	}{
		{bytes.NewBufferString("96"), 96},
		{bytes.NewBufferString("59"), 59},
		{bytes.NewBufferString("1"), 1},
		{bytes.NewBufferString("2"), 2},
		{bytes.NewBufferString("13"), 13},
		{bytes.NewBufferString("20"), 20},
	}
	for _, tc := range testCases {
		t.Run("TestGiveAInt", func(t *testing.T) {
			actual := GiveAInt("test", tc.input)
			if actual != tc.want {
				t.Errorf("got %s; want %d", tc.input, actual)
			}
		})
	}

}
