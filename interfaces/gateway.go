package interfaces

import (
	"github.com/postech-soat2-grupo16/producao-api/entities"
)

type ProducaoPedidoGatewayI interface {
	Save(producaoPedido entities.ProducaoPedido) (*entities.ProducaoPedido, error)
	Update(producaoPedido entities.ProducaoPedido) (*entities.ProducaoPedido, error)
	Delete(producaoPedidoID string) error
	GetByID(producaoPedidoID string) (*entities.ProducaoPedido, error)
	GetAll() ([]entities.ProducaoPedido, error)
	GetByStatus(category string) ([]entities.ProducaoPedido, error)
}
