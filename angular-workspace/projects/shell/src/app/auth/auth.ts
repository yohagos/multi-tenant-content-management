import { Component, inject, signal } from '@angular/core';
import { Register } from './register/register';
import { Login } from './login/login';
import { MatButtonModule } from '@angular/material/button';
import { AuthMask } from './services/auth.mask';

@Component({
  selector: 'app-auth',
  imports: [
    Register,
    Login,
    MatButtonModule,
],
  templateUrl: './auth.html',
  styleUrl: './auth.scss',
})
export class Auth {
  readonly authMaskService = inject(AuthMask)
  /* readonly isLogin = signal(false)

  showLogin(): void {
    this.isLogin.set(true)
  }

  showRegister(): void {
    this.isLogin.set(false)
  } */
}
