import { Component, inject } from '@angular/core';
import { FormBuilder, FormControl, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDividerModule } from '@angular/material/divider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { AuthService, LoginResponse, TokenService } from 'shared-data-access';
import { AuthMask } from '../services/auth.mask';
import { Router } from '@angular/router';

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
  private readonly tokenService = inject(TokenService)
  readonly authMaskService = inject(AuthMask)

  private readonly router = inject(Router)

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

    if ((email !== null || email !== '') && (password !== null || password !== '')) {
      this.authService.login(email, password).subscribe({
        next: (response: LoginResponse) => {
          this.tokenService.setTokens(response)
          let user = this.tokenService.getUser()
          if (user !== null) {
            this.afterLoginRoute(user.role)
          }
        },
        error: (err) => {
          console.error('Login => error occurred : ', err)
        },
      })
    }
  }

  private afterLoginRoute(role: string) {
    switch (role) {
      case 'editor':
        this.router.navigate(['content-editor/overview'])
        break;
      default:
        break;
    }
  }

  logout() {
    this.tokenService.clear()
    this.authService.logout()
  }

}


