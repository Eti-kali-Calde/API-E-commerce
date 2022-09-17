package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//iniciando el servidor
	//server.Server()

	fmt.Print("Opciones:")
	// fmt.Print("1. Iniciar sesión como cliente")
	// fmt.Print("2. Iniciar sesión como administrador")
	// fmt.Print("3. Salir")

	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/tarea_1_sd")
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
}

// func connectDB() (*sql.DB, error) {

// }
