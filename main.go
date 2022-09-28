package main

import (
	"API-E-commerce/server"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Todo struct {
	Nombre              string `json:"nombre"`
	Cantidad_Disponible int    `json:"cantidad_disponible"`
	Precio_Unitario     int    `json:"precio_unitario"`
}
type Client struct {
	Id       int
	Password string
}

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

	if firstOption == "1" { //iniciar como cliente
		var userId int
		var userPass string
		fmt.Print("Ingrese su id: ")
		fmt.Scan(&userId)
		fmt.Print("Ingrese su contraseña: ")
		fmt.Scan(&userPass)
		todo := Client{userId, userPass}
		jsonReq, err := json.Marshal(todo)
		resp, err := http.Post("http://localhost:5000/api/clientes/iniciar_sesion", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
		if err != nil {
			fmt.Print("Error, no hay ninguna coincidencia con los datos ingresados.")
		}
		aux, _ := ioutil.ReadAll(resp.Body)
		userExist := string(aux)
		if userExist == "" {
			fmt.Print("Error, no hay ninguna coincidencia con los datos ingresados.")
		} else {
			fmt.Println("Inicio de sesión exitoso")
			fmt.Println("Opciones:")
			fmt.Println("1. Ver lista de productos")
			fmt.Println("2. Hacer compra")
			fmt.Println("3. Salir")
			var secondOption string
			fmt.Print("Ingrese una opción: ")
			fmt.Scan(&secondOption)
			if secondOption == "1" { //ver producto
				resp, err := http.Get("http://localhost:5000/api/productos")
				if err != nil {
					log.Fatalln(err)
				}

				defer resp.Body.Close()
				bodyBytes, _ := ioutil.ReadAll(resp.Body)

				// Convert response body to string
				bodyString := string(bodyBytes)
				fmt.Println("API Response as String:\n" + bodyString)
			} else if secondOption == "2" { //hacer compras

			} else if secondOption == "3" { //salir
				fmt.Print("Hasta luego!")
			}
		}

	} else if firstOption == "2" { //iniciar como admin
		var adminPass string
		fmt.Println("Ingrese contraseña de administrador: ")
		fmt.Scan(&adminPass)
		fmt.Println("Opciones:")
		fmt.Println("1. Ver lista de productos")
		fmt.Println("2. Crear producto")
		fmt.Println("3. Eliminar producto")
		fmt.Println("4. Ver estadísticas")
		fmt.Println("5. Salir")
		var secondOption string
		fmt.Print("Ingrese una opción: ")
		fmt.Scan(&secondOption)
		if secondOption == "1" {
			resp, err := http.Get("http://localhost:5000/api/productos")
			if err != nil {
				log.Fatalln(err)
			}

			defer resp.Body.Close()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)

			// Convert response body to string
			bodyString := string(bodyBytes)
			fmt.Println("API Response as String:\n" + bodyString)
		} else if secondOption == "2" {
			var name string
			var cantidad int
			var precio int
			fmt.Scan(&name)
			fmt.Scan(&cantidad)
			fmt.Scan(&precio)
			todo := Todo{name, cantidad, precio}
			jsonReq, err := json.Marshal(todo)
			resp, err := http.Post("http://localhost:5000/api/productos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
			if err != nil {
				log.Fatalln(err)
			}

			defer resp.Body.Close()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)

			// Convert response body to string
			bodyString := string(bodyBytes)
			fmt.Println(bodyString)
		} else if secondOption == "3" {
			var id string
			fmt.Println("Ingrese id a eliminar ")
			fmt.Scan(&id)
			todo := "1"
			jsonReq, err := json.Marshal(todo)
			req, err := http.NewRequest(http.MethodDelete, "http://localhost:5000/api/productos/"+id, bytes.NewBuffer(jsonReq))
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}

			defer resp.Body.Close()
			bodyBytes, _ := ioutil.ReadAll(resp.Body)

			// Convert response body to string
			bodyString := string(bodyBytes)
			fmt.Println(bodyString)
		} else if secondOption == "4" {

		} else if secondOption == "5" {
			fmt.Print("Hasta luego!")
		}
	} else if firstOption == "3" {
		fmt.Println("Hasta luego!")
	}
}
