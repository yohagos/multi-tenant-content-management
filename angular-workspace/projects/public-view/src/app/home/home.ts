import { Component, inject, ChangeDetectionStrategy } from '@angular/core';
import { MatButton } from '@angular/material/button';
import { MatDialog } from '@angular/material/dialog';
import { ConfirmDialog } from '../../../../shared-ui/src/public-api';

@Component({
  selector: 'public-view-home',
  standalone: true,
  imports: [MatButton],
  templateUrl: './home.html',
  changeDetection: ChangeDetectionStrategy.Eager,
  styleUrl: './home.scss',
})
export class Home {
  readonly dialog = inject(MatDialog);

  openDialog() {
    this.dialog.open(ConfirmDialog, {
      width: '500px',
      disableClose: true,
      data: {
        title: 'Test',
        subtitle: '',
      },
    });
  }
}
