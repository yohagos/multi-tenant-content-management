import { Component, ChangeDetectionStrategy, input } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatTabsModule } from "@angular/material/tabs";
import { Building } from 'shared-data-access';

@Component({
  selector: 'app-edit-building',
  imports: [
    MatCardModule,
    MatButtonModule,
    MatTabsModule,
  ],
  templateUrl: './edit.building.html',
  changeDetection: ChangeDetectionStrategy.Eager,
  styleUrl: './edit.building.scss',
  standalone: true,
})
export class EditBuilding {
  building = input.required<Building>()
}
