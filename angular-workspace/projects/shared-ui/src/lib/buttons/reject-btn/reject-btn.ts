import { Component } from '@angular/core';

@Component({
  selector: 'shared-reject-btn',
  imports: [],
  templateUrl: './reject-btn.html',
  styleUrl: './reject-btn.scss',
})
export class RejectBtn {
  test() {
    console.log('clicked reject btn')
  }
}
