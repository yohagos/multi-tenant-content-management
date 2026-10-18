import { Injectable, signal } from '@angular/core';
import { LoginResponse } from '../models/login-response.model';
import { User } from 'shared-data-access';

@Injectable({
  providedIn: 'root',
})
export class TokenService {
  private readonly AUTH_TOKEN = 'auth_token'
  private readonly REFRESH_TOKEN = 'refresh_token'
  private readonly USER = 'auth_user'

  private readonly tokenSignal = signal<string | null>(this.getStoredToken())
  readonly token = this.tokenSignal.asReadonly()

  private readonly userSignal = signal<User | null>(this.getStoredUser())

  setTokens(response: LoginResponse) {
    localStorage.setItem(this.AUTH_TOKEN, response.token)
    localStorage.setItem(this.REFRESH_TOKEN, response.refresh_token)
    localStorage.setItem(this.USER, JSON.stringify(response.user))

    this.tokenSignal.set(response.token)
    this.userSignal.set(response.user)
  }

  clear() {
    localStorage.removeItem(this.AUTH_TOKEN)
    localStorage.removeItem(this.REFRESH_TOKEN)
    localStorage.removeItem(this.USER)

    this.tokenSignal.set(null)
    this.userSignal.set(null)
  }

  getToken(): string | null {
    return this.tokenSignal()
  }

  getRefreshToken(): string | null {
    return localStorage.getItem(this.REFRESH_TOKEN)
  }

  getUser(): User | null {
    return this.userSignal()
  }

  isAuthenticated(): boolean {
    return this.tokenSignal() != null
  }

  private getStoredToken(): string | null {
    return localStorage.getItem(this.AUTH_TOKEN)
  }

  private getStoredUser(): User | null {
    const user = localStorage.getItem(this.USER)
    return user ? JSON.parse(user) : null
  }
}
