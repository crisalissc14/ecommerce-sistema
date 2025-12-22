package main
// main.go (consola)
// Punto de entrada del sistema por consola.
// Se encarga del login y del control de acceso por rol.


import (
	"fmt"
	"ecommerce-sistema/internal"
)

func main() {
	fmt.Println("===== LOGIN =====")

	var user, pass string
	fmt.Print("Usuario: ")
	fmt.Scanln(&user)
	fmt.Print("Contraseña: ")
	fmt.Scanln(&pass)

	u, err := internal.Login(user, pass)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	u.Info()

	if u.Rol == "admin" {
		internal.MenuAdmin()
	} else {
		internal.MenuCliente()
	}
}
