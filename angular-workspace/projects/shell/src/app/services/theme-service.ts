import { Injectable, signal } from '@angular/core';

export type Theme = 'light' | 'dark'

@Injectable({
  providedIn: 'root',
})
export class ThemeService {
  readonly theme = signal<Theme>(
    (localStorage.getItem('theme') as Theme) || 'dark'
  )

  constructor() {
    this.applyTheme(this.theme())
  }

  toggle(): void {
    this.setTheme(this.theme() === 'light' ? 'dark' : 'light')
  }

  setTheme(theme: Theme): void {
    this.theme.set(theme)
    localStorage.setItem('theme', theme)
    this.applyTheme(theme)
  }

  private applyTheme(theme: Theme): void {
    document.documentElement.classList.toggle(
      'dark',
      theme === 'dark'
    )
  }
}
