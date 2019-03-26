package test

import (
	"bracketbuilder/functions"
	"fmt"
	"testing"
)

func TestHTML(t *testing.T) {
	test, err := functions.GetHTML("this is not a url")
	fmt.Println(test, err)
	err == error("Get this%20is%20not%20a%20url: unsupported protocol scheme")
}
