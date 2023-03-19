import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiService } from '../drivers/api.service';

@Injectable({
  providedIn: 'root'
})
export class ClothesRepoService {

  constructor(
    private apiSvc: ApiService,
  ) { }

  public getClothes(
    query: any
  ): Observable<any> {
    return this.apiSvc.get(`clothes`, query)
  }
}
