package soal3

import "fmt"

func Main() {
	numbers := [5]int8{30, 64, 19, 88, 21}
	newNumbers := numbers[0:]
	fmt.Println(newNumbers)
	fmt.Println("nilai terbesar adalah", bigNumber(numbers))
}

func bigNumber(numbers [5]int8) int8 {
	result := numbers[0]
	for _, number := range numbers {
		if number > result {
			result = number
		}
	}
	return result
}
