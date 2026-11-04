import { Injectable, signal } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class AuthMask {
  readonly isLogin = signal(false)

  showLogin() {
    this.isLogin.set(false)
  }

  showRegister() {
    this.isLogin.set(true)
  }
}
