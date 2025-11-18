package internal

type Pedido struct {
	ID       uint    `json:"id"`
	ClienteID uint    `json:"cliente_id"`
	Itens    []ItemCarrito `json:"itens"`
	Data     string `json:"data"`
	Total    float64 `json:"total"`
}
