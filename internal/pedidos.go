package internal

import "fmt"

// Pedido básico
type Pedido struct {
	ID      int
	Items   []Producto
	Total   float64
}

var pedidos []Pedido
var contadorPedidos = 1

func ConfirmarPedidoFlujo() {
	if len(carrito) == 0 {
		fmt.Println("El carrito está vacío.")
		return
	}

	total := 0.0
	for _, p := range carrito {
		total += p.Precio
	}

	pedido := Pedido{
		ID:    contadorPedidos,
		Items: carrito,
		Total: total,
	}

	pedidos = append(pedidos, pedido)
	contadorPedidos++
	carrito = []Producto{}

	fmt.Println("Pedido confirmado correctamente.")
}

func ListarPedidos() {
	fmt.Println("\n=== LISTA DE PEDIDOS ===")

	if len(pedidos) == 0 {
		fmt.Println("No hay pedidos registrados.")
		return
	}

	for _, p := range pedidos {
		fmt.Printf("\nPedido #%d | Total: %.2f\n", p.ID, p.Total)
		fmt.Println("Items:")
		for _, it := range p.Items {
			fmt.Printf("- %s (%.2f)\n", it.Nombre, it.Precio)
		}
	}
}
