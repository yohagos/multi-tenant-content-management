import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AcceptBtn } from './accept-btn';

describe('AcceptBtn', () => {
  let component: AcceptBtn;
  let fixture: ComponentFixture<AcceptBtn>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AcceptBtn],
    }).compileComponents();

    fixture = TestBed.createComponent(AcceptBtn);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
