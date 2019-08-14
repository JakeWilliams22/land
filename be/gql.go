package main

import (
    "context"
    "net/http"

    "github.com/samsarahq/thunder/graphql"
    "github.com/samsarahq/thunder/graphql/graphiql"
    "github.com/samsarahq/thunder/graphql/introspection"
    "github.com/samsarahq/thunder/graphql/schemabuilder"
)

type Server struct {
}

type Args struct {
}

func landPageListRet(ctx context.Context, args Args) ([]*LandingPage, error) {
    mc := MCQuestion{
        Question: "How are you",
        Answers:  []string{"good"},
    }

    oe := OpenEndedQuestion{
        Question: "Why",
    }

    return []*LandingPage{
        {
            Title:    "HummingBird",
            SubTitle: "Sing like your favorite artists",
            JoinEmailList: JoinEmailList{
                JoinPrompt:     "Keep up to date",
                JoinButtonText: "Join",
            },
            Questions: Questions{
                QuestionsPrompt: "Help us build something for you",
                McQuestions:     []MCQuestion{mc},
            },
        },
        {
            Title:    "Durgs",
            SubTitle: "Do some durgs",
            JoinEmailList: JoinEmailList{
                JoinPrompt:     "Find out about durgs",
                JoinButtonText: "Segrada Familia",
            },
            Questions: Questions{
                QuestionsPrompt:    "Help us build something for you",
                OpenEndedQuestions: []OpenEndedQuestion{oe},
            },
        },
    }, nil
}

func (s *Server) registerQuery(schema *schemabuilder.Schema) {
    object := schema.Query()

    object.FieldFunc("landingPages", landPageListRet)
}

func (s *Server) registerMutation(schema *schemabuilder.Schema) {
    object := schema.Mutation()

    object.FieldFunc("echo", func(ctx context.Context, args struct{ Text string }) (string, error) {
            return args.Text, nil
    })
}

func (s *Server) Schema() *graphql.Schema {
    schema := schemabuilder.NewSchema()

    s.registerQuery(schema)
    s.registerMutation(schema)

    return schema.MustBuild()
}

func main() {
    server := &Server{}
    graphqlSchema := server.Schema()
    introspection.AddIntrospectionToSchema(graphqlSchema)

    http.Handle("/graphql", graphql.Handler(graphqlSchema))
    http.Handle("/graphiql/", http.StripPrefix("/graphiql/", graphiql.Handler()))

    if err := http.ListenAndServe(":3030", nil); err != nil {
        panic(err)
    }
}