package internal

import (
	"fmt"
)

// Actualizar stock de un producto
func ActualizarStock(id, nuevoStock int) error {
	for i := range productos {
		if productos[i].ID == id {
			productos[i].Stock = nuevoStock
			return nil
		}
	}
	return fmt.Errorf("producto con ID %d no encontrado", id)
}

// Verifica productos con bajo stock
func RevisarStockBajo() {
	fmt.Println("\n=== PRODUCTOS CON STOCK BAJO ===")
	for _, p := range productos {
		if p.Stock < 5 {
			fmt.Printf("⚠️  %s (ID %d) | Stock: %d\n", p.Nombre, p.ID, p.Stock)
		}
	}
}

// Menú de control de inventario
func MenuInventario() {
	for {
		fmt.Println("\n===== CONTROL DE INVENTARIO =====")
		fmt.Println("1. Actualizar stock")
		fmt.Println("2. Revisar stock bajo")
		fmt.Println("3. Volver")
		fmt.Print("Opción: ")

		var op int
		fmt.Scanln(&op)

		switch op {
		case 1:
			var id, stock int
			fmt.Print("ID del producto: ")
			fmt.Scanln(&id)
			fmt.Print("Nuevo stock: ")
			fmt.Scanln(&stock)
			if err := ActualizarStock(id, stock); err != nil {
				fmt.Println("ERROR:", err)
			} else {
				fmt.Println("Stock actualizado.")
			}

		case 2:
			RevisarStockBajo()

		case 3:
			return

		default:
			fmt.Println("Opción inválida")
		}
	}
}
