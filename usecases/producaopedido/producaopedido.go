package producaopedido

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/postech-soat2-grupo16/producao-api/entities"
	"github.com/postech-soat2-grupo16/producao-api/interfaces"
	"github.com/postech-soat2-grupo16/producao-api/util"

	"gorm.io/gorm"
)

type UseCase struct {
	producaoPedidoGateway interfaces.ProducaoPedidoGatewayI
	messageGateway        interfaces.QueueGatewayI
}

func NewUseCase(producaoPedidoGateway interfaces.ProducaoPedidoGatewayI, messageGateway interfaces.QueueGatewayI) UseCase {
	return UseCase{
		producaoPedidoGateway: producaoPedidoGateway,
		messageGateway:        messageGateway,
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

func (p UseCase) GetByID(pedidoID string) (*entities.ProducaoPedido, error) {
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

func (p UseCase) Create(pedidoID string) (*entities.ProducaoPedido, error) {
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

	fmt.Printf("Produção Criada para o Pedido %s", result.PedidoID)

	return result, nil
}

func (p UseCase) Update(pedidoID string, status string) (*entities.ProducaoPedido, error) {
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

	fmt.Printf("Status da produção do pedido %s atualizado para %s", result.PedidoID, result.Status)

	p.messageGateway.SendMessage(result)

	return result, nil
}

func (p UseCase) Delete(pedidoID string) (*entities.ProducaoPedido, error) {
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
