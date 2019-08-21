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

type LandingPageMutation struct {
  LandingPage LandingPage
}

func landPageListRet(ctx context.Context, args Args) ([]*LandingPage, error) {
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

    return []*LandingPage{
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
    }, nil
}

func (s *Server) registerQuery(schema *schemabuilder.Schema) {
    object := schema.Query()

    object.FieldFunc("landingPages", landPageListRet)
}

func (s *Server) registerMutation(schema *schemabuilder.Schema) {
    object := schema.Mutation()

    object.FieldFunc("landingPage", func(ctx context.Context, args LandingPageMutation) (LandingPage, error) {
            return args.LandingPage, nil
    })
}

func (s *Server) Schema() *graphql.Schema {
    schema := schemabuilder.NewSchema()

    s.registerQuery(schema)
    s.registerMutation(schema)

    // sssassaxcxdsasuv wefrddssu - Milly

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