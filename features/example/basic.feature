Feature: Example http request
  In order to be happy
  As a gherkin scenario
  (That does not replace in any way meaningful testing)
  I will check that it's possible to make an HTTP request

  Scenario: Get of postmanecho.com
    Given I set paramA/paramB as "foo" and "bar"
    When I make an http request to get postmanecho
    Then there should be evidence of a successful http request
