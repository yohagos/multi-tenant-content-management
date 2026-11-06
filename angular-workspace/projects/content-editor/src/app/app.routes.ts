import { Routes } from '@angular/router';
import { ContentOverview } from './content.overview/content.overview';
import { AddBuildings } from './buildings/add.buildings/add.buildings';
import { EditBuilding } from './buildings/edit.building/edit.building';
import { AddUnit } from './units/add.unit/add.unit';
import { EditUnit } from './units/edit.unit/edit.unit';
import { ContentToolbar } from './content.overview/content.toolbar/content.toolbar';

export const routes: Routes = [
  {
    path: '',
    component: ContentToolbar,
    children: [
      {
        path: 'overview',
        component: ContentOverview,
      },
    ],
  },
  {
    path: 'building/add',
    component: AddBuildings,
  },
  {
    path: 'building/edit/:id',
    component: EditBuilding,
  },
  {
    path: 'unit/add',
    component: AddUnit,
  },
  {
    path: 'unit/edit',
    component: EditUnit,
  },
  {
    path: '*',
    redirectTo: '',
  },
];
