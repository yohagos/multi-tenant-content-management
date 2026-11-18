import { Component, signal, ChangeDetectionStrategy } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
  selector: 'tenant-admin-root',
  standalone: true,
  imports: [RouterOutlet],
  changeDetection: ChangeDetectionStrategy.Eager,
  template: `<router-outlet></router-outlet>`,
  /* templateUrl: './app.html',
  styleUrl: './app.scss', */
})
export class App {
  protected readonly title = signal('tenant-admin');
}
