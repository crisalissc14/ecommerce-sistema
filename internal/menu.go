package internal

import "fmt"
// --------------------- MENU ADMIN -------------------------
func MenuAdmin() {
	for {
		fmt.Println("\n===== MENÚ ADMINISTRADOR =====")
		fmt.Println("1. Gestión de productos")
		fmt.Println("2. Control de inventario")
		fmt.Println("3. Ver pedidos")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			MenuProductos()
		case 2:
			MenuInventario()
		case 3:
			ListarPedidos()
		case 4:
			fmt.Println("Saliendo del menú administrador...")
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}

// --------------------- MENU CLIENTE -------------------------
func MenuCliente() {
	for {
		fmt.Println("\n===== MENÚ DE COMPRAS =====")
		fmt.Println("1. Ver productos")
		fmt.Println("2. Agregar producto al carrito")
		fmt.Println("3. Ver carrito")
		fmt.Println("4. Confirmar compra")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			ListarProductos()
		case 2:
			AgregarCarritoFlujo()
		case 3:
			MostrarCarrito()
		case 4:
			ConfirmarPedidoFlujo()
		case 5:
			fmt.Println("Volviendo...")
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}
