package main

import (
	"fmt"
	"regexp"
)

const LogoBucketBase = "https://land-logos.s3.us-east-2.amazonaws.com/"

func validateEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

func toLogoUrl(name string) string {
	fmt.Print(LogoBucketBase + name)
	return LogoBucketBase + name
}
