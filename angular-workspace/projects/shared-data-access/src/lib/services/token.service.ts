import { Injectable, signal } from '@angular/core';
import { LoginResponse } from '../models/login-response.model';

@Injectable({
  providedIn: 'root',
})
export class TokenService {
  private readonly AUTH_TOKEN = 'auth_token'
  private readonly REFRESH_TOKEN = 'refresh_token'

  private readonly tokenSignal = signal<string | null>(this.getStoredToken())
  readonly token = this.tokenSignal.asReadonly()

  setTokens(response: LoginResponse) {
    localStorage.setItem(this.AUTH_TOKEN, response.token)
    localStorage.setItem(this.REFRESH_TOKEN, response.refreshToken)
    this.tokenSignal.set(response.token)
  }

  clear() {
    localStorage.removeItem(this.AUTH_TOKEN)
    localStorage.removeItem(this.REFRESH_TOKEN)
    this.tokenSignal.set(null)
  }

  getToken(): string | null {
    return this.tokenSignal()
  }

  getRefreshToken(): string | null {
    return localStorage.getItem(this.REFRESH_TOKEN)
  }

  isAuthenticated(): boolean {
    return this.tokenSignal() != null
  }

  private getStoredToken(): string | null {
    return localStorage.getItem(this.AUTH_TOKEN)
  }
}
