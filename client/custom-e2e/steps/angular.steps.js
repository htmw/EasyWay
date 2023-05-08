'use strict';

var { Given } = require('cucumber');
var { When } = require('cucumber');
var { Then } = require('cucumber');

// Use the external Chai As Promised to deal with resolving promises in
// expectations
const chai = require('chai');
const chaiAsPromised = require('chai-as-promised');
chai.use(chaiAsPromised);
const expect = chai.expect;

Given(/^I view the "([^"]*)"$/, function (url, callback) {
    browser.get(url).then(function () {
        callback();
    });
});

When(/^I click on button$/, function (callback) {
    callback();
});

When(/^I wait for (\d+) ms$/, function (timeToWait, callback) {
    setTimeout(callback, timeToWait);
});

Then(/the page title should be "([^"]*)"$/, function (text, callback) {
    expect(browser.getTitle()).to.eventually.equal(text).and.notify(callback);
});
