package main

import "fmt"

type comida interface {
	tiempoPreparacion() string
}
type pizza struct {
	peso   int
	precio float64
}

func (p pizza) tiempoPreparacion() string {
	return "Tiempo de preparacion de la pizza: 30 minutos"
}
func imprimirTiempo(c comida) {
	fmt.Println(c.tiempoPreparacion())
}

type Persona struct {
	name string
	age  int
}

func main() {
	Nicolas := Persona{name: "Nicolas", age: 30}
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
	fmt.Println("Persona:", Nicolas)
	//AddFiveYears(&Nicolas)
	Nicolas.AddFiveYears()
	fmt.Println("Name after sum ages: ", Nicolas)
	x := 2
	y := &x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of y (address of x):", y)
	fmt.Println("Value at address y (value of x):", *y)
}

func (p *Persona) AddFiveYears() {
	p.age += 5
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
