package internal

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
//                    INTERFACES
// ---------------------------------------------------------

// Gestionable define métodos obligatorios para administrar
// entidades dentro del sistema. REPRESENTA ENCAPSULACIÓN + POLIMORFISMO.
type Gestionable interface {
	Crear() error
	Listar()
	Eliminar(id int) error
	Editar(id int) error
}

// ---------------------------------------------------------
//                    STRUCT PRODUCTO
// ---------------------------------------------------------

type Producto struct {
	ID     int
	Nombre string
	Precio float64
	Stock  int
}

// Base de datos en memoria
var productos = []Producto{
	{1, "Pizza Hawaiana", 9.50, 20},
	{2, "Pizza Pepperoni", 10.00, 15},
	{3, "Coca Cola 1L", 2.00, 30},
}

// Constructor
func NewProducto(id int, nombre string, precio float64, stock int) *Producto {
	return &Producto{ID: id, Nombre: nombre, Precio: precio, Stock: stock}
}

// Utilidad para leer texto con espacios
func leerLinea() string {
	reader := bufio.NewReader(os.Stdin)
	texto, _ := reader.ReadString('\n')
	return strings.TrimSpace(texto)
}

// ---------------------------------------------------------
//           MÉTODOS PRINCIPALES (Implementa interfaz)
// ---------------------------------------------------------

// Crear agrega un nuevo producto validando ID duplicado
func (p Producto) Crear() error {
	for _, prod := range productos {
		if prod.ID == p.ID {
			return errors.New("el ID ya existe")
		}
	}
	productos = append(productos, p)
	return nil
}

// Listar imprime todos los productos del sistema
func (p Producto) Listar() {
	ListarProductos()
}

// Eliminar busca un producto y lo borra por ID
func (p Producto) Eliminar(id int) error {
	for i, prod := range productos {
		if prod.ID == id {
			productos = append(productos[:i], productos[i+1:]...)
			return nil
		}
	}
	return errors.New("producto no encontrado")
}

// Editar permite cambiar nombre, precio o stock
func (p Producto) Editar(id int) error {
	for i, prod := range productos {
		if prod.ID == id {

			fmt.Printf("Nuevo nombre (%s): ", prod.Nombre)
			nuevoNombre := leerLinea()
			if nuevoNombre != "" {
				prod.Nombre = nuevoNombre
			}

			fmt.Printf("Nuevo precio (%.2f): ", prod.Precio)
			var nuevoPrecio float64
			_, err := fmt.Scanln(&nuevoPrecio)
			if err == nil {
				prod.Precio = nuevoPrecio
			}

			fmt.Printf("Nuevo stock (%d): ", prod.Stock)
			var nuevoStock int
			_, err = fmt.Scanln(&nuevoStock)
			if err == nil {
				prod.Stock = nuevoStock
			}

			productos[i] = prod
			return nil
		}
	}
	return errors.New("producto no encontrado")
}

// ---------------------------------------------------------
//                 FUNCIONES DE FLUJO
// ---------------------------------------------------------

// Muestra todos los productos
func ListarProductos() {
	fmt.Println("\n=== LISTA DE PRODUCTOS ===")
	for _, p := range productos {
		fmt.Printf("ID: %d | %s | Precio: %.2f | Stock: %d\n",
			p.ID, p.Nombre, p.Precio, p.Stock)
	}
}

// Crear producto pidiendo datos al usuario
func CrearProductoFlujo() {
	fmt.Print("ID: ")
	var id int
	fmt.Scanln(&id)

	fmt.Print("Nombre: ")
	nombre := leerLinea()

	fmt.Print("Precio: ")
	var precio float64
	fmt.Scanln(&precio)

	fmt.Print("Stock: ")
	var stock int
	fmt.Scanln(&stock)

	p := NewProducto(id, nombre, precio, stock)

	if err := p.Crear(); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Producto agregado correctamente.")
}

// Eliminar un producto
func EliminarProductoFlujo() {
	var id int
	fmt.Print("ID del producto a eliminar: ")
	fmt.Scanln(&id)

	var p Producto
	if err := p.Eliminar(id); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Producto eliminado.")
}

// Editar un producto
func EditarProductoFlujo() {
	var id int
	fmt.Print("ID del producto a editar: ")
	fmt.Scanln(&id)

	var p Producto
	if err := p.Editar(id); err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Producto editado correctamente.")
}

// ---------------------------------------------------------
//                     MENÚ COMPLETO
// ---------------------------------------------------------

func MenuProductos() {
	for {
		fmt.Println("\n===== MENÚ PRODUCTOS =====")
		fmt.Println("1. Listar productos")
		fmt.Println("2. Crear producto")
		fmt.Println("3. Eliminar producto")
		fmt.Println("4. Editar producto")
		fmt.Println("5. Volver")
		fmt.Print("Opción: ")

		var op int
		fmt.Scanln(&op)

		switch op {
		case 1:
			ListarProductos()
		case 2:
			CrearProductoFlujo()
		case 3:
			EliminarProductoFlujo()
		case 4:
			EditarProductoFlujo()
		case 5:
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}
