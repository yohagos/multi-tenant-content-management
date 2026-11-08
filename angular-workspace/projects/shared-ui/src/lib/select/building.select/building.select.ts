import { Overlay, OverlayRef } from '@angular/cdk/overlay';
import { TemplatePortal } from '@angular/cdk/portal';
import { Component, EventEmitter, inject, OnInit, signal, TemplateRef, ViewChild, ViewContainerRef } from '@angular/core';
import { FormControl, ReactiveFormsModule } from "@angular/forms";
import { MatAnchor } from '@angular/material/button';
import { BuildingTransferService } from './building.transfer.service';

export interface Building {
  id?: string
  tenant_id?: string
  name: string
  street: string
  house_number: string
  postal_code: string
  city: string
  country: string
}

@Component({
  selector: 'app-building-select',
  imports: [
    ReactiveFormsModule,
    MatAnchor
],
  templateUrl: './building.select.html',
  styleUrl: './building.select.scss',
  standalone: true,
})
export class BuildingSelect implements OnInit {
  @ViewChild('dropdownTemplate') dropdownTemplate!: TemplateRef<unknown>
  private readonly overlay = inject(Overlay)
  private readonly viewContainerRef = inject(ViewContainerRef)
  private readonly buildingTransferService = inject(BuildingTransferService)

  private overlayRef?: OverlayRef

  readonly buildingList = signal<Building[]>([])
  readonly selectedBuilding = signal<Building | null>(null)

  readonly controll = new FormControl<string | null>(null)

  ngOnInit(): void {
    this.buildingTransferService.transferSubject$.subscribe({
      next: (data) => {
        if (data.length > 0) {
          this.buildingList.set(data)
        }
      }
    })
  }

  toggle(): void {
    if (this.overlayRef?.hasAttached()) {
      this.close()
      return
    }

    const positionStrategy = this.overlay
      .position()
      .flexibleConnectedTo(this.viewContainerRef.element)
      .withPositions([
        {originX: 'start', originY: 'bottom', overlayX: 'start', overlayY: 'top', offsetY: 4},
        {originX: 'start', originY: 'top', overlayX: 'start', overlayY: 'bottom', offsetY: -4},
      ])

    this.overlayRef = this.overlay.create({
      positionStrategy,
      hasBackdrop: true,
      backdropClass: 'cdk-overlay-transparent-backdrop',
      scrollStrategy: this.overlay.scrollStrategies.reposition(),
      minWidth: this.viewContainerRef.element.nativeElement.offsetWidth,
    })

    this.overlayRef.attach(new TemplatePortal(this.dropdownTemplate, this.viewContainerRef))
    this.overlayRef.backdropClick().subscribe(() => this.close())
  }

  select(building: Building) {
    if (building && building.id !== undefined) {
      this.buildingTransferService.setSelectedBuilding(building)
      this.selectedBuilding.set(building)
      this.controll.setValue(building.id)
      this.close()
    }
  }

  private close() {
    this.overlayRef?.dispose()
    this.overlayRef = undefined
  }
}
