import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UnitStatusSelect } from './unit.status.select';

describe('UnitStatusSelect', () => {
  let component: UnitStatusSelect;
  let fixture: ComponentFixture<UnitStatusSelect>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [UnitStatusSelect],
    }).compileComponents();

    fixture = TestBed.createComponent(UnitStatusSelect);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
