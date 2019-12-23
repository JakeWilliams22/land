package main

import (
	"fmt"
)

func getLandingPage(id string) LandingPage {
	landingPage := getLandingPageBase(id)
	if joinEmailList, exist := getJoinEmailList(id); exist {
		landingPage.JoinEmailList = &joinEmailList
	}
	fmt.Printf("%+v", landingPage)
	return landingPage
}

func getLandingPageBase(id string) LandingPage {
	var landingPage LandingPage
	result := query(db, "SELECT ID, TITLE, SUB_TITLE, BODY_TEXT FROM LANDING_PAGES WHERE ID = "+id)
	landingPage = parseLandingPageResult(result)[0]
	return landingPage
}

func getJoinEmailList(landingPageId string) (JoinEmailList, bool) {
	result := query(db, "SELECT PROMPT, BUTTON_TEXT FROM JOIN_EMAIL_LIST WHERE LANDING_PAGE_ID = "+landingPageId)
	joinEmailListResults := parseJoinEmailListResult(result)
	return joinEmailListResults[0], len(joinEmailListResults) != 0
}
