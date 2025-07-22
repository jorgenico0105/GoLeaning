package main

import (
	"fmt"
	"strconv"
)

type Transporte interface {
	sonido() string
	consumo() string
}
type Carro struct {
	modelo string
	motor  uint
}
type Moto struct {
	modelo string
	motor  uint
}
type Barco struct {
	modelo string
	motor  uint
}

func (c Carro) sonido() string {
	return "Vroom Vroom"
}
func (c Carro) consumo() string {
	return "10 L/100km"
}
func (m Moto) sonido() string {
	return "Brrr Brrr"
}
func (m Moto) consumo() string {
	return "5 L/100km"
}
func (b Barco) sonido() string {
	return "Splash Splash"
}
func calcularConsumo(t Transporte) {
	fmt.Println(t.consumo())
}
func imprimirSonido(t Transporte) {
	fmt.Println(t.sonido())
}

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
	var carro Transporte = Carro{modelo: "Toyota", motor: 2000}
	var moto Transporte = Moto{modelo: "Yamaha", motor: 150}
	imprimirSonido(carro)
	imprimirSonido(moto)
	fmt.Println("Edad:", nicolas.Age)
	nicolas.Age = 31
	fmt.Println("Nueva Edad:", nicolas.Age)
	sumAge(&nicolas)
	fmt.Println("Edad despues de la funcion:", nicolas.Age)
	nombre := "Nicolas"
	nombre2 := "Gonzalez"
	_, apellido := returnNames(nombre, nombre2)
	fmt.Println("El apellido es:", apellido)
	i, err := strconv.Atoi("145")
	if err != nil {
		fmt.Println("Error converting string to int:", err.Error())
		return
	}
	fmt.Println("Converted integer:", i)
	us, err := getUserById(3)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	// while en go
	//cont := 0
	//for cont < 5{
	//  fmt.Println("Contador:", cont)
	// cont++
	//}
	fmt.Println("User found:", us)
}
func returnNames(nombre1, nombre2 string) (string, string) {
	return nombre1, nombre2
}
func sumAge(persona *Persona) {
	persona.Age += 20
}

type UserId struct {
	id   uint
	name string
}

func getUserById(id uint) (string, error) {
	users := []UserId{{id: 1, name: "Nicolas"}, {id: 2, name: "Javier"}}
	for _, user := range users {
		if user.id == id {
			return user.name, nil
		}

	}
	return "", fmt.Errorf("User with id %d not found", id)
}
