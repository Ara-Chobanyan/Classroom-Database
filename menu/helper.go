package menu

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

func GiveAInt(a string, r io.Reader) int {
	fmt.Println(a)
	var i int
	_, err := fmt.Fscanf(r, "%d", &i)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func GiveAString(a string, r io.Reader) string {
	fmt.Println(a)
	var n string
	_, err := fmt.Fscanf(r, "%s", &n)
	if err != nil {
		fmt.Println(err)
	}
	return n
}

func CreateAStudentGrade(r io.Reader) float64 {
	var grades []int
	fmt.Println("Please Enter grades with space in between")

	scanner := bufio.NewScanner(r)

	scanner.Scan()
	grades = ConvertStringToSlice(scanner.Text())

	a := FindAverage(grades)
	return math.RoundToEven(a)
}

func ConvertStringToSlice(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func FindAverage(grades []int) float64 {
	length := len(grades)
	sum := 0

	for i := 0; i < length; i++ {
		sum += grades[i]
	}
	return (float64(sum)) / (float64(length))
}
