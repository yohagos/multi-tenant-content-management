import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class UnitStatusTransferService {
  private selectedStatus = new BehaviorSubject<string | null>(null)
  selectedStatus$ = this.selectedStatus.asObservable()

  setSelectedStatus(status: string) {
    this.selectedStatus.next(status)
  }
}
