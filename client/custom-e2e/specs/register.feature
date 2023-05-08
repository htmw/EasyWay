Feature: Test Register Page

  Scenario: Go to Get Register page
    Given I view the "http://localhost:4200/register"
    When I click on button
    When I wait for 3000 ms
    Then the page title should be "EasyWay"
