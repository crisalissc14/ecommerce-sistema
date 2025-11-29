package internal

import "fmt"

// Carrito del cliente (siempre uno)
var carrito []Producto

func AgregarCarritoFlujo() {
	var id int
	fmt.Print("ID del producto: ")
	fmt.Scanln(&id)

	for _, p := range productos {
		if p.ID == id {
			carrito = append(carrito, p)
			fmt.Println("Producto agregado al carrito.")
			return
		}
	}
	fmt.Println("Producto no encontrado.")
}

func MostrarCarrito() {
	fmt.Println("\n=== CARRITO ===")
	if len(carrito) == 0 {
		fmt.Println("Carrito vacío.")
		return
	}

	total := 0.0
	for _, p := range carrito {
		fmt.Printf("%s - %.2f\n", p.Nombre, p.Precio)
		total += p.Precio
	}

	fmt.Printf("\nTOTAL: %.2f\n", total)
}
