import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MatButtonModule } from "@angular/material/button";
import { MatCardModule } from "@angular/material/card";
import { Auth } from '../auth/auth';

@Component({
  selector: 'shell-app-home',
  imports: [
    MatButtonModule,
    MatCardModule,

    RouterOutlet,
    Auth,
  ],
  templateUrl: './home.html',
  styleUrl: './home.scss',
})
export class Home {
}
