package mocks

import (
	"github.com/postech-soat2-grupo16/producao-api/entities"
	"github.com/stretchr/testify/mock"
)

type MockProducaoPedidoUseCase struct {
	mock.Mock
}

func (m *MockProducaoPedidoUseCase) List() ([]entities.ProducaoPedido, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.ProducaoPedido), args.Error(1)
}

func (m *MockProducaoPedidoUseCase) Create(pedidoID string) (*entities.ProducaoPedido, error) {
	args := m.Called(pedidoID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.ProducaoPedido), args.Error(1)
}

func (m *MockProducaoPedidoUseCase) GetByID(pedidoID string) (*entities.ProducaoPedido, error) {
	args := m.Called(pedidoID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.ProducaoPedido), args.Error(1)
}

func (m *MockProducaoPedidoUseCase) GetByStatus(status string) ([]entities.ProducaoPedido, error) {
	args := m.Called(status)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.ProducaoPedido), args.Error(1)
}

func (m *MockProducaoPedidoUseCase) Update(pedidoID string, status string) (*entities.ProducaoPedido, error) {
	args := m.Called(pedidoID, status)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.ProducaoPedido), args.Error(1)
}

func (m *MockProducaoPedidoUseCase) Delete(pedidoID string) (*entities.ProducaoPedido, error) {
	args := m.Called(pedidoID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.ProducaoPedido), args.Error(1)
}
