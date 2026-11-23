import { ApplicationConfig, importProvidersFrom, provideBrowserGlobalErrorListeners } from '@angular/core';
import { provideRouter } from '@angular/router';
import { MatNativeDateModule, provideNativeDateAdapter } from "@angular/material/core";
import { MatDateFnsModule } from "@angular/material-date-fns-adapter";

import { routes } from './app.routes';

export const appConfig: ApplicationConfig = {
  providers: [
    provideBrowserGlobalErrorListeners(),
    provideRouter(routes),
    provideNativeDateAdapter(),
    importProvidersFrom([MatNativeDateModule, MatDateFnsModule,]),
  ],
};
