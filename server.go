package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"

	"github.com/jmagoga/new-equimper-go-graphql/graph"
	"github.com/jmagoga/new-equimper-go-graphql/graph/generated"
	"github.com/jmagoga/new-equimper-go-graphql/postgres"
)

const defaultPort = "8081"

func main() {

	// CORS begin

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"http://locahost:8081"},
		AllowCredentials: true,
		Debug: true,
	}).Handler)

	// CORS end

	DB := postgres.New(&pg.Options{
		User: "postgres",
		Password: "secret",
		Database: "meetmeup",
	})
	
	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := generated.Config{Resolvers: &graph.Resolver{
		RequestsRepo: postgres.RequestsRepo{DB: DB},
		BeesRepo: postgres.BeesRepo{DB: DB},
	}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	// CORS begin2 
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// check agains your desired domains here
				return r.Host == "locahost:3000"
			},
			ReadBufferSize: 1024,
			WriteBufferSize: 1024,
		},
	})
	// CORS end2

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
