import { initFederation } from '@angular-architects/native-federation';

initFederation({'public-view': './remoteEntry.js'})
  .catch((err) => console.error(err))
  .then((_) => import('./bootstrap'))
  .catch((err) => console.error(err));
