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

  public getCloth(
    id: number,
    query: any
  ): Observable<any> {
    return this.apiSvc.get(`cloth/${id}`, query)
  }


  public createCloth(
    body: any
  ): Observable<any> {
    return this.apiSvc.post(`cloth`, body)
  }
}
