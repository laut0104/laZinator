import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ApiService } from '../drivers/api.service';

@Injectable({
  providedIn: 'root'
})
export class ProposeService {

  constructor(
    private apiSvc: ApiService
  ) { }

  public getTemp(
    query: any
  ): Observable<any> {
    return this.apiSvc.get(`temp`, query)
  }
}
