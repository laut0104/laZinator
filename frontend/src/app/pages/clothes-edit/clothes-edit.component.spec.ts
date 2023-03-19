import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClothesEditComponent } from './clothes-edit.component';

describe('ClothesEditComponent', () => {
  let component: ClothesEditComponent;
  let fixture: ComponentFixture<ClothesEditComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ClothesEditComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ClothesEditComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
