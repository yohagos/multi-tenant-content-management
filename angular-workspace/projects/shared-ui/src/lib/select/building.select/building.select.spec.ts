import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BuildingSelect } from './building.select';

describe('BuildingSelect', () => {
  let component: BuildingSelect;
  let fixture: ComponentFixture<BuildingSelect>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [BuildingSelect],
    }).compileComponents();

    fixture = TestBed.createComponent(BuildingSelect);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
