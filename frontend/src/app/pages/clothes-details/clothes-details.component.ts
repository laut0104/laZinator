import { Component, OnInit } from '@angular/core';

export interface PeriodicElement {
  name: string;
  position: number;
}

@Component({
  selector: 'app-clothes-details',
  templateUrl: './clothes-details.component.html',
  styleUrls: ['./clothes-details.component.scss']
})
export class ClothesDetailsComponent implements OnInit {
  
  ELEMENT_DATA: PeriodicElement[] = [
    {position: 1, name: 'Hydrogen'},
    {position: 2, name: 'Helium'},
    {position: 3, name: 'Lithium'},
    {position: 4, name: 'Beryllium'},
    {position: 5, name: 'Boron'},
    {position: 6, name: 'Carbon'}
  ];
  constructor() { }

  ngOnInit(): void {
  }

  displayedColumns: string[] = ['position', 'name'];
  dataSource = this.ELEMENT_DATA;
}
