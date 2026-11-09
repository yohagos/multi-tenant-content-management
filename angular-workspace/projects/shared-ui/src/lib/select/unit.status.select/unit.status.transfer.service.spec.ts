import { TestBed } from '@angular/core/testing';

import { UnitStatusTransferService } from './unit.status.transfer.service';

describe('UnitStatusTransferService', () => {
  let service: UnitStatusTransferService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UnitStatusTransferService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
