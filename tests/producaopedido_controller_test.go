package tests

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/postech-soat2-grupo16/producao-api/controllers"
	"github.com/postech-soat2-grupo16/producao-api/interfaces/mocks"
	"github.com/postech-soat2-grupo16/producao-api/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	errUsecaseFailure  = errors.New("ErrUsecaseFailed")
	errUsecaseNotFound = errors.New("ErrUsecaseNotFound")
)

func TestProducaoPedidoController_GetAll_Error(t *testing.T) {
	useCase := new(mocks.MockProducaoPedidoUseCase)
	useCase.On("List").Return(nil, errUsecaseFailure)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/producao_pedidos", nil)

	c := chi.NewRouter()
	controllers.NewProducaoPedidoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "Internal Server Error response is expected")
}

func TestProducaoPedidoController_GetByID_Error(t *testing.T) {
	useCase := new(mocks.MockProducaoPedidoUseCase)
	useCase.On("GetByID", mock.Anything).Return(nil, errUsecaseNotFound)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/producao_pedidos/1", nil)

	c := chi.NewRouter()
	controllers.NewProducaoPedidoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code)
}

func TestProducaoPedidoController_Create_ErrorParse(t *testing.T) {
	useCase := new(mocks.MockProducaoPedidoUseCase)

	res := httptest.NewRecorder()
	badJSON := `{"invalid json`
	req, _ := http.NewRequest("POST", "/producao_pedidos", strings.NewReader(badJSON))
	c := chi.NewRouter()
	controllers.NewProducaoPedidoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code, "Internal Server Error response is expected")

}

func TestProducaoPedidoController_Create_ErrorDomain(t *testing.T) {
	useCase := new(mocks.MockProducaoPedidoUseCase)
	useCase.On("Create", mock.Anything).Return(nil, util.NewErrorDomain("domain"))

	res := httptest.NewRecorder()
	badJSON := `{}`
	req, _ := http.NewRequest("POST", "/producao_pedidos", strings.NewReader(badJSON))
	c := chi.NewRouter()
	controllers.NewProducaoPedidoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnprocessableEntity, res.Code, "Internal Server Error response is expected")

}

func TestProducaoPedidoController_PUT_ErrorParse(t *testing.T) {
	useCase := new(mocks.MockProducaoPedidoUseCase)

	res := httptest.NewRecorder()
	badJSON := `{"invalid json`
	req, _ := http.NewRequest("PUT", "/producao_pedidos/1", strings.NewReader(badJSON))
	c := chi.NewRouter()
	controllers.NewProducaoPedidoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code, "Internal Server Error response is expected")

}

func TestProducaoPedidoController_PUT_ErrorDomain(t *testing.T) {
	useCase := new(mocks.MockProducaoPedidoUseCase)
	useCase.On("Update", mock.Anything, mock.Anything).Return(nil, util.NewErrorDomain("domain"))

	res := httptest.NewRecorder()
	badJSON := `{}`
	req, _ := http.NewRequest("PUT", "/producao_pedidos/1", strings.NewReader(badJSON))
	c := chi.NewRouter()
	controllers.NewProducaoPedidoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusUnprocessableEntity, res.Code, "Internal Server Error response is expected")
}

func TestProducaoPedidoController_Delete_ErrorUseCase(t *testing.T) {
	useCase := new(mocks.MockProducaoPedidoUseCase)
	useCase.On("Delete", mock.Anything).Return(nil, util.NewErrorDomain("domain"))

	res := httptest.NewRecorder()
	badJSON := `{}`
	req, _ := http.NewRequest("DELETE", "/producao_pedidos/1", strings.NewReader(badJSON))
	c := chi.NewRouter()
	controllers.NewProducaoPedidoController(useCase, c)

	c.ServeHTTP(res, req)

	assert.Equal(t, http.StatusInternalServerError, res.Code, "Internal Server Error response is expected")
}
