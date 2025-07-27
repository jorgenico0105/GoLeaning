package main

import "fmt"

type User struct {
	id    int
	name  string
	rol   string
	roles []string
}

func main() {
	usuarios := []User{{name: "Nico", id: 1, rol: "admin"}, {name: "Juan", id: 2, rol: "mecanico"}, {name: "Pedro", id: 3}, {name: "Nico", id: 1, rol: "facturador"}}
	mapper := make(map[int]User)
	//Fomra de declarar un map o construit
	fmt.Println("Usuarios:", usuarios)
	for _, user := range usuarios {
		if _, existe := mapper[user.id]; !existe {
			mapper[user.id] = user
		}
		u := mapper[user.id]
		u.roles = append(u.roles, user.rol)
		mapper[user.id] = u
	}
	fmt.Println("Usuarios:", mapper)
	for k, v := range mapper {
		fmt.Println("ID:", k, "Usuario:", v)
	}

	fmt.Println("Usuarios con roles:")
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
	_, number2 := testFunction(5, 8)
	fmt.Println("Number2:", number2)
}

func agregar(numeros []int) {
	numeros[0] = 1
	numeros = append(numeros, 15)
}
func agregarNames(names []string) {
	names = append(names, "Roger")
}
func testFunction(number, number2 int) (sum, sum2 int) {
	sum = number + number2
	sum2 = (number * 3) + number2
	return number, number2
}
