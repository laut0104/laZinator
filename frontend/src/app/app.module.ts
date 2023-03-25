import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LiffInitComponent } from './pages/liff-init/liff-init.component';
import { HttpRequestInterceptor } from './interceptors/http.interceptor';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialModule } from './material/material.module';
import { HeaderComponent } from './components/header/header.component';
import { ClothesListComponent } from './pages/clothes-list/clothes-list.component';
import { FooterComponent } from './components/footer/footer.component';
import { ClothesDetailsComponent } from './pages/clothes-details/clothes-details.component';

@NgModule({
  declarations: [
    AppComponent,
    LiffInitComponent,
    HeaderComponent,
    ClothesListComponent,
    FooterComponent,
    ClothesDetailsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MaterialModule,
  ],
  providers: [
    {
      provide: HTTP_INTERCEPTORS,
      useClass: HttpRequestInterceptor,
      multi: true,
    },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
