package main

import (
	"fmt"
	"log"
	"strconv"
)

func getLandingPage(id string) LandingPage {
	landingPage := getLandingPageBase(id)
	logoUri := toLogoUrl(*landingPage.LogoName)
	landingPage.LogoName = &logoUri
	if joinEmailList, exist := getJoinEmailList(id); exist {
		landingPage.JoinEmailList = &joinEmailList
	}
	if questions := getQuestions(id); len(questions) > 0 {
		landingPage.Questions = questions
	}
	return landingPage
}

func getLandingPageBase(id string) LandingPage {
	queryStr := "SELECT ID, TITLE, SUB_TITLE, BODY_TEXT, GOOGLE_ANALYTICS_ID, LOGO_NAME FROM LANDING_PAGES WHERE ID = " + id
	result := query(db, queryStr)
	landingPage := parseLandingPageResult(result)[0]
	return landingPage
}

func getJoinEmailList(landingPageId string) (JoinEmailList, bool) {
	result := query(db, "SELECT PROMPT, BUTTON_TEXT FROM JOIN_EMAIL_LIST WHERE LANDING_PAGE_ID = "+landingPageId)
	joinEmailListResults := parseJoinEmailListResult(result)
	if exist := len(joinEmailListResults) != 0; exist {
		return joinEmailListResults[0], true
	}
	var joinEmailList JoinEmailList
	return joinEmailList, false
}

func getQuestions(landingPageId string) []Question {
	questions := make([]Question, 0)
	mcQuestions := getMCQuestions(landingPageId)
	mcQuestionAddition := make([]Question, len(mcQuestions))
	for i, mcq := range mcQuestions {
		mcQuestionAddition[i] = mcq
	}
	questions = append(questions, mcQuestionAddition...)

	oeQuestions := getOEQuestions(landingPageId)
	oeQuestionAddition := make([]Question, len(oeQuestions))
	for i, oeq := range oeQuestions {
		oeQuestionAddition[i] = oeq
	}
	questions = append(questions, oeQuestionAddition...)
	return questions
}

func getMCQuestions(landingPageId string) []MCQuestion {
	questionResult := query(db, "SELECT ID, QUESTION FROM MULTIPLE_CHOICE_QUESTIONS WHERE LANDING_PAGE_ID = "+landingPageId)
	questionIds, questions := parseMCQuestionResult(questionResult)
	for index, _ := range questions {
		questionId := questionIds[index]
		answerQuery := "SELECT ANSWER_TEXT FROM MULTIPLE_CHOICE_ANSWERS WHERE QUESTION_ID = " + strconv.Itoa(questionId)
		answerResult := query(db, answerQuery)
		answers := parseMCAnswersResult(answerResult)
		questions[index].Answers = answers
	}
	return questions
}

func getOEQuestions(landingPageId string) []OpenEndedQuestion {
	questionResult := query(db, "SELECT ID, QUESTION FROM OPEN_ENDED_QUESTIONS WHERE LANDING_PAGE_ID = "+landingPageId)
	_, questions := parseOEQuestionResult(questionResult)
	return questions
}

func addEmailSubscriber(email string, landingPageId string) bool {
	if valid := validateEmail(email); valid {
		sql := fmt.Sprintf(
			"INSERT INTO EMAILS VALUES ('%s', '%s', NOW())",
			email,
			landingPageId)
		return insertQuery(db, sql)
	} else {
		log.Printf("Email '%s' is not valid", email)
	}
	return false
}
