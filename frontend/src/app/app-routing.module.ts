import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './guards/auth.guard';
import { ClothesEditComponent } from './pages/clothes-edit/clothes-edit.component';
import { ClothesListComponent } from './pages/clothes-list/clothes-list.component';
import { ImageComponent } from './pages/image/image.component';
import { ClothesAddComponent } from './pages/clothes-add/clothes-add.component';
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
  {
    path: 'clothes-edit',
    canActivate: [AuthGuard],
    component: ClothesEditComponent,
  },
  {
    path: 'clothes-add',
    canActivate: [AuthGuard],
    component: ClothesAddComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
