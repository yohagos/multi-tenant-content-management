import { withNativeFederation, fromPackageJson } from '@angular-architects/native-federation/config';
export default withNativeFederation({
  name: 'tenant-admin',

  exposes: {
    './Routes': './projects/tenant-admin/src/app/app.routes.ts',
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
