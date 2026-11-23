import { Routes } from '@angular/router';
import { loadRemoteModule } from '@angular-architects/native-federation';
import { Home } from './home/home';
import { Auth } from './auth/auth';

export const routes: Routes = [
  {
    path: '',
    component: Home,
    children: [
      {
        path: 'authentication',
        component: Auth,
      },
      {
        path: 'register',
        redirectTo: 'auth',
        pathMatch: 'full',
      },
      {
        path: 'login',
        redirectTo: 'auth',
        pathMatch: 'full',
      },
    ],
  },
  {
    path: 'tenant-admin',
    loadChildren: () => loadRemoteModule('tenant-admin', './Routes').then((m) => m.routes),
  },
  {
    path: 'content-editor',
    loadChildren: () => loadRemoteModule('content-editor', './Routes').then((m) => m.routes),
  },
  {
    path: 'public-view',
    loadChildren: () => loadRemoteModule('public-view', './Routes').then((m) => m.routes),
  },
  /* {
    path: 'tenant-admin',
    loadChildren: () => loadRemoteModule({
      remoteEntry: 'http://localhost:4201/remoteEntry.js',
      remoteName: 'tenant-admin',
      exposedModule: './Routes',
    }).then(m => m.routes),
  },
  {
    path: 'content-editor',
    loadChildren: () => loadRemoteModule('content-editor', './Routes').then(m => m.routes),
  },
  {
    path: 'public-view',
    loadChildren: () => loadRemoteModule({
      remoteEntry: 'http://localhost:4203/remoteEntry.js',
      remoteName: 'public-view',
      exposedModule: './Routes',
    }).then(m => m.routes),
  }, */
];
