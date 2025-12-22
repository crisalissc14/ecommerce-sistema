package internal

import (
	"encoding/json"
	"net/http"
)

// 1️⃣ GET /usuarios
func ListarUsuariosAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// 2️⃣ GET /usuario?id=
func ObtenerUsuarioAPI(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	for _, u := range usuarios {
		if u.Usuario == id {
			json.NewEncoder(w).Encode(u)
			return
		}
	}
	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}

// 3️⃣ POST /usuario/crear
func CrearUsuarioAPI(w http.ResponseWriter, r *http.Request) {
	var nuevo Usuario
	json.NewDecoder(r.Body).Decode(&nuevo)

	usuarios = append(usuarios, nuevo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevo)
}

// 4️⃣ PUT /usuario/actualizar (simulado)
func ActualizarUsuarioAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"mensaje":"Usuario actualizado (simulado)"}`))
}

// 5️⃣ DELETE /usuario/eliminar
func EliminarUsuarioAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"mensaje":"Usuario eliminado (simulado)"}`))
}

// 6️⃣ GET /usuarios/rol
func UsuariosPorRolAPI(w http.ResponseWriter, r *http.Request) {
	rol := r.URL.Query().Get("rol")
	var filtrados []Usuario

	for _, u := range usuarios {
		if u.Rol == rol {
			filtrados = append(filtrados, u)
		}
	}
	json.NewEncoder(w).Encode(filtrados)
}

// 7️⃣ GET /usuarios/activos (simulado)
func UsuariosActivosAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"usuarios_activos":2}`))
}

// 8️⃣ GET /estado
func EstadoAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"estado":"Sistema operativo"}`))
}
