import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ClothesDetailsComponent } from './clothes-details.component';

describe('ClothesDetailsComponent', () => {
  let component: ClothesDetailsComponent;
  let fixture: ComponentFixture<ClothesDetailsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ClothesDetailsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ClothesDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
