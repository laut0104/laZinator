import { Component, OnInit } from '@angular/core';
import { Cloth } from 'src/app/models/model';
import { ClothesRepoService } from 'src/app/repositories/clothes-repo.service';
import { ProposeService } from 'src/app/services/propose.service';
import { UserService } from 'src/app/services/user.service';
import { AngularFireStorage } from '@angular/fire/compat/storage';

@Component({
  selector: 'app-clothes-list',
  templateUrl: './clothes-list.component.html',
  styleUrls: ['./clothes-list.component.scss']
})
export class ClothesListComponent implements OnInit {
  public userId: number = 0;
  public clothes: any[] = [];
  public weathers: string[] = [];
  public lat: string | null = "";
  public lon: string | null = "";
  public temperatures: string[] = [];
  imgUrl: string[]=[];

  constructor(
    public userSvc: UserService,
    private clothesRepoSvc: ClothesRepoService,
    private proposeSvc: ProposeService,
  ) { }

  ngOnInit(): void {
    this.userId = this.userSvc.user$.getValue().id
    this.getClothes();
  }

  public getClothes() {
    const query = {}
    this.clothesRepoSvc.getClothes(query).subscribe((res: any) => {
      if(res.clothes){
        console.log(res.clothes);
        this.clothes = res.clothes;
        this.clothes.forEach((cloth, index) => {
          cloth.temperature = cloth.temperature.slice(1,-1)
          this.clothes[index].temperatures = cloth.temperature.split(',')
        })
      }
    })
  }

  public getMyposition() {
    const options = {
      enableHighAccuracy: true,
      maximumAge: 30000,
      timeout: 27000
    };
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition((position) => {
        var latitude = position.coords.latitude.toString();
        var longitude = position.coords.longitude.toString();
        const body = {
          "lat": latitude,
          "lon": longitude,
        }
        this.proposeSvc.proposeCloth(body).subscribe((res) => console.log(res))
        console.log("Latitude: " + latitude + " Longitude: " + longitude);
      }, (error) => {
        console.log(error);
      }, );
    } else {
      console.log("Geolocation is not supported by this browser.");
    }
  }

}
