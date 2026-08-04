import { initFederation } from '@angular-architects/native-federation';

initFederation({'tenant-admin': './remoteEntry.json'})
  .catch((err) => console.error(err))
  .then((_) => import('./bootstrap'))
  .catch((err) => console.error(err));
