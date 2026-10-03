import { Component, inject } from '@angular/core';
import { FormBuilder, FormControl, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDividerModule } from '@angular/material/divider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { AuthService } from 'shared-data-access';

@Component({
  selector: 'app-login',
  imports: [
    MatButtonModule,
    MatCardModule,
    MatDividerModule,
    MatFormFieldModule,
    MatInputModule,

    ReactiveFormsModule,
  ],
  templateUrl: './login.html',
  styleUrl: './login.scss',
})
export class Login {
  private readonly fb = inject(FormBuilder)
  private readonly authService = inject(AuthService)

  readonly loginForm = this.fb.group({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required, Validators.minLength(8)]),
  })

  onSubmit() {
    if (this.loginForm.invalid) {
      this.loginForm.markAsTouched()
      return
    }

    const data = this.loginForm.getRawValue()
    let email = data['email'] || ''
    let password = data['password'] || ''
    console.log(`Login Request => ${data.email} ${data.password}`)

    if ((email !== null || email !== '') && (password !== null || password !== '')) {
      this.authService.login(email, password).subscribe({
        next: (response) => {
          console.log(`Login => User ${email} logged in successfully!!`)
          console.log(`Login => Backend Response : ${response}`)
        },
        error: (err) => console.error('Login => error occurred : ', err)
      })
    }
  }
}
