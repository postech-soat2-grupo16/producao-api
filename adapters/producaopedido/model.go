package producaopedido

type ProducaoPedido struct {
	PedidoID uint32 `json:"pedido_id"`
	Status   string `json:"status"`
}
