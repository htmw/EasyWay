Feature: Test Home Page

  Scenario: Go to Get Home page
    Given I view the "http://localhost:4200/"
    When I click on button
    When I wait for 3000 ms
    Then the page title should be "EasyWay"
