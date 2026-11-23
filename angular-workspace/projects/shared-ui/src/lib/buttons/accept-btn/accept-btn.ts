import { Component, ChangeDetectionStrategy } from '@angular/core';

@Component({
  selector: 'shared-accept-btn',
  standalone: true,
  imports: [],
  templateUrl: './accept-btn.html',
  changeDetection: ChangeDetectionStrategy.Eager,
  styleUrl: './accept-btn.scss',
})
export class AcceptBtn {}
