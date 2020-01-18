package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

func getLandingPageById(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"]
	if id, ok := id.(string); ok {
		return getLandingPage(id), nil
	}
	panic("Invalid type for ID field in query")
	// return landingPages()[1], nil
}

func resolveAddEmailSubscriber(params graphql.ResolveParams) (interface{}, error) {
	email := params.Args["email"].(string)
	landingPageId := params.Args["landingPageId"].(string)
	success := addEmailSubscriber(email, landingPageId)
	return success, nil
}

func landingPages() []LandingPage {
	mc := MCQuestion{
		Question: "How are you",
		Answers:  []string{"good"},
	}

	q1List := []Question{mc}

	oe := OpenEndedQuestion{
		Question: "Why",
	}

	q2List := []Question{oe}

	t1 := "Hummingnerd"
	t2 := "Durgs"

	st1 := "Sing like your favorite fartists"
	st2 := "Do some durgs"

	bt1 := "You found us while Hummingnerd is still in the Lab! " +
		"We'll let you know as soon as it's ready!"
	bt2 := "2You found us while Hummingbird is still in the Lab! " +
		"We'll let you know as soon as it's ready!"

	jel1 := JoinEmailList{
		JoinPrompt:     "Keep up to d8",
		JoinButtonText: "Joins",
	}
	jel2 := JoinEmailList{
		JoinPrompt:     "Find out about durgs",
		JoinButtonText: "Segrada Familia",
	}

	return []LandingPage{
		{
			Title:         &t1,
			SubTitle:      &st1,
			BodyText:      &bt1,
			JoinEmailList: &jel1,
			Questions:     q1List,
		},
		{
			Title:         &t2,
			SubTitle:      &st2,
			BodyText:      &bt2,
			JoinEmailList: &jel2,
			Questions:     q2List,
		},
	}
}

func createSchema() graphql.Schema {
	rootQuery, rootMutation := getGQLConfigs()
	schemaConfig :=
		graphql.SchemaConfig{
			Query:    graphql.NewObject(rootQuery),
			Mutation: graphql.NewObject(rootMutation)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	schema.AppendType(oeQuestionGraphQL)
	schema.AppendType(mcQuestionGraphQL)
	return schema
}

func testAPI() {
	schema := createSchema()

	// Query
	query := `
        {
          landingPage(id: "-2") {
            id
            title
            subTitle
            bodyText
            googleAnalyticsId
            questions {
                questionPrompt
                questionType
                ... on MCQuestion {
                    answers
                }
            }
            joinEmailList {
                joinPrompt
                joinButtonText
            }
          }
        }
    `

	mutation := `
		mutation {
		  addEmailSubscriber(email: "test email", landingPageId: "-1")
		}
	`

	// sssassaxcxdsasuv wefrddssu - Milly the dog
	fmt.Print(query != "")
	params := graphql.Params{Schema: schema, RequestString: mutation}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf(
			"failed to execute graphql operation, errors: %+v",
			r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("\n%s \n", rJSON) // {“data”:{“hello”:”world”}}
}
