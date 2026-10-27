import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ContentOverview } from './content.overview';

describe('ContentOverview', () => {
  let component: ContentOverview;
  let fixture: ComponentFixture<ContentOverview>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ContentOverview],
    }).compileComponents();

    fixture = TestBed.createComponent(ContentOverview);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
