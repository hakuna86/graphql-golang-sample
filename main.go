package main

import (
	"fmt"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/graph-gophers/graphql-transport-ws/graphqlws"
	schema_ws "github.com/hakuna86/graphql-golang-sample/schema-ws"
	"net/http"
)

func main() {
	//opts := []graphql.SchemaOpt{graphql.UseFieldResolvers(), graphql.MaxParallelism(20)}
	//http.Handle("/graphql", &relay.Handler{Schema: graphql.MustParseSchema(schema.Schema, &schema.Resolver{}, opts...)})
	//http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.Write(template)
	//}))
	//log.Fatal(http.ListenAndServe(":8080", nil))


	// init graphQL schema
	s, err := graphql.ParseSchema(schema_ws.Schema, schema_ws.NewResolver())
	if err != nil {
		panic(err)
	}

	// graphQL handler
	graphQLHandler := graphqlws.NewHandlerFunc(s, &relay.Handler{Schema: s})
	http.HandleFunc("/graphql", graphQLHandler)
	http.HandleFunc("/", http.HandlerFunc(temp))

	// start HTTP server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil); err != nil {
		panic(err)
	}

}