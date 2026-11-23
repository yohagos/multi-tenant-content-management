import { Component, ChangeDetectionStrategy } from '@angular/core';

@Component({
  selector: 'lib-shared-ui',
  imports: [],
  template: ` <p>shared-ui works!</p> `,
  changeDetection: ChangeDetectionStrategy.Eager,
  styles: ``,
})
export class SharedUi {}
