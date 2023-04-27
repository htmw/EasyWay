// Import the necessary Angular libraries and modules
import { enableProdMode } from '@angular/core';
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

// Import the main AppModule and the environment configuration
import { AppModule } from './app/app.module';
import { environment } from './environments/environment';

// If we are in production mode, enable Angular's production mode
if (environment.production) {
  enableProdMode();
}

// Bootstrap the AppModule and catch any errors
platformBrowserDynamic().bootstrapModule(AppModule)
  .catch(err => console.error(err));
