import { Component, signal, ChangeDetectionStrategy } from '@angular/core';

@Component({
  selector: 'public-view-root',
  imports: [],
  templateUrl: './app.html',
  changeDetection: ChangeDetectionStrategy.Eager,
  styleUrl: './app.scss',
})
export class App {
  protected readonly title = signal('public-view');
}
