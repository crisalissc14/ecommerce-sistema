package main

import (
	"log"
	"net/http"
	"ecommerce-sistema/internal"
)

func main() {
	http.HandleFunc("/usuarios", internal.ListarUsuariosAPI)
	http.HandleFunc("/usuario", internal.ObtenerUsuarioAPI)
	

	log.Println("Servidor activo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
