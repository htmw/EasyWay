exports.config = {
  seleniumAddress: 'http://localhost:4444/wd/hub',
  specs: ['./specs/*.feature'],
  capabilities: {
    browserName: 'chrome',
    'safari.options': {
      technologyPreview: true
    }
  },
  framework: 'custom',
  frameworkPath: require.resolve('protractor-cucumber-framework'),
  cucumberOpts: {
    require: './steps/*.js',
  },
};
