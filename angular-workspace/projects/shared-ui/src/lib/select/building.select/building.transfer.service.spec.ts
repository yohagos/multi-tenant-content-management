import { TestBed } from '@angular/core/testing';

import { BuildingTransferService } from './building.transfer.service';

describe('BuildingTransferService', () => {
  let service: BuildingTransferService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BuildingTransferService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
