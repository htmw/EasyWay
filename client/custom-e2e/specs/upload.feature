Feature: Test Upload Page

  Scenario: Go to Get Login page
    Given I view the "http://localhost:4200/upload"
    When I click on button
    When I wait for 3000 ms
    Then the page title should be "EasyWay"
