import { inject, Injectable } from '@angular/core';
import { MatSnackBar } from "@angular/material/snack-bar";

@Injectable({
  providedIn: 'root',
})
export class SnackbarService {
  private readonly _snackbar = inject(MatSnackBar)

  successSnackBar(message: string) {
    this._snackbar.open(message, 'Done.')
  }

  errorSnackBar(message: string) {
    this._snackbar.open(message, 'Err')
  }
}
