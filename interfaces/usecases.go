package interfaces

import (
	"github.com/postech-soat2-grupo16/producao-api/entities"
)

type ProducaoPedidoUseCase interface {
	List() ([]entities.ProducaoPedido, error)
	Create(pedidoID string) (*entities.ProducaoPedido, error)
	GetByID(pedidoID string) (*entities.ProducaoPedido, error)
	GetByStatus(status string) ([]entities.ProducaoPedido, error)
	Update(pedidoID string, status string) (*entities.ProducaoPedido, error)
	Delete(pedidoID string) (*entities.ProducaoPedido, error)
}
