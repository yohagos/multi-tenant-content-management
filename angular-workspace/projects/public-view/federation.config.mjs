import { withNativeFederation, shareAll } from '@angular-architects/native-federation/config';
export default withNativeFederation({
  name: 'public-view',

  exposes: {
    './Routes': './projects/public-view/src/app/app.routes.ts',
  },

  shared: fromPackageJson({
    singleton: true,
    strictVersion: true,
    requiredVersion: 'auto',
  })
    .skip([
      'rxjs/ajax',
      'rxjs/fetch',
      'rxjs/testing',
      'rxjs/webSocket',
    ])
    .patch(
      ['@angular/*'],
      {
        includeSecondaries: {
          keepAll: true,
        },
      },
    ),

  features: {
    ignoreUnusedDeps: true,
    denseChunking: true,
    denseExternals: true,
  },
});
