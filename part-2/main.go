package main

import "fmt"

type User struct {
	id   int
	name string
}

func main() {
	usuarios := []User{{name: "Nico", id: 1}, {name: "Juan", id: 2}, {name: "Pedro", id: 3}}
	mapper := make(map[User]int)
	//Fomra de declarar un map o construit
	for _, user := range usuarios {
		mapper[user] = user.id
	}
	fmt.Println("Usuarios:", mapper)
	fmt.Println("This is Class2")
	numeros := []int{7, 8, 9, 10, 11, 12, 13}
	masNumeros := []int{14, 15, 16, 17, 18, 19, 20}
	nombres := make([]string, 10)
	nombres[0] = "Nico"
	nombres[1] = "Juan"
	nombres[2] = "Pedro"
	fmt.Println("Antes de agregar:", numeros)
	agregar(numeros)
	fmt.Println("Despues de agregar:", numeros)
	fmt.Println("Antes de agregar nombres:", nombres)
	agregarNames(nombres)
	fmt.Println("Despues de agregar nombres:", nombres)
	numeros = append(numeros, masNumeros...)
	fmt.Println("Despues de agregar mas numeros:", numeros)
}

func agregar(numeros []int) {
	numeros[0] = 1
	numeros = append(numeros, 15)
}
func agregarNames(names []string) {
	names = append(names, "Roger")
}
