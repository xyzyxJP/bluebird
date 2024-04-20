package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/xyzyxJP/bluebird/src/graphql"
	"github.com/xyzyxJP/bluebird/src/mock"
	"github.com/xyzyxJP/bluebird/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "4000"

func main() {
	connectionString := "host=db port=5432 user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.ShelfItem{}, &model.ShelfCategory{}, &model.ShelfTag{}, &model.ShelfLocation{})
	db.Create(mock.MockShelfCategory())
	db.Create(mock.MockShelfTag())
	db.Create(mock.MockShelfLocation())

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{DB: db}}))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true})
	handler := c.Handler(http.DefaultServeMux)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
