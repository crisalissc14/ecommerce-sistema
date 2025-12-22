package internal

import "fmt"

// Ejemplo simple de goroutines
func ProcesarUsuariosConcurrente(lista []UsuarioAPI) {
	c := make(chan string)

	for _, u := range lista {
		go func(user UsuarioAPI) {
			c <- fmt.Sprintf("Procesado %s", user.Usuario)
		}(u)
	}

	for range lista {
		fmt.Println(<-c)
	}
}
