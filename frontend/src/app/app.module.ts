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
import { ClothesEditComponent } from './pages/clothes-edit/clothes-edit.component';
import { ClothesAddComponent } from './pages/clothes-add/clothes-add.component';
import { FooterComponent } from './components/footer/footer.component';
import { AngularFireModule } from '@angular/fire/compat';
import { AngularFireAuthModule } from '@angular/fire/compat/auth';
import { AngularFireStorageModule } from '@angular/fire/compat/storage';
import { environment } from '../environments/environment';
import { AngularFireDatabaseModule } from "@angular/fire/compat/database";
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    AppComponent,
    LiffInitComponent,
    HeaderComponent,
    ClothesListComponent,
    FooterComponent,
    ClothesEditComponent,
    ClothesAddComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MaterialModule,
    AngularFireModule.initializeApp(environment.firebase),
    AngularFireAuthModule,
    AngularFireStorageModule,
    ReactiveFormsModule,
    AngularFireDatabaseModule,
    FormsModule,
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
