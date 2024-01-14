package producaopedido

import (
	"log"

	"github.com/postech-soat2-grupo16/producao-api/entities"
	"gorm.io/gorm"
)

type Gateway struct {
	repository *gorm.DB
}

func NewGateway(repository *gorm.DB) *Gateway {
	return &Gateway{repository: repository}
}

func (g *Gateway) Save(item entities.ProducaoPedido) (*entities.ProducaoPedido, error) {
	result := g.repository.Create(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &item, nil
}

func (g *Gateway) Update(item entities.ProducaoPedido) (*entities.ProducaoPedido, error) {
	result := g.repository.Updates(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &item, nil
}

func (g *Gateway) Delete(itemID uint32) error {
	item := entities.ProducaoPedido{
		PedidoID: itemID,
	}
	result := g.repository.Delete(&item)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (g *Gateway) GetByID(pedidoID uint32) (*entities.ProducaoPedido, error) {
	item := entities.ProducaoPedido{
		PedidoID: pedidoID,
	}
	result := g.repository.First(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (g *Gateway) GetByStatus(status string) ([]entities.ProducaoPedido, error) {
	producaoPedido := entities.ProducaoPedido{
		Status: status,
	}
	var producaoPedidos []entities.ProducaoPedido
	result := g.repository.Find(&producaoPedidos, producaoPedido)
	if result.Error != nil {
		return nil, result.Error
	}

	return producaoPedidos, nil
}

func (g *Gateway) GetAll() (producaoPedidos []entities.ProducaoPedido, err error) {
	result := g.repository.Find(&producaoPedidos)
	if result.Error != nil {
		log.Println(result.Error)
		return producaoPedidos, result.Error
	}

	return producaoPedidos, err
}
