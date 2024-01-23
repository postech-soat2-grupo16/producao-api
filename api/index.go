package api

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/postech-soat2-grupo16/producao-api/controllers"
	"github.com/postech-soat2-grupo16/producao-api/external"
	producaopedidoGateway "github.com/postech-soat2-grupo16/producao-api/gateways/db/producaopedido"
	"github.com/postech-soat2-grupo16/producao-api/usecases/producaopedido"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	dialector := external.GetPostgresDialector()
	db := external.NewORM(dialector)

	return db
}

func SetupQueue() *sqs.SQS {
	return external.GetSqsClient()
}

func SetupRouter(db *gorm.DB, queue *sqs.SQS) *chi.Mux {
	r := chi.NewRouter()
	r.Use(commonMiddleware)

	mapRoutes(r, db, queue)

	return r
}

func mapRoutes(r *chi.Mux, orm *gorm.DB, queue *sqs.SQS) {
	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler())

	// Injections
	// Gateways
	producaoPedidoGateway := producaopedidoGateway.NewGateway(orm)
	//queueGateway := message.NewGateway(queue)
	// Use cases
	producaoPedidoUseCase := producaopedido.NewUseCase(producaoPedidoGateway)
	// Handlers
	_ = controllers.NewProducaoPedidoController(producaoPedidoUseCase, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
