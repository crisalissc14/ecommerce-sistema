// Login valida las credenciales del usuario
// Retorna error si el usuario no existe o la contraseña es incorrecta

package internal

import (
	"errors"
	"fmt"
)

// Usuario representa un usuario del sistema
type Usuario struct {
	Nombre     string
	Usuario    string
	Contrasena string
	Rol        string // admin o cliente
}

// Simulación de BD de usuarios
var usuarios = []Usuario{
	{"Administrador", "admin", "1234", "admin"},
	{"Cliente Compra", "cliente", "1111", "cliente"},
}

// Constructor encapsulado
func NewUsuario(nombre, usuario, contrasena, rol string) *Usuario {
	return &Usuario{
		Nombre:     nombre,
		Usuario:    usuario,
		Contrasena: contrasena,
		Rol:        rol,
	}
}

// Método para validar contraseña
func (u *Usuario) ValidarContrasena(pass string) bool {
	return u.Contrasena == pass
}

// Login con manejo de errores
func Login(user, pass string) (*Usuario, error) {
	for _, u := range usuarios {
		if u.Usuario == user {
			if u.ValidarContrasena(pass) {
				return &u, nil
			}
			return nil, errors.New("contraseña incorrecta")
		}
	}
	return nil, errors.New("usuario no encontrado")
}

// Mostrar perfil
func (u *Usuario) Info() {
	fmt.Printf("Bienvenido %s (%s)\n", u.Nombre, u.Rol)
}
