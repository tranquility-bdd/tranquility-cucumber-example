package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/tranquility-bdd/tranquility"
	"strings"
)

var paramA, paramB string
var resp *tranquility.Response
var err error

func iSetParamAparamBAsAnd(arg1, arg2 string) error {
	tranquility.Env.Set("paramA", arg1)
	tranquility.Env.Set("paramB", arg2)
	return nil
}

func iMakeAnHttpRequestToGetPostmanecho() error {
	var url = "https://postman-echo.com/get?foo1={{.paramA}}+&foo2={{.paramB}}"
	var action = tranquility.Action{Method: "GET", URL: url}
	resp, err = action.Run()
	return nil
}

func thereShouldBeEvidenceOfASuccessfulHttpRequest() error {
	if resp.StatusCode != 200 {
		return fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}
	var str = "https://postman-echo.com/get?foo1=foo+&foo2=bar"
	if !strings.Contains(resp.Body, str) {
		return fmt.Errorf("Response body: %v doesn't contain sting: %v", resp.Body, str)
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
