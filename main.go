package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/postech-soat2-grupo16/producao-api/api"
	_ "github.com/postech-soat2-grupo16/producao-api/docs"
)

//	@title			Producao API
//	@version		1.0
//	@description	Here you will find everything you need to have the best possible integration with our APIs.
//	@termsOfService	http://fastfood.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.fastfood.io/support
//	@contact.email	support@fastfood.io

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	db := api.SetupDB()
	queue := api.SetupQueue()
	r := api.SetupRouter(db, queue)

	server := &http.Server{
		Addr:              ":8000",
		ReadHeaderTimeout: 3 & time.Second,
		Handler:           r,
	}
	log.Println(server.ListenAndServe())
}
