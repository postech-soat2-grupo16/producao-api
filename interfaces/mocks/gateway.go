package mocks

import (
	"github.com/postech-soat2-grupo16/producao-api/entities"
	"github.com/stretchr/testify/mock"
)

type MockProducaoPedidoGateway struct {
	mock.Mock
}

func (m *MockProducaoPedidoGateway) Save(pedido entities.ProducaoPedido) (*entities.ProducaoPedido, error) {
	args := m.Called(pedido)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.ProducaoPedido), args.Error(1)
}

func (m *MockProducaoPedidoGateway) Update(pedido entities.ProducaoPedido) (*entities.ProducaoPedido, error) {
	args := m.Called(pedido)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.ProducaoPedido), args.Error(1)
}

func (m *MockProducaoPedidoGateway) Delete(pedidoID string) error {
	args := m.Called(pedidoID)
	return args.Error(0)
}

func (m *MockProducaoPedidoGateway) GetByID(pedidoID string) (*entities.ProducaoPedido, error) {
	args := m.Called(pedidoID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.ProducaoPedido), args.Error(1)
}

func (m *MockProducaoPedidoGateway) GetAll() ([]entities.ProducaoPedido, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.ProducaoPedido), args.Error(1)
}

func (m *MockProducaoPedidoGateway) GetByStatus(status string) ([]entities.ProducaoPedido, error) {
	args := m.Called(status)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.ProducaoPedido), args.Error(1)
}

func (m *MockProducaoPedidoGateway) SendMessage(producaopedido *entities.ProducaoPedido) error {
	args := m.Called(producaopedido)
	return args.Error(0)
}
