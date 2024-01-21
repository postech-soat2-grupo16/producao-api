package entities

import (
	"time"

	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

const (
	StatusPedidoCriado       = "CRIADO"
	StatusPedidoRecebido     = "RECEBIDO"
	StatusPedidoEmPreparacao = "EM_PREPARACAO"
	StatusPedidoPronto       = "PRONTO"
	StatusPedidoEntregue     = "ENTREGUE"
	StatusPedidoFinalizado   = "FINALIZADO"
)

type ProducaoPedido struct {
	PedidoID  string `gorm:"primary_key;not null;"`
	Status    string `gorm:"size:100;not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (i *ProducaoPedido) IsStatusValid() bool {
	status := []string{StatusPedidoCriado,
		StatusPedidoRecebido,
		StatusPedidoEmPreparacao,
		StatusPedidoPronto,
		StatusPedidoEntregue,
		StatusPedidoFinalizado}

	return slices.Contains(status, i.Status)

}

func (i *ProducaoPedido) CopyPedidoWithNewValues(status string) ProducaoPedido {
	return ProducaoPedido{
		PedidoID: i.PedidoID,
		Status:   status,
	}
}
