package main

import (
	"API-E-commerce/server"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	go server.Server()
	time.Sleep(500 * time.Millisecond)
	terminalInterface()
}

func terminalInterface() {
	fmt.Print("Bienvenido \n")
	fmt.Println("Opciones:")
	fmt.Println("1. Iniciar sesión como cliente")
	fmt.Println("2. Iniciar sesión como administrador")
	fmt.Println("3. Salir")

	var firstOption string
	fmt.Print("Ingrese una opción: ")
	fmt.Scan(&firstOption)

	if firstOption == "1" {
		var userId string
		var userPass string
		fmt.Print("Ingrese su id: ")
		fmt.Scan(&userId)
		fmt.Print("Ingrese su contraseña: ")
		fmt.Scan(&userPass)
		//llamar a la api y ver si el usuario existe
	} else if firstOption == "2" {
		var adminPass string
		fmt.Println("Ingrese contraseña de administrador: ")
		fmt.Scan(&adminPass)
	} else if firstOption == "3" {
		fmt.Println("eligio la opcion 3")
	}
}
