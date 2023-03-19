import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './guards/auth.guard';
import { ClothesListComponent } from './Pages/clothes-list/clothes-list.component';
import { LiffInitComponent } from './Pages/liff-init/liff-init.component';

const routes: Routes = [
  {
    path: '',
    canActivate: [AuthGuard],
    component: LiffInitComponent,
  },
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
