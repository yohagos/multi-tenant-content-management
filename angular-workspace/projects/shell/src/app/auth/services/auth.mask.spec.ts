import { TestBed } from '@angular/core/testing';

import { AuthMask } from './auth.mask';

describe('AuthMask', () => {
  let service: AuthMask;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AuthMask);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
