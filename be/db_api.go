package main

import (
    "fmt"
)

func getLandingPage(id string) LandingPage {
  landingPage := getLandingPageBase(id)
  joinEmailList := getJoinEmailList(id)
  
  landingPage.JoinEmailList = &joinEmailList
  fmt.Printf("%+v",landingPage)
  return landingPage
}

func getLandingPageBase(id string) LandingPage {
  var landingPage LandingPage
  result := query(db, "SELECT ID, TITLE, SUB_TITLE, BODY_TEXT FROM LANDING_PAGES WHERE ID = " + id)
  landingPage = parseLandingPageResult(result)[0]
  return landingPage
}

func getJoinEmailList(landingPageId string) JoinEmailList {
  var joinEmailList JoinEmailList
  result := query(db, "SELECT PROMPT, BUTTON_TEXT FROM JOIN_EMAIL_LIST WHERE LANDING_PAGE_ID = " + landingPageId)
  joinEmailList = parseJoinEmailListResult(result)[0]
  return joinEmailList
}

// func getQuestionsForLandingPage(landingPageId string) []Question {
  // TODO
  
// }