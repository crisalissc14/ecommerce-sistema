package internal

// ======================================================================
//  INTERFACES DEL SISTEMA — REQUISITO OBLIGATORIO SEGÚN LA TAREA
// ======================================================================

// Gestiona usuarios
type IUsuarioManager interface {
    CrearUsuario(nombre string) error
    EliminarUsuario(id int) error
    ListarUsuarios()
}

// Gestiona productos
type IProductoManager interface {
    RegistrarProducto() error
    ListarProductos()
    BuscarProducto()
    EliminarProducto()
}

// Gestiona inventario
type IInventarioManager interface {
    MostrarInventario()
    AumentarStock()
    DisminuirStock()
}

// Gestiona el carrito
type ICarrito interface {
    Agregar(id int, cantidad int) error
    Mostrar()
    Vaciar()
    Total() float64
}

// Gestiona los pedidos
type IPedidoManager interface {
    CrearPedido(cliente, producto string, cantidad int, precio float64) Pedido
    ListarPedidos() []Pedido
    BuscarPedido(id int) (*Pedido, error)
    CambiarEstado(id int, nuevoEstado string) error
    ImprimirPedido(id int) error
}
