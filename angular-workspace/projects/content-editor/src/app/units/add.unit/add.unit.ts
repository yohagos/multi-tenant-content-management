import { AfterViewInit, Component, inject, OnInit, signal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatStepperModule } from '@angular/material/stepper';
import { Router } from '@angular/router';
import { Building, BuildingFilter, BuildingList, BuildingService, Unit, UnitService, UnitStatus } from 'shared-data-access';
import { SnackbarService, BuildingSelect, BuildingTransferService } from 'shared-ui';

@Component({
  selector: 'app-add.unit',
  imports: [
    MatButtonModule,
    MatCardModule,
    MatCheckboxModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatStepperModule,

    ReactiveFormsModule,
    BuildingSelect,
  ],
  standalone: true,
  templateUrl: './add.unit.html',
  styleUrl: './add.unit.scss',
})
export class AddUnit implements OnInit, AfterViewInit {
  private readonly _fb = inject(FormBuilder)
  private readonly _router = inject(Router)
  private readonly _snackBarService = inject(SnackbarService)
  private readonly _unitService = inject(UnitService)
  private readonly _buildingService = inject(BuildingService)
  private readonly _buildingTransferService = inject(BuildingTransferService)

  chosenBuilding: Partial<Building> | null = null

  unitTitle = signal<string>('')

  readonly buildings = signal<BuildingList>({
    data: [],
    total: 0,
    limit: 20,
    offset: 0,
  })

  buildingFilter: BuildingFilter = {
    limit: 20,
    offset: 0,
  }

  searchForBuilding = ''

  readonly unitStatus = signal<UnitStatus[]>([
    'draft',
    'available',
    'reserved',
    'rented',
    'archived',
  ])

  readonly buildingForm = this._fb.group({
    building_id: ['', Validators.required]
  })

  readonly firstForm = this._fb.group({
    unit_number: ['', Validators.required],
    /* title: [''], */
    description: [''],
    floor: ['', Validators.required],
    rooms: ['', Validators.required],
    bedrooms: ['', Validators.required],
    bathrooms: ['', Validators.required],
  })

  readonly secondForm = this._fb.group({
    living_area_sqm: ['', Validators.required],
    total_area_sqm: ['', Validators.required],
    balcony_area_sqm: ['', Validators.required],
    ceiling_height_m: ['', Validators.required],
  })

  readonly thirdForm = this._fb.group({
    cold_rent: ['', Validators.required],
    warm_rent: ['', Validators.required],
    deposit: ['', Validators.required],
    available_from: ['', Validators.required],
    status: ['', Validators.required],
    published: ['', Validators.required],
    published_at: [''],
  })

  ngOnInit() {
    this._buildingService.list().subscribe({
      next: (data) => {
        this._buildingTransferService.setBuildings(data.data)
      }
    })
  }

  ngAfterViewInit(): void {
    this._buildingTransferService.selectedBuilding$.subscribe({
      next: (data) => {
        if (data !== null && data.id !== undefined) {
          this.chosenBuilding = data
          this.unitTitle.set(data.name + '-')
        }
      }
    })

    this.firstForm.controls.unit_number.valueChanges.subscribe((value) => {
      console.log(value)
      if (value !== null) {
        this.unitTitle().concat(value)
      }
    })
  }

  save() {
    if (!this.firstForm.valid || !this.secondForm.valid || !this.thirdForm.valid) {
      this._snackBarService.errorSnackBar('Please check if required fields are filled.')
      return
    }

    let unit = this.transformData()

    this._unitService.create(unit).subscribe({
      next: () => {
        this._snackBarService.successSnackBar(`Unit ${unit.unit_number} added successfully.`)
      },
      error: (err) => {
        this._snackBarService.errorSnackBar(`Error occurred: ${err}`)
      },
      complete: () => {
        this.backToOverview()
      },
    })
  }

  private transformData(): Unit {
    return {
      ...this.firstForm.getRawValue(),
      ...this.secondForm.getRawValue(),
      ...this.thirdForm.getRawValue(),
      building_id: this.chosenBuilding,
      floor: Number(this.firstForm.getRawValue().floor),
      rooms: Number(this.firstForm.getRawValue().rooms),
      bathrooms: Number(this.firstForm.getRawValue().bathrooms),
      bedrooms: Number(this.firstForm.getRawValue().bedrooms),

      living_area_sqm: Number(this.secondForm.getRawValue().living_area_sqm),
      total_area_sqm: Number(this.secondForm.getRawValue().total_area_sqm),
      ceiling_height_m: Number(this.secondForm.getRawValue().ceiling_height_m),
      balcony_area_sqm: Number(this.secondForm.getRawValue().balcony_area_sqm),

      cold_rent: Number(this.thirdForm.getRawValue().cold_rent),
      warm_rent: Number(this.thirdForm.getRawValue().warm_rent),
      deposit: Number(this.thirdForm.getRawValue().deposit),
    } as unknown as Unit
  }

  cancel() {
    this.firstForm.reset()
    this.secondForm.reset()
    this.thirdForm.reset()
    this.backToOverview()
  }

  private backToOverview() {
    this._router.navigate(['content-editor'])
  }
}
