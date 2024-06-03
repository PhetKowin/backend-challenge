package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decode(code string) string {
	num := []int{}
	lowestNegativeValue := 0
	value := 0
	var builder strings.Builder
	for i := 0; i < len(code); i++ {
		if code[i] == 'L' {
			num = append(num, value)
			value = value - 1
		} else if code[i] == 'R' {
			num = append(num, value)
			value = value + 1
		} else if code[i] == '=' {
			num = append(num, value)
		}
		if value < 0 && value < lowestNegativeValue {
			lowestNegativeValue = value
		}
	}
	num = append(num, value)

	if lowestNegativeValue < 0 {
		for i := 0; i < len(num); i++ {
			num[i] = num[i] - lowestNegativeValue
			builder.WriteString(strconv.Itoa(num[i]))
		}
	} else {
		for i := 0; i < len(num); i++ {
			builder.WriteString(strconv.Itoa(num[i]))
		}
	}
	return builder.String()
}

func main() {
	var encoded string
	fmt.Scanf("%s", &encoded)
	output := decode(encoded)
	fmt.Println(output)
}
