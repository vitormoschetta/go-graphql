package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vitormoschetta/go-graphql/graph"
	"github.com/vitormoschetta/go-graphql/internal/database"
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "./internal/database/database.db")
	if err != nil {
		log.Fatal("error opening the database connection: ", err)
	}
	defer db.Close()

	categoryDb := database.NewCategory(db)
	productDb := database.NewProduct(db)

	executableSchema := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDb,
		ProductDB:  productDb,
	}})

	srv := handler.NewDefaultServer(executableSchema)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
