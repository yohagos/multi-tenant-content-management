import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UnitTable } from './unit.table';

describe('UnitTable', () => {
  let component: UnitTable;
  let fixture: ComponentFixture<UnitTable>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [UnitTable],
    }).compileComponents();

    fixture = TestBed.createComponent(UnitTable);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
