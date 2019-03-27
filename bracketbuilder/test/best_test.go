package test

import (
	"bracketbuilder/functions"
	"fmt"
	"testing"
)

// Test that an error is returned when a bad url is given
func TestBadHTML(t *testing.T) {
	_, err := functions.GetHTML("this is not a url")
	if err != nil {
		fmt.Println("properly detected invalid url", err)
	}
}

// Test that checks if an error is returned when a good url is given
func TestGoodHtml(t *testing.T) {
	_, err := functions.GetHTML("https://www.example.com")
	if err == nil {

	} else {
		t.Error("Unexpected Error, check connectivity and try again")
	}
}
