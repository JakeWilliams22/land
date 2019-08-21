package main

import (
    "fmt"

    "github.com/samsarahq/thunder/graphql/introspection"
    "github.com/samsarahq/thunder/graphql/schemabuilder"
)


func main2() {
  //Instantiate a server and run the introspection query on it.
  server := &Server{}

  builderSchema := schemabuilder.NewSchema()
  server.registerQuery(builderSchema)
  server.registerMutation(builderSchema)
  //...

  valueJSON, err := introspection.ComputeSchemaJSON(*builderSchema)
  if err != nil {
    panic(err)
  }

  fmt.Print(string(valueJSON))
}