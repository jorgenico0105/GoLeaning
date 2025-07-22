package main

import "fmt"

type Persona struct {
	Age       uint
	Height    float64
	Weight    float64
	SkinColor string
	Familia   struct {
		numeroHijos int
		padreNombre string
	}
	mascotas []Animal
}
type Animal struct {
	Especies string
	Sonido   string
}
type Rect struct {
	h uint
	w uint
}
type User struct {
	username string
	password string
}

func (user User) getUsername() (string, string) {
	return user.username, user.password
}
func (r Rect) area() uint {
	return r.h * r.w
}
func main() {
	rec := Rect{h: 10, w: 20}
	fmt.Println(rec.area())
	fmt.Println("Hello, World!")
	//javier := Persona{}
	nicolas := Persona{
		Age:       30,
		Height:    1.75,
		Weight:    70.5,
		SkinColor: "Moreno",
		mascotas:  []Animal{{Especies: "Perro", Sonido: "Guau"}, {Especies: "Gato", Sonido: "Miau"}},
	}
	user := User{username: "nicolas", password: "1234"}
	user2 := User{username: "javier", password: "5678"}
	usuarios := []User{user, user2}
	for i, u := range usuarios {
		password, username := u.getUsername()
		fmt.Println("Usuario:", i, "Username:", username, "Password:", password)
	}
	fmt.Println("Edad:", nicolas.Age)
	nicolas.Age = 31
	fmt.Println("Nueva Edad:", nicolas.Age)
	sumAge(&nicolas)
	fmt.Println("Edad despues de la funcion:", nicolas.Age)
	nombre := "Nicolas"
	nombre2 := "Gonzalez"
	_, apellido := returnNames(nombre, nombre2)
	fmt.Println("El apellido es:", apellido)
}
func returnNames(nombre1, nombre2 string) (string, string) {
	return nombre1, nombre2
}
func sumAge(persona *Persona) {
	persona.Age += 20
}
