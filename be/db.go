package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// Logic for parsing query results into object data structures should live here.
// AKA do not create structs for db rows/query results.
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


func testDb() {
// func main() {
    fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    //TODO: Password environment variable
    db, err := sql.Open("mysql", "jisawesome:password!@tcp(land.cyoywlf7jxiv.us-east-2.rds.amazonaws.com:3306)/Land")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()
    
    result, err := db.Query(
      "SELECT ID, TITLE, SUB_TITLE, BODY_TEXT FROM LANDING_PAGES")
    
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    
    landingPages := parseLandingPageResult(result)
    
    for _, lp := range  landingPages {
      fmt.Printf(*lp.BodyText)
    }
}