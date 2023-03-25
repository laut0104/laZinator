import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClothesAddComponent } from './clothes-add.component';

describe('ClothesAddComponent', () => {
  let component: ClothesAddComponent;
  let fixture: ComponentFixture<ClothesAddComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ClothesAddComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ClothesAddComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
