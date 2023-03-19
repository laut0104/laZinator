import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from './guards/auth.guard';
import { CodeListComponent } from './Pages/code-list/code-list.component';
import { HeaderComponent } from './Pages/header/header.component';
import { LiffInitComponent } from './Pages/liff-init/liff-init.component';



const routes: Routes = [
  {
    path: 'code-list',
    component: CodeListComponent,
  },
  {
    path: 'header',
    component: HeaderComponent,
  },
  {
    path: '',
    canActivate: [AuthGuard],
    component: LiffInitComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
