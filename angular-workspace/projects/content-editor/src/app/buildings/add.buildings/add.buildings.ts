import { STEPPER_GLOBAL_OPTIONS } from '@angular/cdk/stepper';
import { Component, inject, signal } from '@angular/core';
import { FormBuilder, FormControl, ReactiveFormsModule, Validators, FormsModule } from "@angular/forms";
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
    FormsModule,
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
  readonly selectedStatus = signal<BuildingStatus>('draft')

  readonly publishedDateSignal = signal<Date>

  readonly firstCreateForm  = this._fb.group({
    name: new FormControl('', [Validators.required]),
    year_built: new FormControl('', [Validators.required]),
    total_floors: new FormControl('', [Validators.required]),
    total_units: new FormControl('', [Validators.required]),
    description: new FormControl(''),
    status: new FormControl('', [Validators.required]),
    published: new FormControl(false, [Validators.required]),
  })

  readonly secondCreateForm  = this._fb.group({
    street: new FormControl('', [Validators.required]),
    house_number: new FormControl('', [Validators.required]),
    postal_code: new FormControl('', [Validators.required]),
    city: new FormControl('', [Validators.required]),
    country: new FormControl('', [Validators.required]),
    latitude: new FormControl(''),
    longitude: new FormControl(''),
  })

  save() {
    if (!this.firstCreateForm.valid || !this.secondCreateForm.valid) {
      return
    }

    let building = this.transformData()

    console.log(`Building to be added : ${building}`)

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
    return {
      ...this.firstCreateForm.getRawValue(),
      ...this.secondCreateForm.getRawValue(),
      status: this.selectedStatus(),
    } as unknown as Building
  }

  cancel() {
    this.firstCreateForm.reset()
    this.secondCreateForm.reset()
    this.backToOverview()
  }

  backToOverview() {
    this._router.navigate(['/content-editor'])
  }
}
