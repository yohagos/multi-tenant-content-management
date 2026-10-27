import { Component, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar';
import { Router, RouterOutlet } from '@angular/router';

export interface ContentToolbarItems {
  route: string;
  name: string;
  icon: string;
}

@Component({
  selector: 'app-content-toolbar',
  imports: [MatButtonModule, MatIconModule, MatToolbarModule, RouterOutlet],
  standalone: true,
  templateUrl: './content.toolbar.html',
  styleUrl: './content.toolbar.scss',
})
export class ContentToolbar {
  private readonly _router = inject(Router);
  readonly toolbarItems: ContentToolbarItems[] = [
    {
      route: 'content-editor/building/add',
      name: 'Building',
      icon: 'add',
    },
    {
      route: 'content-editor/unit/add',
      name: 'Unit',
      icon: 'add',
    },
  ];

  navigateTo(link: string) {
    console.log(link);
    this._router.navigate([link]);
  }
}
