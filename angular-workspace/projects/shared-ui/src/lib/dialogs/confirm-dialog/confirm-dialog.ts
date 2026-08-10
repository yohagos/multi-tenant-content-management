import { Component, inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogContent, MatDialogRef, MatDialogTitle, MatDialogActions } from '@angular/material/dialog';
import { AcceptBtn, RejectBtn } from '../../../public-api';

@Component({
  selector: 'shared-confirm-dialog',
  imports: [
    AcceptBtn,
    RejectBtn,
    MatDialogTitle,
    MatDialogContent,
    MatDialogActions
],
  templateUrl: './confirm-dialog.html',
  styleUrl: './confirm-dialog.scss',
})
export class ConfirmDialog {
  dialogRef = inject(MatDialogRef<ConfirmDialog>)
  data = inject(MAT_DIALOG_DATA)

  title: string = this.data['title']
  subtitle: string = this.data['subtitle']

  confirm() {
    this.dialogRef.close(true)
  }

  decline() {
    this.dialogRef.close(false)
  }
}
