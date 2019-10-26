package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/graphql-go/graphql"
)

func landingPages() []LandingPage {
    mc := MCQuestion{
        Question: "How are you",
        Answers:  []string{"good"},
    }

    oe := OpenEndedQuestion{
        Question: "Why",
    }

    t1 := "Hummingbird"
    t2 := "Durgs"

    st1 := "Sing like your favorite artists"
    st2 := "Do some durgs"

    jel1 := JoinEmailList{
        JoinPrompt:     "Keep up to date",
        JoinButtonText: "Join",
    }
    jel2 := JoinEmailList{
        JoinPrompt:     "Find out about durgs",
        JoinButtonText: "Segrada Familia",
    }

    q1 := Questions{
        QuestionsPrompt: "Help us build something for you",
        McQuestions:     []MCQuestion{mc},
    }
    q2 := Questions{
        QuestionsPrompt:    "Help us build something for you",
        OpenEndedQuestions: []OpenEndedQuestion{oe},
    }

    return []LandingPage{
        {
            Title:    &t1,
            SubTitle: &st1,
            JoinEmailList: &jel1,
            Questions: &q1,
        },
        {
            Title:    &t2,
            SubTitle: &st2,
            JoinEmailList: &jel2,
            Questions: &q2,
        },
    }
}

func createSchema() graphql.Schema {
  // Schema
    fields := graphql.Fields{
        "landingPages": &graphql.Field{
            Type: graphql.NewList(landingPageGQL),
            Resolve: func(p graphql.ResolveParams) (interface{}, error) {
                return landingPages(), nil
            },
        },
    }
    rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
    schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
    schema, err := graphql.NewSchema(schemaConfig)
    if err != nil {
        log.Fatalf("failed to create new schema, error: %v", err)
    }
    return schema
}

func main2() {
    schema := createSchema()

    // Query
    query := `
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
    
    // sssassaxcxdsasuv wefrddssu - Milly the dog
    
    params := graphql.Params{Schema: schema, RequestString: query}
    r := graphql.Do(params)
    if len(r.Errors) > 0 {
        log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
    }
    rJSON, _ := json.Marshal(r)
    fmt.Printf("%s \n", rJSON) // {“data”:{“hello”:”world”}}
}