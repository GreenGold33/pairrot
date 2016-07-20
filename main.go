package main

import "fmt"

func combinations(array []string, numPerGroup int) [][]string {
	var pairs [][]string
	// For each element
	//   iterate over the remaining elements
	//   and add them to an array and add them to pairs
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			pairs = append(pairs, []string{array[i], array[j]})
		}
	}
	return pairs
}

func main() {
	students := []string{
		"Federico",
		"John",
		"Rachel",
		"Christie",
		"Weili",
		"Andres",
		"Carolina",
		"Jonathan",
		"Alberto",
		"Alison",
		"Devorah",
		"Jonathan",
		"Minyoung",
		"kkkjjjkkDaniel",
		"John",
		"Sebastian",
		"Joseph",
		"Gabriel",
	}
	pairs := combinations(students, 2)
	fmt.Print(pairs)
}

//func printPairs(pairs [][]string) {
//	fmt.Println("Pairs!")
//	for pair := range pairs {
//		fmt.Print("Pair: ")
//		for student := range pair {
//
//		}
//	}
//}
