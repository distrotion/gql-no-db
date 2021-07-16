package main

import (
	"log"
	"minefit_auth/graph"
	"minefit_auth/graph/generated"
	"minefit_auth/internal/auth"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//---------------------- route -----------------------------

	router := chi.NewRouter()
	router.Use(auth.Middleware())

	//---------------------- route -----------------------------

	//---------------------- db -----------------------------

	// maindb.Getcol()
	// singoutdb.Getcol()

	//---------------------- db -----------------------------

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
