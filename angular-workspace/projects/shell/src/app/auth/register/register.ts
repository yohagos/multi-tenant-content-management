import { Component, inject } from '@angular/core';
import { FormBuilder, FormControl, ReactiveFormsModule, Validators } from "@angular/forms";
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDividerModule } from "@angular/material/divider";
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from "@angular/material/input";

import { AuthService, RegisterRequest } from 'shared-data-access';

@Component({
  selector: 'app-register',
  imports: [
    MatButtonModule,
    MatCardModule,
    MatDividerModule,
    MatFormFieldModule,
    MatInputModule,

    ReactiveFormsModule,
  ],
  templateUrl: './register.html',
  styleUrl: './register.scss',
})
export class Register {
  private readonly fb = inject(FormBuilder)
  private readonly authService = inject(AuthService)

  readonly registerForm = this.fb.group({
    first_name: new FormControl('', [Validators.minLength(3)]),
    last_name: new FormControl('', [Validators.min(3)]),
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required, Validators.minLength(8)]),
    tenant_id: new FormControl(''),
  })

  onSubmit() {
    if (this.registerForm.invalid) {
      this.registerForm.markAllAsTouched()
      return
    }

    const data = this.registerForm.getRawValue()

    console.log(data)
    this.authService.register(data as RegisterRequest).subscribe({
      next: response => {
        console.log(`Register reponse for user => ${response}`)
      },
      error: err => {
        console.error(err)
      }
    })
  }
}
