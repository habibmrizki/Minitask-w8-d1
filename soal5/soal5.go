package soal5

import "fmt"

func Main() {

	originalSlice := []int{50, 75, 66, 20, 32, 90}
	valueToInsert := 88

	mid := len(originalSlice) / 2
	originalSlice = append(originalSlice[:mid], append([]int{valueToInsert}, originalSlice[mid:]...)...)

	fmt.Println(originalSlice)

	// originalSlice := []int{50, 75, 66, 20, 32, 90}
	// index := 3
	// valueToInsert := 88
	// newSlice := append(originalSlice[:index], valueToInsert)
	// newSlice = append(newSlice, originalSlice[index:]...)
	// fmt.Println(newSlice)

	// 	originalSlice := []int{50, 75, 66, 20, 32, 90}
	// index := 3
	// valueToInsert := 88
	// // newSlice := append(originalSlice[:index], valueToInsert)
	// // newSlice = append(newSlice, originalSlice[index:]...)

	// firstPart := originalSlice[:index]
	// secondPart := originalSlice[index:]

	// originalSlice = append(firstPart, valueToInsert)
	// originalSlice = append(originalSlice, secondPart...)

	// fmt.Println(originalSlice)

}
