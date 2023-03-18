import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable()
export class HttpRequestInterceptor implements HttpInterceptor {

  constructor() {}

  intercept(req: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {
    let headers = req.headers.set('Content-Type', 'application/json');

    if (environment.ngrokSkipBrowserWarning)
      headers = headers.set('ngrok-skip-browser-warning', 'skip');

    const request = req.clone({ headers: headers });

    return next.handle(request);
  }
}
