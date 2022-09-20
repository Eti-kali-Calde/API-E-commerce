package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//iniciando el servidor
	//server.Server()
	//conectando a la base de datos
	var db *sql.DB = connectDB()

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

	type Cliente struct {
		Id       int
		Name     string
		Password string
	}

	type Compra struct {
		Id_Compra  int
		Id_Cliente int
	}

	type Detalle struct {
		Id_Compra   int
		Id_Producto int
		Cantidad    int
		fecha       string
	}

	type Producto struct {
		Id_Producto         int
		Nombre              string
		Cantidad_Disponible int
		Precio_Unitario     int
	}

	fmt.Print(db)
}

func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/tarea_1_sd")
	if err != nil {
		fmt.Println("error validating sql.open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.ping")
		panic(err.Error())
	}
	return db
}
