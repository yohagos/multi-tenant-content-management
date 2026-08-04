import { Component, signal } from '@angular/core';

@Component({
  selector: 'public-view-root',
  imports: [],
  templateUrl: './app.html',
  styleUrl: './app.scss',
})
export class App {
  protected readonly title = signal('public-view');
}
