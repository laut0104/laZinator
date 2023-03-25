import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './guards/auth.guard';
import { ClothesListComponent } from './pages/clothes-list/clothes-list.component';
import { ImageComponent } from './pages/image/image.component';
import { LiffInitComponent } from './pages/liff-init/liff-init.component';

const routes: Routes = [
  {
    path: '',
    canActivate: [AuthGuard],
    component: LiffInitComponent,
  },
  {
    path: 'clothes-list',
    canActivate: [AuthGuard],
    component: ClothesListComponent,
  },
  {
    path: 'image',
    component: ImageComponent,children:[
      {path: 'upload', component:ImageComponent} //image/upload
    ]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
