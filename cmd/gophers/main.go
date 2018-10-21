package main

import (
	"io/ioutil"
	"log"
	"net/http"

	graphqlsample "../.."
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	schemaData, err := ioutil.ReadFile("../../schema.graphql")
	if err != nil {
		panic(err)
	}
	schema := graphql.MustParseSchema(string(schemaData), &graphqlsample.Query{})

	log.Printf("Listening on :7070")
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":7070", nil))
}
