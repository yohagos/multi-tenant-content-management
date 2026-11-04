import { Component } from '@angular/core';
import { BuildingTable } from "./building.table/building.table";
import { UnitTable } from "./unit.table/unit.table";

@Component({
  selector: 'app-content.overview',
  imports: [
    BuildingTable,
    UnitTable,
  ],
  templateUrl: './content.overview.html',
  styleUrl: './content.overview.scss',
})
export class ContentOverview {
}
