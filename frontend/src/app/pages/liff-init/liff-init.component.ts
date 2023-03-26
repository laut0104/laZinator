import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-liff-init',
  templateUrl: './liff-init.component.html',
  styleUrls: ['./liff-init.component.scss']
})
export class LiffInitComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
    navigator.geolocation.getCurrentPosition(this.successCallback, this.errorCallback);
  }

  successCallback(position: any) {
    console.log(position)
  }
  errorCallback(error: any){
    console.log(error)
  }
}
