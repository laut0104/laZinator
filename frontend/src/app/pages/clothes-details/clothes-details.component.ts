import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Subscription } from 'rxjs';
import { ClothesRepoService } from 'src/app/repositories/clothes-repo.service';
import { UserService } from 'src/app/services/user.service';

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
  
  public userId: number = 0;
  public clothes: any;
  public weathers: string[] = [];
  public clothId!: number;
  public all_events: any;
  public temperatures: any;
  public all_details: any;
  private subscriptions: Subscription[] = [];

  constructor(
    private userSvc: UserService,
    private clothRepoSvc: ClothesRepoService,
    private route: ActivatedRoute,
  ) { }

  ngOnInit(): void {
    this.userId = this.userSvc.user$.getValue().id
    
    this.subscriptions.push(
      this.route.params.subscribe(params => {
        this.clothId = Number(params['id']);
        console.log(this.clothId)
        this.getCloth();
      })
    );
  }

  public getCloth() {
    const query = {}
    this.clothRepoSvc.getCloth(this.clothId,query).subscribe((res: any) => {
      console.log(res);

        this.clothes = res;
        this.clothes.details = this.clothes.details.slice(1,-1)
        this.clothes.all_details = this.clothes.details.split('|,')
        console.log(this.clothes.all_details)

        this.clothes = res;
        this.clothes.weather = this.clothes.weather.slice(1,-1)
        this.clothes.weathers = this.clothes.weather.split(',')

        this.clothes = res;
        this.clothes.events = this.clothes.events.slice(1,-1)
        this.clothes.all_events = this.clothes.events.split(',')

        this.clothes = res;
        this.clothes.temperature = this.clothes.temperature.slice(1,-1)
        this.clothes.temperatures = this.clothes.temperature.split(',')

        

      // if(res.clothes){
      //   this.clothes = res.clothes;
      //   this.clothes.forEach((cloth, index) => {
      //     cloth.events = cloth.events.slice(1,-1)
      //     this.clothes[index].all_events = cloth.events.split(',')
      //   })
      // }
    })
  }
}
