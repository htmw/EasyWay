Feature: Test Service Page

  Scenario: Go to Get Service page
    Given I view the "http://localhost:4200/services"
    When I click on button
    When I wait for 3000 ms
    Then the page title should be "EasyWay"
