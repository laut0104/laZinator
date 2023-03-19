import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
// // import { CdkAccordionModule } from '@angular/cdk/accordion';
// // import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatDialogModule } from '@angular/material/dialog';
// // import { MatRippleModule } from '@angular/material/core';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
// // import { MatSelectModule } from '@angular/material/select';
import { MatTableModule } from '@angular/material/table';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatChipsModule } from '@angular/material/chips';

@NgModule({
  declarations: [],
  imports: [CommonModule],
  exports: [
    MatButtonModule,
    MatIconModule,
    MatCardModule,
    // // CdkAccordionModule,
    // // MatProgressSpinnerModule,
    MatDialogModule,
    // // MatRippleModule,
    MatToolbarModule,
    MatFormFieldModule,
    MatInputModule,
    // // MatSelectModule,
    MatTableModule,
    MatSnackBarModule,
    MatChipsModule,
  ],
})
export class MaterialModule {}
