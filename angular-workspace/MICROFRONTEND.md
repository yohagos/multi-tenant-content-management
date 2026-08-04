# Angular Micro Frontend

Created this MICROFRONTEND.md for future reference, if there might be a need. I took me a couple of days to get it to run, because of issues with the newer angular version 22.

## Documentation

I had to use the [Documentation](https://www.npmjs.com/package/@angular-architects/native-federation-v4/v/21.2.7) to get it to work. Even though I also used Deepseek, ChatGPT and Grok, they had their issues to comprehend the problems/errors.

## Step by Step

- Creating a new angular project without an application

```sh
ng new angular-workspace --create-application=false --standalone --style=scss
```

- Installing native-federation

```sh
npm i -D @angular-architects/native-federation
```

- Creating shell and remote's

```sh
ng generate application shell --standalone --style=scss
ng generate application tenant-admin --standalone --style=scss
ng generate application content-editor --standalone --style=scss
ng generate application public-view --standalone --style=scss
```

- Initializing shell & remote's

```sh
ng add @angular-architects/native-federation --project shell --port 4200 --type dynamic-host
ng add @angular-architects/native-federation --project tenant-admin --port 4201 --type remote
ng add @angular-architects/native-federation --project content-editor --port 4202 --type remote
ng add @angular-architects/native-federation --project public-view --port 4203 --type remote
```

- Checking and adjusting /shell/public/federation.config.json

The Key/Value have to be 

```json
{
  "tenant-admin": "http://localhost:4201/remoteEntry.json",
  "content-editor": "http://localhost:4202/remoteEntry.json",
  "public-view": "http://localhost:4203/remoteEntry.json"
}
```

- Checking and adjusting federation.config.mjs for each remotes

```ts
export default withNativeFederation({
  // has to be the same as corresponding Key in federation.config.json
  name: 'tenant-admin', 

  exposes: {
    // exposing the right key and class
    // init values are app component and Module as key
    './Routes': './projects/tenant-admin/src/app/app.routes.ts',
  },

  ..
})
```

- Adjusting shell routes with loadRemoteModules

```ts
import { loadRemoteModule } from '@angular-architects/native-federation';

export const routes: Routes = [
  {
    path: '',
    component: Home,
  },
  {
    path: 'tenant-admin',
    loadChildren: () => loadRemoteModule({
      // Value from /shell/public/federation.config.json
      remoteEntry: 'http://localhost:4201/remoteEntry.js', 
      // Key from /shell/public/federation.config.json
      remoteName: 'tenant-admin',
      exposedModule: './Routes', // Exposed Key
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
```

- (manuelly) Adjusting all selectors of shell, remotes and child components

Even though remotes are generate with --type remote, there can be an problem, in which they will be generated as new application and will each be named 'app-root'. If the shell and remotes each have 'app-root' as selector, it will cause an error (name conflict).

- Update main.ts of each remote

```ts
import { initFederation } from '@angular-architects/native-federation';

// Adding remoteEntry path for each remote with their project/key
initFederation({ 'tenant-admin': './remoteEntry.json' })
  .catch((err) => console.error(err))
  .then((_) => import('./bootstrap'))
  .catch((err) => console.error(err));
```

- Generating shared components

```sh
# Used for better structuring the project and make all functionalities accessible for shell and remotes
ng g library shared-ui --standalone
ng g library shared-data-access --standalone
ng g library shared-utils --standalone
```
