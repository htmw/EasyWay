//  This file is used by Karma configuration file to recursively load all the .spec and framework files.

import 'zone.js/testing'; // Import the zone.js/testing for zone testing utilities
import { getTestBed } from '@angular/core/testing'; // Import getTestBed from angular core testing library
import {
BrowserDynamicTestingModule,
platformBrowserDynamicTesting
} from '@angular/platform-browser-dynamic/testing';

// First, initialize the Angular testing environment.
getTestBed().initTestEnvironment(
BrowserDynamicTestingModule,
platformBrowserDynamicTesting(),
); // Load all the .spec.ts files using the context object
