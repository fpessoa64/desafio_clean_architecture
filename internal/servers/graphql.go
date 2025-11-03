package servers

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fpessoa64/desafio_clean_arch/graph"
	"github.com/fpessoa64/desafio_clean_arch/internal/usecase"
	"github.com/vektah/gqlparser/v2/ast"
)

type GraphQL struct {
	port string
}

func NewGraphQL() *GraphQL {
	graphqlPort := os.Getenv("GRAPHQL_PORT")
	if graphqlPort == "" {
		graphqlPort = "8081"
	}
	return &GraphQL{port: graphqlPort}
}

func (gr *GraphQL) Start(uc *usecase.OrderUsecase) error {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL Orders", gr.port)
	log.Fatal(http.ListenAndServe(":"+gr.port, nil))
	return nil
}
