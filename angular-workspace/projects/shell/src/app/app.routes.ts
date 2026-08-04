import { Routes } from '@angular/router';
import { loadRemoteModule } from '@angular-architects/native-federation';
import { Nav } from './nav/nav';
import { Home } from './home/home';

export const routes: Routes = [
  {
    path: '',
    component: Home,
  },
  {
    path: 'tenant-admin',
    loadChildren: () => loadRemoteModule({
      remoteEntry: 'http://localhost:4201/remoteEntry.js',
      remoteName: 'tenant-admin',
      exposedModule: './Routes',
    }).then(m => m.routes),
  },
  {
    path: 'content-editor',
    loadChildren: () => loadRemoteModule({
      remoteEntry: 'http://localhost:4202/remoteEntry.js',
      remoteName: 'content-editor',
      exposedModule: './Routes',
    }).then(m => m.routes),
  },
  {
    path: 'public-view',
    loadChildren: () => loadRemoteModule({
      remoteEntry: 'http://localhost:4203/remoteEntry.js',
      remoteName: 'public-view',
      exposedModule: './Routes',
    }).then(m => m.routes),
  },
];
