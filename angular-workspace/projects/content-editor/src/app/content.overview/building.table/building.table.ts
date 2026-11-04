import { Component, computed, inject, signal, ViewChild } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatTableModule } from '@angular/material/table';
import { forkJoin } from 'rxjs';
import { BuildingList, BuildingService } from 'shared-data-access';

@Component({
  selector: 'app-building-table',
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
  templateUrl: './building.table.html',
  styleUrl: './building.table.scss',
  standalone: true,
})
export class BuildingTable {
  private readonly buildingService = inject(BuildingService)

  @ViewChild(MatPaginator) paginator!: MatPaginator
  @ViewChild(MatSort) sort!: MatSort

  readonly buildings = signal<BuildingList>({
    data: [],
    total: 0,
    limit: 20,
    offset: 0,
  })

  readonly buildingFilter = signal('')

  readonly filteredBuildings = computed(() => {
    const value = this.buildingFilter().toLowerCase()
    const list = this.buildings()

    if (list.data === null) {
      list.data = []
    }

    return {
      ...list,
      data: list.data.filter(b => b.name.toLowerCase().includes(value)),
    }
  })

  isLoadingResults = signal(true)

  displayColumnsBuilding = ['name', 'year_built', 'total_floors', 'total_units', 'status', 'published']

    ngOnInit(): void {
      forkJoin({
        buildings: this.buildingService.list(),
      }).subscribe(({buildings}) => {
        this.isLoadingResults.set(false)
        this.buildings.set(buildings)
      })
    }

    applyBuildingFilter(event: Event) {
      this.buildingFilter.set((event.target as HTMLInputElement).value)
    }
}
