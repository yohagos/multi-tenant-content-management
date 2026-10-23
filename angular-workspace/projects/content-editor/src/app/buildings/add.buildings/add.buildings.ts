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
import { BuildingService, BuildingStatus } from 'shared-data-access';

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
    year_build: new FormControl('', [Validators.required]),
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
    latitude: new FormControl('', [Validators.required]),
    longitude: new FormControl('', [Validators.required]),
  })

  
}
