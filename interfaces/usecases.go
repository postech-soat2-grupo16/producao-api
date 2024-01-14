package interfaces

import (
	"github.com/postech-soat2-grupo16/producao-api/entities"
)

type ProducaoPedidoUseCase interface {
	List() ([]entities.ProducaoPedido, error)
	Create(pedidoID uint32) (*entities.ProducaoPedido, error)
	GetByID(pedidoID uint32) (*entities.ProducaoPedido, error)
	GetByStatus(status string) ([]entities.ProducaoPedido, error)
	Update(pedidoID uint32, status string) (*entities.ProducaoPedido, error)
	Delete(itemID uint32) (*entities.ProducaoPedido, error)
}
