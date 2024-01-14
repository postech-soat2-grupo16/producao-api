package producaopedido

import (
	"errors"
	"log"
	"strings"

	"github.com/postech-soat2-grupo16/producao-api/entities"
	"github.com/postech-soat2-grupo16/producao-api/interfaces"
	"github.com/postech-soat2-grupo16/producao-api/util"

	"gorm.io/gorm"
)

type UseCase struct {
	producaoPedidoGateway interfaces.ProducaoPedidoGatewayI
}

func NewUseCase(producaoPedidoGateway interfaces.ProducaoPedidoGatewayI) UseCase {
	return UseCase{
		producaoPedidoGateway: producaoPedidoGateway,
	}
}

func (p UseCase) List() ([]entities.ProducaoPedido, error) {
	pedidos, err := p.producaoPedidoGateway.GetAll()
	if err != nil {
		log.Println(err)
		return pedidos, err
	}

	return pedidos, err
}

func (p UseCase) GetByID(pedidoID uint32) (*entities.ProducaoPedido, error) {
	result, err := p.producaoPedidoGateway.GetByID(pedidoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (p UseCase) GetByStatus(status string) ([]entities.ProducaoPedido, error) {
	pedido := entities.ProducaoPedido{
		Status: strings.ToUpper(status),
	}

	if !pedido.IsStatusValid() {
		return nil, util.NewErrorDomain("Categoria inválida")
	}

	result, err := p.producaoPedidoGateway.GetByStatus(pedido.Status)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entities.ProducaoPedido{}, nil
		}
		log.Println(err)
		return nil, err
	}

	return result, nil
}

func (p UseCase) Create(pedidoID uint32) (*entities.ProducaoPedido, error) {
	pedido := entities.ProducaoPedido{
		PedidoID: pedidoID,
		Status:   entities.StatusPedidoCriado,
	}

	if !pedido.IsStatusValid() {
		return nil, util.NewErrorDomain("Status inválido")
	}

	result, err := p.producaoPedidoGateway.Save(pedido)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (p UseCase) Update(pedidoID uint32, status string) (*entities.ProducaoPedido, error) {
	pedido, err := p.producaoPedidoGateway.GetByID(pedidoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	updatedItem := pedido.CopyPedidoWithNewValues(status)

	if !updatedItem.IsStatusValid() {
		return nil, util.NewErrorDomain("Status inválido")
	}

	result, err := p.producaoPedidoGateway.Update(updatedItem)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func (p UseCase) Delete(pedidoID uint32) (*entities.ProducaoPedido, error) {
	pedido, err := p.producaoPedidoGateway.GetByID(pedidoID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	err = p.producaoPedidoGateway.Delete(pedido.PedidoID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return pedido, nil
}
