Feature: Test Login Page

  Scenario: Go to Login page and verify title
    Given I view the "http://localhost:4200/login"
    When I click on button
    When I wait for 3000 ms
    Then the page title should be "EasyWay"
