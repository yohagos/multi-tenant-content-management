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
import { UnitService, UnitList } from 'shared-data-access';

@Component({
  selector: 'app-unit-table',
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
  templateUrl: './unit.table.html',
  styleUrl: './unit.table.scss',
  standalone: true,
})
export class UnitTable {
  private readonly unitService = inject(UnitService)

  readonly units = signal<UnitList>({
    data: [],
    total: 0,
    limit: 20,
    offset: 0,
  })

  readonly unitFilter = signal('')

  readonly filteredUnits = computed(() => {
    const value = this.unitFilter().toLowerCase()
    let list = this.units()

    if (list.data === null) {
      list.data = []
    }

    return {
      ...list,
      data: list.data.filter(u => u.title.toLowerCase().includes(value)),
    }
  })

  @ViewChild(MatPaginator) paginator!: MatPaginator
  @ViewChild(MatSort) sort!: MatSort

  isLoadingResults = signal(true)

  displayColumnsUnit = ['unit_number', 'title', 'floor', 'rooms', 'living_area_sqm', 'cold_rent', 'warm_rent', 'status', 'deposit']

  ngOnInit(): void {
    forkJoin({
      units: this.unitService.list(),
    }).subscribe(({units}) => {
      this.isLoadingResults.set(false)
      this.units.set(units)
    })
  }

  applyUnitFilter(event: Event) {
    this.unitFilter.set((event.target as HTMLInputElement).value)
  }
}
