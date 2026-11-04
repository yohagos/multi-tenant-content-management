import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BuildingTable } from './building.table';

describe('BuildingTable', () => {
  let component: BuildingTable;
  let fixture: ComponentFixture<BuildingTable>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BuildingTable],
    }).compileComponents();

    fixture = TestBed.createComponent(BuildingTable);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
