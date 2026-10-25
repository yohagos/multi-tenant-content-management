import { Component, inject } from '@angular/core';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';

@Component({
  selector: 'lib-snackbar',
  imports: [
    MatSnackBarModule,
  ],
  templateUrl: './snackbar.html',
  styleUrl: './snackbar.css',
})
export class Snackbar {
  private _snackBar = inject(MatSnackBar)

  successSnackBar(message: string) {
    this._snackBar.open(message, 'Done')
  }

  errorSnackBar(message: string) {
    this._snackBar.open(message, 'Error')
  }
}
