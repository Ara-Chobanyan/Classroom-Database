package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GiveAInt(a string) int {
	fmt.Println(a)
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func GiveAString(a string) string {
	fmt.Println(a)
	var i string
	_, err := fmt.Scanf("%s", &i)
	if err != nil {
		fmt.Println(err)
	}
	return i
}

func CreateAStudentGrade() float32 {
	var grades []int
	fmt.Println("Please Enter grades with space in between")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	grades = ConvertStringToSlice(scanner.Text())

	a := FindAverage(grades)
	return a
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

func FindAverage(grades []int) float32 {
	length := len(grades)
	sum := 0

	for i := 0; i < length; i++ {
		sum += grades[i]
	}
	return (float32(sum)) / (float32(length))
}
