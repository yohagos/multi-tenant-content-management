import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ContentToolbar } from './content.toolbar';

describe('ContentToolbar', () => {
  let component: ContentToolbar;
  let fixture: ComponentFixture<ContentToolbar>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ContentToolbar],
    }).compileComponents();

    fixture = TestBed.createComponent(ContentToolbar);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
