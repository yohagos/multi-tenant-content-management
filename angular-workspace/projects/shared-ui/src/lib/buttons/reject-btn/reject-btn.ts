import { Component, ChangeDetectionStrategy } from '@angular/core';

@Component({
  selector: 'shared-reject-btn',
  imports: [],
  templateUrl: './reject-btn.html',
  changeDetection: ChangeDetectionStrategy.Eager,
  styleUrl: './reject-btn.scss',
})
export class RejectBtn {
  test() {
    console.log('clicked reject btn');
  }
}
