import { ChangeDetectionStrategy, Component, computed, inject } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { MatButton, MatIconButton } from "@angular/material/button";
import { MatButtonToggleModule } from "@angular/material/button-toggle";
import { MatIconModule } from "@angular/material/icon";
import { MatMenuModule } from "@angular/material/menu";
import { MatToolbarModule } from "@angular/material/toolbar";
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ThemeService } from '../services/theme-service';
import { AuthService, TokenService } from 'shared-data-access';

export interface Links {
  name: string
  link: string
}

export interface Menu {
  title: string
  links: Links[]
}

export interface NavLinks {
  name: string
  icon: string | null
  link: string
  menu: Menu
}

@Component({
  selector: 'shell-nav',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    FormsModule,
    MatButton,
    MatButtonToggleModule,
    MatIconModule,
    MatIconButton,
    MatMenuModule,
    MatToolbarModule,
    RouterLink
],
  templateUrl: './nav.html',
  styleUrl: './nav.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Nav {
  private readonly router = inject(Router)
  readonly tokenService = inject(TokenService)
  private readonly authService = inject(AuthService)
  readonly themeService = inject(ThemeService)
  theme: string = 'light'

  items: NavLinks[] = [
    {
      name: 'Home',
      icon: 'home',
      link: '',
      menu: {
        title: 'home',
        links: [
          {
            link: 'home',
            name: 'Home',
          },
          {
            link: 'register',
            name: 'Registration',
          },
          {
            link: 'login',
            name: 'Login',
          },
        ]
      }
    },
    {
      name: 'Tenant Admin',
      icon: 'home',
      link: 'tenant-admin',
      menu: {
        title: 'tenant',
        links: [
        ]
      }
    },
    {
      name: 'Content Editor',
      icon: 'dashboard',
      link: 'content-editor',
      menu: {
        title: 'content-home',
        links: [
          {
            name: 'Overview',
            link: 'content-editor/overview',
          },
        ]
      }
    },
    {
      name: 'Public View',
      icon: 'public',
      link: 'public-view',
      menu: {
        title: 'public-home',
        links: [
        ]
      }
    },
  ]

  readonly visibleItems = computed(() => {
    const user = this.tokenService.getUser()

    if (!user) {
      return this.items.filter(item => item.name === 'Home')
    }

    switch (user.role) {
      case 'admin':
        return this.items

      case 'editor':
        return this.items.filter(items => ['Home', 'Content Editor'].includes(items.name))

      case 'tenant':
        return this.items.filter(items => ['Home', 'Content Editor'].includes(items.name))

      case 'public':
      default:
        return this.items.filter(item => item.name === 'Home')
    }
  })

  navigate(link: string) {
    this.router.navigate([link])
  }

  logout() {
    this.tokenService.clear()
    this.authService.logout()

    this.router.navigate([''])
  }
}
