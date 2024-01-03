package main

import "fmt"

type Alumno struct {
	Name     string
	Apellido string
	DNI      string
	Fecha    string
}

func (a *Alumno) detalle() {
	fmt.Printf("Name: %s \nApellido: %s \nDNI: %s \nFecha: %s \n", a.Name, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	alumno := Alumno{
		Name:     "John",
		Apellido: "Doe",
		DNI:      "12345678",
		Fecha:    "01/01/2000",
	}

	alumno.detalle()
}
