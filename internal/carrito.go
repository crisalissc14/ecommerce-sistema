package internal

type Carrito struct {
	Itens []ItemCarrito `json:"itens"`
}

type ItemCarrito struct {
	ProdutoID uint `json:"produto_id"`
	Quantidade uint `json:"quantidade"`
}
