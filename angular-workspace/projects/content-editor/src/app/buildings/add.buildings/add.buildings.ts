import { STEPPER_GLOBAL_OPTIONS } from '@angular/cdk/stepper';
import { Component, inject, signal } from '@angular/core';
import { FormBuilder, FormControl, ReactiveFormsModule, Validators } from "@angular/forms";
import { MatButtonModule } from "@angular/material/button";
import { MatCardModule } from "@angular/material/card";
import { MatCheckboxModule } from "@angular/material/checkbox";
import { MatFormFieldModule } from "@angular/material/form-field";
import { MatInputModule } from "@angular/material/input";
import { MatSelectModule } from "@angular/material/select";
import { MatStepperModule } from "@angular/material/stepper";
import { Router } from '@angular/router';
import { Building, BuildingService, BuildingStatus } from 'shared-data-access';
import { SnackbarService } from "shared-ui";

@Component({
  selector: 'app-add.buildings',
  imports: [
    MatButtonModule,
    MatCardModule,
    MatCheckboxModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatStepperModule,

    ReactiveFormsModule,
  ],
  templateUrl: './add.buildings.html',
  styleUrl: './add.buildings.scss',
  providers: [
    {
      provide: STEPPER_GLOBAL_OPTIONS,
      useValue: { showError : true },
    },
  ]
})
export class AddBuildings {
  private readonly _fb = inject(FormBuilder)
  private readonly _router = inject(Router)
  private readonly _snackBarService = inject(SnackbarService)
  private readonly buildingService = inject(BuildingService)

  readonly buildingStatus = signal<BuildingStatus[]>([
    'draft',
    'published',
    'archived',
  ])

  readonly publishedDateSignal = signal<Date>

  readonly firstCreateForm  = this._fb.group({
    name: ['', Validators.required],
    year_built: ['', Validators.required],
    total_floors: ['', Validators.required],
    total_units: ['', Validators.required],
    description: [''],
    status: this._fb.control<BuildingStatus | null>(null, Validators.required),
    published: [false, Validators.required]
  })

  readonly secondCreateForm  = this._fb.group({
    street: ['', Validators.required],
    house_number: ['', Validators.required],
    postal_code: ['', Validators.required],
    city: ['', Validators.required],
    country: ['', Validators.required],
    latitude: [''],
    longitude: [''],
  })

  save() {
    if (!this.firstCreateForm.valid || !this.secondCreateForm.valid) {
      return
    }

    let building = this.transformData()

    this.buildingService.create(building).subscribe({
      next: (data) => {
        console.log(`AddBuilding => response from backend : ${data}`)
        this._snackBarService.successSnackBar('Saved new building successfully.')
      },
      error: (err) => {
        console.error(`AddBuilding => Error occurred : ${err}`)
        this._snackBarService.errorSnackBar('An error occurred, please ty again.')
      },
      complete: () => {
        this.backToOverview()
      }
    })
  }

  private transformData(): Building {
    var building = {
      ...this.firstCreateForm.getRawValue(),
      ...this.secondCreateForm.getRawValue(),
    } as unknown as Building

    building.slug = building.name + '-' + building.city + '-' + building.street

    console.log(`Building with slug => ${JSON.stringify(building)}`)

    return building
  }

  cancel() {
    this.firstCreateForm.reset()
    this.secondCreateForm.reset()
    this.backToOverview()
  }

  backToOverview() {
    this._router.navigate(['content-editor'])
  }
}
