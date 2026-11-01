import { Component, inject, OnInit, signal, ViewChild } from '@angular/core';
import { MatButtonModule } from "@angular/material/button";
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatPaginator, MatPaginatorModule } from "@angular/material/paginator";
import { MatProgressSpinnerModule } from "@angular/material/progress-spinner";
import { MatSort, MatSortModule } from "@angular/material/sort";
import { MatTableModule } from "@angular/material/table";
import { forkJoin } from 'rxjs';
import { BuildingList, BuildingService, ContentService, UnitList, UnitService } from "shared-data-access";

@Component({
  selector: 'app-content.overview',
  imports: [
    MatButtonModule,
    MatCardModule,
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
    MatPaginatorModule,
    MatProgressSpinnerModule,
    MatSortModule,
    MatTableModule,
  ],
  templateUrl: './content.overview.html',
  styleUrl: './content.overview.scss',
})
export class ContentOverview implements OnInit {
  private readonly contentService = inject(ContentService)
  private readonly buildingService = inject(BuildingService)
  private readonly unitService = inject(UnitService)

  readonly buildings = signal<BuildingList>({
    data: [],
    total: 0,
    limit: 20,
    offset: 0,
  })
  readonly units = signal<UnitList>({
    data: [],
    total: 0,
    limit: 20,
    offset: 0,
  })

  @ViewChild(MatPaginator) paginator!: MatPaginator
  @ViewChild(MatSort) sort!: MatSort

  isLoadingResults = true

  displayColumnsBuilding = ['name', 'year_built', 'total_floors', 'total_units', 'status', 'published']
  displayColumnsUnit = ['unit_number', 'title', 'floor', 'rooms', 'living_area_sqm', 'cold_rent', 'warm_rent', 'status', 'deposit']

  ngOnInit(): void {
    forkJoin({
      buildings: this.buildingService.list(),
      units: this.unitService.list(),
    }).subscribe(({buildings, units}) => {
      this.isLoadingResults = false
      this.buildings.set(buildings)
      this.units.set(units)
    })
  }

  applyBuildingFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value

    this.buildings()?.data.filter(a => filterValue !== null ? a.name.toLowerCase().includes(filterValue) : '')

    /* this.buildings.update(
      buildings => ({
        ...buildings,
        data: buildings.data.filter(building => building.name.toLowerCase().includes(filterValue))
      })
    ) */
  }

  applyUnitFilter(value: string) {
    const filterValue = value.toLowerCase()

    this.units.update(
      units => ({
        ...units,
        data: units.data.filter(unit => unit.title.toLowerCase().includes(filterValue))
      })
    )
  }
}
