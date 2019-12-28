package main

import (
	"strconv"
)

func getLandingPage(id string) LandingPage {
	landingPage := getLandingPageBase(id)
	if joinEmailList, exist := getJoinEmailList(id); exist {
		landingPage.JoinEmailList = &joinEmailList
	}
	if questions := getQuestions(id); len(questions) > 0 {
		landingPage.Questions = questions
	}
	return landingPage
}

func getLandingPageBase(id string) LandingPage {
	result := query(db, "SELECT ID, TITLE, SUB_TITLE, BODY_TEXT FROM LANDING_PAGES WHERE ID = "+id)
	landingPage := parseLandingPageResult(result)[0]
	return landingPage
}

func getJoinEmailList(landingPageId string) (JoinEmailList, bool) {
	result := query(db, "SELECT PROMPT, BUTTON_TEXT FROM JOIN_EMAIL_LIST WHERE LANDING_PAGE_ID = "+landingPageId)
	joinEmailListResults := parseJoinEmailListResult(result)
	return joinEmailListResults[0], len(joinEmailListResults) != 0
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
