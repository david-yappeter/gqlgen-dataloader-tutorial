package main

import (
	"log"
	"myapp/config"
	"myapp/dataloader"
	"myapp/graph"
	"myapp/graph/generated"
	"myapp/migration"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func init() {
	// Load Environment from .env
	godotenv.Load()

	// Initialize db Value
	config.ConnectGorm()
}

func main() {
	// Defer close database connection (optional)
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Migration
	migration.MigrateTable()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := mux.NewRouter()
	router.Use(dataloader.Middleware)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
