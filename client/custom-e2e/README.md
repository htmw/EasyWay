# Cucumber

Cucumber is a tool that supports BDD.

## What is BDD?

Behavior Driven Development (BDD) is an extension to the concept of TDD, but instead of testing your code you are testing your product, and specifically that your product behaves as you desire it to.

## What is Gherkin?

Gherkin allows for test scripts to be written in a human readable format, which can then be shared between all of the stakeholders in the product development.

It gives you the ability to remove logic details from behavior tests. Gherkin serves two purposes: serving as your product documentation and automated tests.

Gherkin files typically have a `.feature` file extension that contain tests, written in the Gherkin language.

```feature
Given I navigate to the "nazmul website" page
When I click my profile link
And I search for "Mac"
Then I see title "Mac | Search Results | Nazmul Website"
```

Straight away we can see that this test tells us **what** to do and not **how** to do it. It is written in a language that makes sense to anyone reading it, and — importantly — that will most likely be correct no matter how the end product might be tweaked. The product UI could be changed completely, but as long as the functionality is equivalent then the Gherkin is still accurate.

## What is Cucumber?

Gherkin is a simple syntax for natural language tests, and Cucumber is a testing framework for behavior driven development that can execute them.

## What is the advantages of Cucumber?

- It is helpful to involve business stakeholders who can't easily read code.
- Cucumber Testing focuses on end-user experience.
- Style of writing tests allow for easier reuse of code in the tests.
- Quick and easy set up and execution.
- Efficient tool for testing.

## Installation:

As we are using JavaScript so please install <a href="http://nodejs.org">Node.js</a>.

* Clone the repository

```
git clone https://github.com/htmw/EasyWay.git
```
* Install NodeJS LTS version from https://nodejs.org/en/ for your Operating System.
* Navigate to client folder and install required libraries:
```
cd ./client/
```

* Install Cucumber.js:

Add `cucumber` as a development dependency:

```
npm install cucumber --save-dev
```

* Install Chai assertion library:

```
npm install chai --save-dev
```

* Install Selenium webdriver library:

```
npm install selenium-webdriver
```

* Running Selenium webdriver library:

```
webdriver-manager start
```

* Running Cucumber:

```
protractor custom-e2e/protractor.conf.js
```
