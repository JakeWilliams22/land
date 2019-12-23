package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

func ping(w http.ResponseWriter, r *http.Request) {
	message := "Hi"
	w.Write([]byte(message))
}

var graphqlPingQuery = `
        {
            landingPages {
                title
            }
        }
    `

func testGraphql(w http.ResponseWriter, r *http.Request) {
	schema := createSchema()
	gqlResponse := queryGraph(graphqlPingQuery, schema)
	rJSON, _ := json.Marshal(gqlResponse)
	w.Write([]byte(rJSON))
}

func handleGraphql(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		schema := createSchema()
		decoder := json.NewDecoder(r.Body)
		var postBody GqlRequest
		if err := decoder.Decode(&postBody); err != nil {
			panic(err)
		}
		query := postBody.GqlQuery

		gqlResponse := queryGraph(query, schema)
		log.Println("Query " + query)
		rJSON, _ := json.Marshal(gqlResponse)
		fmt.Printf("responseJson: %s \n", rJSON)
		w.Write([]byte(rJSON))
	} else {
		panic("This function requires a POST request")
	}
}

func queryGraph(query string, schema graphql.Schema) *graphql.Result {
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	return r
}

func startServer() {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})

	r := mux.NewRouter()
	r.HandleFunc("/", ping)
	r.HandleFunc("/graphql", handleGraphql)
	r.HandleFunc("/testGraphql", testGraphql)

	if err := http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
		panic(err)
	}
}
