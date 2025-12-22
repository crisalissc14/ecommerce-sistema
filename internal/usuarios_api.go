// Usuario representa una entidad del sistema
// Aplica encapsulación mediante métodos


package internal

import (
	"encoding/json"
	"net/http"
)

// 1️⃣ GET /usuarios
func ListarUsuariosAPI(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(usuarios)
}

// 2️⃣ GET /usuario?user=
func ObtenerUsuarioAPI(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")

	for _, u := range usuarios {
		if u.Usuario == user {
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

// 4️⃣ PUT /usuario/actualizar?user=
func ActualizarUsuarioAPI(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")
	var datos Usuario
	json.NewDecoder(r.Body).Decode(&datos)

	for i, u := range usuarios {
		if u.Usuario == user {
			usuarios[i].Nombre = datos.Nombre
			usuarios[i].Rol = datos.Rol
			json.NewEncoder(w).Encode(usuarios[i])
			return
		}
	}
	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}

// 5️⃣ DELETE /usuario/eliminar?user=
func EliminarUsuarioAPI(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("user")

	for i, u := range usuarios {
		if u.Usuario == user {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			w.Write([]byte(`{"mensaje":"Usuario eliminado"}`))
			return
		}
	}
	http.Error(w, "Usuario no encontrado", http.StatusNotFound)
}

// 6️⃣ GET /usuarios/rol?rol=
func UsuariosPorRolAPI(w http.ResponseWriter, r *http.Request) {
	rol := r.URL.Query().Get("rol")
	var resultado []Usuario

	for _, u := range usuarios {
		if u.Rol == rol {
			resultado = append(resultado, u)
		}
	}
	json.NewEncoder(w).Encode(resultado)
}

// 7️⃣ GET /usuarios/activos
func UsuariosActivosAPI(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]int{
		"usuarios_activos": len(usuarios),
	})
}

// 8️⃣ GET /estado
func EstadoAPI(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"estado": "Sistema operativo",
	})
}
