package servers

import (
	"log"
	"net/http"
	"os"

	"github.com/fpessoa64/desafio_clean_arch/internal/handlers/rest"
	"github.com/fpessoa64/desafio_clean_arch/internal/handlers/rest/routes"
	"github.com/fpessoa64/desafio_clean_arch/internal/usecase"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Rest struct {
	port string
}

func NewRest() *Rest {
	restPort := os.Getenv("REST_PORT")
	if restPort == "" {
		restPort = "8080"
	}
	return &Rest{port: restPort}
}

func (rs *Rest) Start(uc *usecase.OrderUsecase) error {
	handler := rest.NewHandler(uc)
	r := chi.NewRouter()
	routes.RegisterOrderRoutes(r, handler)

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	log.Printf("REST API running on :%s", rs.port)
	if err := http.ListenAndServe(":"+rs.port, r); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
