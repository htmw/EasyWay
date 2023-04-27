//  This file is used by Karma configuration file to recursively load all the .spec and framework files.

import 'zone.js/testing'; // Import the zone.js/testing for zone testing utilities
import { getTestBed } from '@angular/core/testing'; // Import getTestBed from angular core testing library
import {
BrowserDynamicTestingModule,
platformBrowserDynamicTesting
} from '@angular/platform-browser-dynamic/testing'; // Import BrowserDynamicTestingModule and platformBrowserDynamicTesting for testing Angular components

declare const require: {
context(path: string, deep?: boolean, filter?: RegExp): {
<T>(id: string): T;
keys(): string[];
};
};

// First, initialize the Angular testing environment.
getTestBed().initTestEnvironment(
BrowserDynamicTestingModule,
platformBrowserDynamicTesting(),
); // Initializing the testing environment using BrowserDynamicTestingModule and platformBrowserDynamicTesting

// Then we find all the tests.
const context = require.context('./', true, /.spec.ts$/); // Get all the .spec.ts files in current directory recursively

// And load the modules.
context.keys().map(context); // Load all the .spec.ts files using the context object
