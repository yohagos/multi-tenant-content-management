import { Overlay, OverlayRef } from '@angular/cdk/overlay';
import { TemplatePortal } from '@angular/cdk/portal';
import { Component, inject, signal, TemplateRef, ViewChild, ViewContainerRef } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatAnchor } from '@angular/material/button';
import { UnitStatusTransferService } from './unit.status.transfer.service';

@Component({
  selector: 'app-unit-status-select',
  imports: [MatAnchor],
  templateUrl: './unit.status.select.html',
  styleUrl: './unit.status.select.scss',
  standalone: true,
})
export class UnitStatusSelect {
  @ViewChild('dropdownTemplate') dropdownTemplate!: TemplateRef<unknown>
  private readonly overlay = inject(Overlay)
  private readonly viewContainerRef = inject(ViewContainerRef)
  private unitStatusTransferService = inject(UnitStatusTransferService)

  selectedStatus = signal<string>('')

  statusList = signal<string[]>([
    'draft',
    'available',
    'reserved',
    'rented',
    'archived',
  ])

  private overlayRef?: OverlayRef
  readonly controll = new FormControl<string | null>(null)

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

  select(status: string) {
    this.unitStatusTransferService.setSelectedStatus(status)
    this.controll.setValue(status)
    this.close()
  }

  close() {
    this.overlayRef?.dispose()
    this.overlayRef = undefined
  }
}
