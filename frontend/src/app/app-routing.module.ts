import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './guards/auth.guard';
import { ClothesListComponent } from './pages/clothes-list/clothes-list.component';
import { LiffInitComponent } from './pages/liff-init/liff-init.component';
import { ClothesDetailsComponent } from './pages/clothes-details/clothes-details.component';

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
    path: 'clothes-details',
    canActivate: [AuthGuard],
    component: ClothesDetailsComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
