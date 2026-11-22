import { ApplicationConfig, importProvidersFrom, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideRouter } from '@angular/router';

import { routes } from './app.routes';
import { environment } from "./environments/environment";
import { API_BASE_URL, authInterceptor } from "shared-data-access";
import { MatNativeDateModule } from '@angular/material/core';
import { provideHttpClient, withInterceptors } from '@angular/common/http';

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideRouter(routes),
    importProvidersFrom([
      MatNativeDateModule,
    ]),
    provideHttpClient(withInterceptors([authInterceptor])),
  { provide: API_BASE_URL, useValue: environment.apiBaseUrl },
  ],
};
