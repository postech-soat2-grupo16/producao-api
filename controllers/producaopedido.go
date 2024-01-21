package controllers

import (
	"encoding/json"
	"net/http"

	producaopedido "github.com/postech-soat2-grupo16/producao-api/adapters/producaopedido"
	"github.com/postech-soat2-grupo16/producao-api/interfaces"
	"github.com/postech-soat2-grupo16/producao-api/util"

	"github.com/go-chi/chi/v5"
)

type ProducaoPedidoController struct {
	useCase interfaces.ProducaoPedidoUseCase
}

func NewProducaoPedidoController(useCase interfaces.ProducaoPedidoUseCase, r *chi.Mux) *ProducaoPedidoController {
	controller := ProducaoPedidoController{useCase: useCase}
	r.Route("/producao_pedidos", func(r chi.Router) {
		r.Get("/", controller.GetAll())
		r.Post("/", controller.Create())
		r.Get("/{id}", controller.GetByID())
		r.Put("/{id}", controller.Update())
		r.Delete("/{id}", controller.Delete())
	})
	return &controller
}

//	@Summary	Get all producao_pedidos
//
//	@Tags		ProducaoPedidos
//
//	@ID			get-all-producao_pedidos
//
// @Param        status    query     string  false  "status search by status"
//
//	@Produce	json
//
// @Success	200	{object}	producaopedido.ProducaoPedido
// @Failure	500
// @Router		/producao_pedidos [get]
func (c *ProducaoPedidoController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result interface{}
		var err error
		if r.URL.Query().Get("status") != "" {
			result, err = c.useCase.GetByStatus(r.URL.Query().Get("status"))
		} else {
			result, err = c.useCase.List()
		}
		if err != nil {
			if util.IsDomainError(err) {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(err)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(result)
	}
}

// @Summary	Get a producaopedido by ID
//
// @Tags		ProducaoPedidos
//
// @ID			get-producaopedido-by-id
// @Produce	json
// @Param		id	path		string	true	"ProducaoPedido ID"
// @Success	200	{object}	producaopedido.ProducaoPedido
// @Failure	404
// @Router		/producao_pedidos/{id} [get]
func (c *ProducaoPedidoController) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		producaoPedido, err := c.useCase.GetByID(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if producaoPedido == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(producaoPedido)
	}
}

// @Summary	New producaopedido
//
// @Tags		ProducaoPedidos
//
// @ID			create-producaopedido
// @Produce	json
// @Param		data	body		producaopedido.ProducaoPedido	true	"ProducaoPedido data"
// @Success	200		{object}	producaopedido.ProducaoPedido
// @Failure	400
// @Router		/producao_pedidos [post]
func (c *ProducaoPedidoController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i producaopedido.ProducaoPedido
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		producaoPedido, err := c.useCase.Create(i.PedidoID)
		if err != nil {
			if util.IsDomainError(err) {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(err)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(producaoPedido)
	}
}

// @Summary	Update a producaopedido
//
// @Tags		ProducaoPedidos
//
// @ID			update-producaopedido
// @Produce	json
// @Param		id		path		string	true	"ProducaoPedido ID"
// @Param		data	body		producaopedido.ProducaoPedido	true	"ProducaoPedido data"
// @Success	200		{object}	producaopedido.ProducaoPedido
// @Failure	404
// @Failure	400
// @Router		/producao_pedidos/{id} [put]
func (c *ProducaoPedidoController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i producaopedido.ProducaoPedido
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		idStr := chi.URLParam(r, "id")

		producaoPedido, err := c.useCase.Update(idStr, i.Status)
		if err != nil {
			if util.IsDomainError(err) {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(err)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if producaoPedido == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(producaoPedido)
	}
}

// @Summary	Delete a producaopedido by ID
//
// @Tags		ProducaoPedidos
//
// @ID			delete-producaopedido-by-id
// @Produce	json
// @Param		id	path	string	true	"ProducaoPedido ID"
// @Success	204
// @Failure	500
// @Router		/producao_pedidos/{id} [delete]
func (c *ProducaoPedidoController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		producaoPedido, err := c.useCase.Delete(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if producaoPedido == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
