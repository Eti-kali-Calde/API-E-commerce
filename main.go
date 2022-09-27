package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Opciones:")
	fmt.Println("1. Iniciar sesión como cliente")
	fmt.Println("2. Iniciar sesión como administrador")
	fmt.Println("3. Salir")

	var firstOption string
	fmt.Print("Ingrese una opción: ")
	fmt.Scan(&firstOption)

	if firstOption == "1" {
		var secondOption string
		fmt.Println("Opciones:")
		fmt.Println("1. Ver lista de productos")
		fmt.Println("2. Hacer compra")
		fmt.Println("3. Salir")
		fmt.Print("Ingrese una opción: ")
		fmt.Scan(&secondOption)
		if secondOption == "1" {
			response, err := http.Get("http://localhost:5000/api/productos")

			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			responseData, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(responseData))
		}
	} else if firstOption == "2" {
		var adminPass string
		fmt.Println("Ingrese contraseña de administrador: ")
		fmt.Scan(&adminPass)
		if adminPass == "1234" {
			var secondOption string
			fmt.Println("Opciones:")
			fmt.Println("1. Ver lista de productos")
			fmt.Println("2. Hacer compra")
			fmt.Println("3. Eliminar productos")
			fmt.Println("4. Ver estadisticas")
			fmt.Println("5. Salir")
			fmt.Print("Ingrese una opción: ")
			fmt.Scan(&secondOption)
			if secondOption == "1" {
				response, err := http.Get("http://localhost:5000/api/productos")

				if err != nil {
					fmt.Print(err.Error())
					os.Exit(1)
				}

				responseData, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(string(responseData))
			} else if secondOption == "2" {

			} else if secondOption == "3" {
				var id string
				fmt.Println("ingrese id del producto a eliminar")
				fmt.Scan(&id)

			}
		}
	} else if firstOption == "3" {
		fmt.Println("eligio la opcion 3")
	}
}
