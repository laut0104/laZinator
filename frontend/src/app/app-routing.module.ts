import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './guards/auth.guard';
import { ClothesAddComponent } from './pages/clothes-add/clothes-add.component';
import { ClothesListComponent } from './pages/clothes-list/clothes-list.component';
import { LiffInitComponent } from './pages/liff-init/liff-init.component';

const routes: Routes = [
  {
    path: '',
    canActivate: [AuthGuard],
    component: LiffInitComponent,
  },

  {
    path: 'cloth-add',
    canActivate: [AuthGuard],
    component: ClothesAddComponent,
  }
  // {
  //   path: 'clothes-list',
  //   canActivate: [AuthGuard],
  //   component: ClothesListComponent,
  // },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
