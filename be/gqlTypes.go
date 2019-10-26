package main

import (
    "github.com/graphql-go/graphql"
)

var landingPageGQL = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "LandingPage",
        Fields: graphql.Fields{
            "title": &graphql.Field{
                Type: graphql.String,
            },
            "subTitle": &graphql.Field{
                Type: graphql.String,
            },
            "joinEmailList": &graphql.Field{
                Type: joinEmailListGQL,
            },
            "questions": &graphql.Field{
                Type: questionsGQL,
            },
        },
    },
)

var joinEmailListGQL = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "JoinEmailList",
        Fields: graphql.Fields{
            "joinPrompt": &graphql.Field{
                Type: graphql.String,
            },
            "joinButtonText": &graphql.Field{
                Type: graphql.String,
            },
        },
    },
)

var questionsGQL = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "Questions",
        Fields: graphql.Fields{
            "questionsPrompt": &graphql.Field{
              Type: graphql.String,
            },
            "mcQuestions": &graphql.Field{
              Type: graphql.NewList(mcQuestionGraphQL),
            },
            "oeQuestions": &graphql.Field{
              Type: graphql.NewList(oeQuestionGraphQL),
            },
        },
    },
)

var mcQuestionGraphQL = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "MCQuestion",
        Fields: graphql.Fields{
            "question": &graphql.Field{
              Type: graphql.String,
            },
            "answers": &graphql.Field{
              Type: graphql.NewList(graphql.String),
            },
        },
    },
)

var oeQuestionGraphQL = graphql.NewObject(
    graphql.ObjectConfig{
        Name: "OEQuestion",
        Fields: graphql.Fields{
            "question": &graphql.Field{
              Type: graphql.String,
            },
        },
    },
)