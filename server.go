package main

import (
	"database/sql"
	db2 "github.com/another-maverick/gql-demo-server/graph/api/db"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/another-maverick/gql-demo-server/graph"
	"github.com/another-maverick/gql-demo-server/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	db, err := db2.Connect()
	if err != nil {
		panic(err)
	}

	initDB(db)



	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// initDB
func initDB(db *sql.DB) {
	db2.MustExec(db, "DROP TABLE IF EXISTS videos")
	db2.MustExec(db,"DROP TABLE IF EXISTS users")
	db2.MustExec(db, "CREATE TABLE public.users (id SERIAL PRIMARY KEY, name varchar(255), email varchar(255))")
	db2.MustExec(db, "CREATE TABLE public.videos (id SERIAL PRIMARY KEY, name varchar(255), " +
		"description varchar(255), url text,created_at TIMESTAMP)")
	db2.MustExec(db, "INSERT INTO users(name, email) VALUES('maverick', 'maverick@test.me')")
	db2.MustExec(db, "INSERT INTO users(name, email) VALUES('vadlakun', 'vadlakun@test.me')")
}
