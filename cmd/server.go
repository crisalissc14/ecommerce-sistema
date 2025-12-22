package main

import (
	"log"
	"net/http"
	"ecommerce-sistema/internal"
)

func main() {
	http.HandleFunc("/usuarios", internal.ListarUsuariosAPI)
	http.HandleFunc("/usuario", internal.ObtenerUsuarioAPI)
	http.HandleFunc("/usuarios", internal.ListarUsuariosAPI)
	http.HandleFunc("/usuario", internal.ObtenerUsuarioAPI)
	http.HandleFunc("/usuario/crear", internal.CrearUsuarioAPI)
	http.HandleFunc("/usuario/actualizar", internal.ActualizarUsuarioAPI)
	http.HandleFunc("/usuario/eliminar", internal.EliminarUsuarioAPI)
	http.HandleFunc("/usuarios/rol", internal.UsuariosPorRolAPI)
	http.HandleFunc("/usuarios/activos", internal.UsuariosActivosAPI)
	http.HandleFunc("/estado", internal.EstadoAPI)


	log.Println("Servidor activo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
