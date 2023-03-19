import { TestBed } from '@angular/core/testing';

import { ClothesRepoService } from './clothes-repo.service';

describe('ClothesRepoService', () => {
  let service: ClothesRepoService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ClothesRepoService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
