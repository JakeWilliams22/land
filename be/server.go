package main

import (
  "net/http"
  "encoding/json"
  "fmt"
  "log"
  
  "github.com/graphql-go/graphql"
  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
)

func ping(w http.ResponseWriter, r *http.Request) {
  message := "Hi"
  w.Write([]byte(message))
}

var graphqlPingQuery = `
        {
            landingPages {
                title,
                questions {
                  mcQuestions {
                    question
                  }
                }
            }
        }
    `

func handleGraphql(w http.ResponseWriter, r *http.Request) {
    schema := createSchema()
    query := graphqlPingQuery
    if r.Method == "POST" {
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "ParseForm() err: %v", err)
            return
        }
        fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
        query = r.FormValue("gqlQuery")
        fmt.Fprintf(w, "Query = %s\n", query)
    }
    gqlResponse := queryGraph(graphqlPingQuery, schema)
    rJSON, _ := json.Marshal(gqlResponse)
    fmt.Printf("%s \n", rJSON)
    w.Write([]byte(rJSON))
}

func queryGraph(query string, schema graphql.Schema) *graphql.Result {
    params := graphql.Params{Schema: schema, RequestString: query}
    r := graphql.Do(params)
    if len(r.Errors) > 0 {
        log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
    }
    return r
}

func main() {
  headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
  originsOk := handlers.AllowedOrigins([]string{"*"})
  methodsOk := handlers.AllowedMethods([]string{"GET", "POST"})
  
  r := mux.NewRouter()
  r.HandleFunc("/", ping)
  r.HandleFunc("/graphql", handleGraphql)
  
  if err := http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
    panic(err)
  }
}