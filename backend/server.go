package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Nevermore-FMS/poesitory/backend/auth"
	"github.com/Nevermore-FMS/poesitory/backend/database"
	"github.com/Nevermore-FMS/poesitory/backend/graph"
	"github.com/Nevermore-FMS/poesitory/backend/graph/generated"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.Init()
	auth.Init()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	gqlhandler := cors.New(cors.Options{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	}).Handler(srv)

	gqlhandler = auth.Middleware(gqlhandler)

	http.Handle("/api/playground", playground.Handler("GraphQL playground", "/api/graphql"))
	http.Handle("/api/graphql", gqlhandler)

	http.Handle("/api/github/login", auth.GithubLoginHandler())
	http.Handle("/api/github/callback", auth.GithubCallbackHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
