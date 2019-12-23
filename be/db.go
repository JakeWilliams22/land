package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func query(db *sql.DB, query string) *sql.Rows {
    result, err := db.Query(query)
    if err != nil {
        panic(err)
    }
    return result
}

func initDb() *sql.DB {
    // Open up our database connection.
    //TODO: Password environment variable
    db, err := sql.Open("mysql", "user:pass@tcp(land.cyoywlf7jxiv.us-east-2.rds.amazonaws.com:3306)/Land")
    
    if err != nil {
        panic(err.Error())
    }
    
    return db
}

// Logic for parsing query results into object data structures should be here.
// Do not create structs for db rows/query results.
func parseLandingPageResult(result *sql.Rows) []LandingPage {
    landingPages := make([]LandingPage, 0)
    for result.Next() {
        var lp LandingPage
        
        err := result.Scan(&lp.Id, &lp.Title, &lp.SubTitle, &lp.BodyText, )
        if err != nil {
            panic(err.Error()) //TODO: Proper Error Handling
        }
        landingPages = append(landingPages, lp)
    }
    return landingPages
}

func parseJoinEmailListResult(result *sql.Rows) []JoinEmailList {
  joinEmailLists := make([]JoinEmailList, 0)
  for result.Next() {
    var jel JoinEmailList
    if err := result.Scan(&jel.JoinPrompt, &jel.JoinButtonText); err != nil {
      panic(err.Error())
    }
    joinEmailLists = append(joinEmailLists, jel)
  }
  return joinEmailLists
}

// func parseQuestionResult(result *sql.Rows) Questions {
  // TODO
  // return nil
// }

func testDb() {
    db := initDb()
    defer db.Close()
    
    result := query(db, "SELECT ID, TITLE, SUB_TITLE, BODY_TEXT FROM LANDING_PAGES")
    
    landingPages := parseLandingPageResult(result)
    
    for _, lp := range  landingPages {
      fmt.Printf(*lp.BodyText)
    }
}