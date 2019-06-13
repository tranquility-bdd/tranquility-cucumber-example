package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/tranquility-bdd/tranquility"
)

var paramA, paramB string
var resp *tranquility.Response
var err error

func iSetParamAparamBAsAnd(arg1, arg2 string) error {
	paramA = arg1
	paramB = arg2
	return nil
}

func iMakeAnHttpRequestToGetPostmanecho() error {
	var url = "https://postman-echo.com/get?foo1=" + paramA + "&foo2=" + paramB
	var action = tranquility.Action{"GET", url, nil, nil, ""}
	resp, err = action.Run()
	return nil
}

func thereShouldBeEvidenceOfASuccessfulHttpRequest() error {
	if resp.StatusCode != 200 {
		return fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I set paramA\/paramB as "([^"]*)" and "([^"]*)"$`, iSetParamAparamBAsAnd)
	s.Step(`^I make an http request to get postmanecho$`, iMakeAnHttpRequestToGetPostmanecho)
	s.Step(`^there should be evidence of a successful http request$`, thereShouldBeEvidenceOfASuccessfulHttpRequest)
	s.BeforeScenario(func(interface{}) {
		paramA, paramB = "", "" // clean the state before every scenario
	})
}
