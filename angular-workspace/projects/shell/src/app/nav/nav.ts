import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';

export interface NavLinks {
  name: string;
  link: string;
}

@Component({
  selector: 'shell-nav',
  standalone: true,
  imports: [],
  templateUrl: './nav.html',
  styleUrl: './nav.scss',
})
export class Nav {
  private readonly router = inject(Router)

  items: NavLinks[] = [
    {
      name: 'Tenant Admin',
      link: 'tenant-admin',
    },
    {
      name: 'Content Editor',
      link: 'content-editor',
    },
    {
      name: 'Public View',
      link: 'public-view',
    },
  ]

  navigate(link: string) {
    console.log('navigate to => ', link)
    this.router.navigate([link])
  }
}
