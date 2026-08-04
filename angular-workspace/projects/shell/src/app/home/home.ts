import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'shell-app-home',
  imports: [
    RouterOutlet,
  ],
  templateUrl: './home.html',
  styleUrl: './home.scss',
})
export class Home {}
