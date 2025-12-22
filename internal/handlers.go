package internal

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Modelo API
type UsuarioAPI struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Usuario string `json:"usuario"`
	Rol    string `json:"rol"`
	Activo bool   `json:"activo"`
}

// BD en memoria
var usuariosAPI = []UsuarioAPI{
	{1, "Administrador", "admin", "admin", true},
	{2, "Cliente Compra", "cliente", "cliente", true},
}

// GET /usuarios
func ListarUsuariosAPI(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(usuariosAPI)
}

// GET /usuario?id=1
func ObtenerUsuarioAPI(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	for _, u := range usuariosAPI {
		if u.ID == id {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}
