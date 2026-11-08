import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Building } from 'shared-ui';

@Injectable({
  providedIn: 'root',
})
export class BuildingTransferService {
  private transferSubject = new BehaviorSubject<Building[]>([])
  transferSubject$ = this.transferSubject.asObservable()

  private selectedBuilding = new BehaviorSubject<Building | null>(null)
  selectedBuilding$ = this.selectedBuilding.asObservable()

  setBuildings(b: Building[]) {
    this.transferSubject.next(b)
  }

  setSelectedBuilding(b: Building) {
    this.selectedBuilding.next(b)
  }
}
