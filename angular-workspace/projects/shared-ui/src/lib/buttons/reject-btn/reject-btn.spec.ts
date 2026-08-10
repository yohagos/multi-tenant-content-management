import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RejectBtn } from './reject-btn';

describe('RejectBtn', () => {
  let component: RejectBtn;
  let fixture: ComponentFixture<RejectBtn>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RejectBtn],
    }).compileComponents();

    fixture = TestBed.createComponent(RejectBtn);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
