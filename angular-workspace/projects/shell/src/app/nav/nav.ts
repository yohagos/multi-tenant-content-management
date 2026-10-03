import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { MatButton, MatIconButton } from "@angular/material/button";
import { MatButtonToggleModule } from "@angular/material/button-toggle";
import { MatIconModule } from "@angular/material/icon";
import { MatMenuModule } from "@angular/material/menu";
import { MatToolbarModule } from "@angular/material/toolbar";
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ThemeService } from '../services/theme-service';

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

  navigate(link: string) {
    this.router.navigate([link])
  }
}
