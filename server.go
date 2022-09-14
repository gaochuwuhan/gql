package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gaochuwuhan/gql/graph/generated"
	"github.com/gaochuwuhan/gql/graph/resolver"
)

const defaultPort = "9090"
const defaultUri = "mongodb://mongou:mongou@127.0.0.1:27017/nisar-module"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	mongo_uri := os.Getenv("MONGO_URI")
	if mongo_uri == "" {
		mongo_uri = defaultUri
	}
	cfg, err := resolver.New(mongo_uri)

	if err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(*cfg))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http路由是 /query 只有这一个，post方法
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
