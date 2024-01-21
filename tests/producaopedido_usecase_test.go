package tests

import (
	"errors"
	"testing"

	"github.com/postech-soat2-grupo16/producao-api/interfaces/mocks"
	"github.com/postech-soat2-grupo16/producao-api/usecases/producaopedido"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUseCaseListFail(t *testing.T) {
	mockGateway := new(mocks.MockProducaoPedidoGateway)
	mockGateway.On("GetAll").Return(nil, errors.New("database error"))
	usecase := producaopedido.NewUseCase(mockGateway)

	_, err := usecase.List()
	assert.Error(t, err)
}

func TestUseCaseGetByIDFail(t *testing.T) {
	mockGateway := new(mocks.MockProducaoPedidoGateway)
	mockGateway.On("GetByID", mock.Anything).Return(nil, errors.New("database error"))
	usecase := producaopedido.NewUseCase(mockGateway)

	_, err := usecase.GetByID("12345")
	assert.Error(t, err)
}

func TestUseCaseGetByStatusFail(t *testing.T) {
	mockGateway := new(mocks.MockProducaoPedidoGateway)
	mockGateway.On("GetByStatus", mock.Anything).Return(nil, errors.New("database error"))
	usecase := producaopedido.NewUseCase(mockGateway)

	_, err := usecase.GetByStatus("INVALID_STATUS")
	assert.Error(t, err)
}

func TestUseCaseCreateFail(t *testing.T) {
	mockGateway := new(mocks.MockProducaoPedidoGateway)
	mockGateway.On("Save", mock.Anything).Return(nil, errors.New("database error"))
	usecase := producaopedido.NewUseCase(mockGateway)

	_, err := usecase.Create("12345")
	assert.Error(t, err)
}

func TestUseCaseUpdateFail(t *testing.T) {
	mockGateway := new(mocks.MockProducaoPedidoGateway)
	mockGateway.On("GetByID", mock.Anything).Return(nil, errors.New("database error"))
	usecase := producaopedido.NewUseCase(mockGateway)

	_, err := usecase.Update("12345", "INVALID_STATUS")
	assert.Error(t, err)
}

func TestUseCaseDeleteFail(t *testing.T) {
	mockGateway := new(mocks.MockProducaoPedidoGateway)
	mockGateway.On("GetByID", mock.Anything).Return(nil, errors.New("database error"))
	mockGateway.On("Delete", mock.Anything).Return(errors.New("database error"))
	usecase := producaopedido.NewUseCase(mockGateway)

	_, err := usecase.Delete("12345")
	assert.Error(t, err)
}
