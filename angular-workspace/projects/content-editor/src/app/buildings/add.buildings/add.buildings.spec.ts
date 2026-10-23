import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddBuildings } from './add.buildings';

describe('AddBuildings', () => {
  let component: AddBuildings;
  let fixture: ComponentFixture<AddBuildings>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AddBuildings],
    }).compileComponents();

    fixture = TestBed.createComponent(AddBuildings);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
