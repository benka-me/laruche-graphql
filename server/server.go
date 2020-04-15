package server

import (
	"fmt"
	"github.com/benka-me/laruche-graphql/go-pkg/http/rpc"
	"github.com/benka-me/laruche/go-pkg/discover"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/benka-me/laruche-graphql/graph"
	"github.com/benka-me/laruche-graphql/graph/generated"
)

const defaultPort = "8088"

func Run(engine *discover.Engine) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	resolver := &graph.Resolver{
		Clients: rpc.InitClients(*engine, grpc.WithInsecure()),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Query")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Origin", "*")

		handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver})).ServeHTTP(w, r)
	})

	log.Printf("connect to http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	return nil
}
