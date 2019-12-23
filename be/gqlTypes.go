package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

var (
	questionsInterfaceGQL *graphql.Interface
	oeQuestionGraphQL     *graphql.Object
	mcQuestionGraphQL     *graphql.Object
)

func getRootGQLQuery() graphql.ObjectConfig {
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

	questionsInterfaceGQL = graphql.NewInterface(
		graphql.InterfaceConfig{
			Name: "Question",
			Fields: graphql.Fields{
				"questionPrompt": &graphql.Field{
					Type: graphql.String,
				},
				"questionType": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	mcQuestionGraphQL = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "MCQuestion",
			Interfaces: []*graphql.Interface{
				questionsInterfaceGQL,
			},
			Fields: graphql.Fields{
				"questionPrompt": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if q, ok := p.Source.(MCQuestion); ok {
							return q.QuestionText(), nil
						}
						return "", nil
					},
				},
				"questionType": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if q, ok := p.Source.(MCQuestion); ok {
							return q.QuestionType(), nil
						}
						return "", nil
					},
				},
				"answers": &graphql.Field{
					Type: graphql.NewList(graphql.String),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if q, ok := p.Source.(MCQuestion); ok {
							return q.Answers, nil
						}
						return make([]string, 0), nil
					},
				},
			},
		},
	)

	oeQuestionGraphQL = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "OEQuestion",
			Interfaces: []*graphql.Interface{
				questionsInterfaceGQL,
			},
			Fields: graphql.Fields{
				"questionPrompt": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if q, ok := p.Source.(OpenEndedQuestion); ok {
							return q.QuestionText(), nil
						}
						return "", nil
					},
				},
				"questionType": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if q, ok := p.Source.(OpenEndedQuestion); ok {
							return q.QuestionType(), nil
						}
						return "", nil
					},
				},
			},
		},
	)

	questionsInterfaceGQL.ResolveType =
		func(p graphql.ResolveTypeParams) *graphql.Object {
			fmt.Printf("\n%+v\n", p)
			_, ok2 := p.Value.(Question)
			fmt.Printf("\n%s\n", ok2)
			if question, ok := p.Value.(Question); ok {
				fmt.Printf("\n%+v\n", question)
				if question.QuestionType() == 0 {
					return mcQuestionGraphQL
				} else if question.QuestionType() == 1 {
					return oeQuestionGraphQL
				}
			}
			return oeQuestionGraphQL
		}

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
				"bodyText": &graphql.Field{
					Type: graphql.String,
				},
				"joinEmailList": &graphql.Field{
					Type: joinEmailListGQL,
				},
				"questions": &graphql.Field{
					Type: graphql.NewList(questionsInterfaceGQL),
				},
			},
		},
	)

	var rootFields = graphql.Fields{
		"allLandingPages": &graphql.Field{
			Type: graphql.NewList(landingPageGQL),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return landingPages(), nil
			},
		},
		"landingPage": &graphql.Field{
			Type: landingPageGQL,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: getLandingPageById,
		},
	}

	return graphql.ObjectConfig{Name: "RootQuery", Fields: rootFields}
}
