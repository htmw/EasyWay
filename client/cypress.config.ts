import { defineConfig } from "cypress";

export default defineConfig({
  projectId: 'gvmhix',
  videosFolder: "cypress/videos",
  screenshotsFolder: "cypress/screenshots",
  fixturesFolder: "cypress/fixtures",
  videoCompression: false,

  e2e: {
    setupNodeEvents(on, config) {
      return require("./cypress/plugins/index.ts")(on, config);
    },
    baseUrl: "http://localhost:4200",
  },

  component: {
    devServer: {
      framework: "angular",
      bundler: "webpack",
    },
    specPattern: "**/*.cy.ts",
    setupNodeEvents(on, config) {
      // You can add any additional setup logic for component tests here
    },
  },
});
