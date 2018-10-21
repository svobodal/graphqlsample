package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

type body struct {
	Query string `json:"query"`
}

type GraphqlHandler struct {
	Schema graphql.Schema
}

func (handler *GraphqlHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	var body body
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	params := graphql.Params{Schema: handler.Schema, RequestString: body.Query}

	response := graphql.Do(params)
	if len(response.Errors) > 0 {
		fmt.Printf("Query failed: %+v\n", response.Errors)
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: QueryResolver,
	})

	if err != nil {
		panic(err)
	}

	handler := &GraphqlHandler{
		Schema: schema,
	}

	http.Handle("/query", handler)
	fmt.Printf("Listening on :7070\n")
	http.ListenAndServe(":7070", nil)
}
