package main

import "fmt"

func main() {
	fmt.Println("Hello, Advanced Go!")
	suma := sumar(5, 3)
	fmt.Println("The sum is:", suma)
	fullPhrase := concatter()
	fullPhrase("Hello")
	fullPhrase("World")

	fmt.Println("Concatenated phrase:")
	phrase, count := fullPhrase("from Go!")
	fmt.Println(phrase)
	fmt.Println("Word count:", count)
}

func sumar(a, b int) int {
	defer restar(a, b) // This will execute after sumar
	return a + b
}
func restar(a, b int) int {
	fmt.Println("Resting...")
	return a - b
}
func concatter() func(string) (string, int) {
	doc := ""
	count := 0
	return func(word string) (string, int) {
		doc += word + " "
		count++
		return doc, count
	}
}
