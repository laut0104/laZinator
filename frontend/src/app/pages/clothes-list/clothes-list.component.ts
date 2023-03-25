import { Component, OnInit } from '@angular/core';
import { Cloth } from 'src/app/models/model';
import { ClothesRepoService } from 'src/app/repositories/clothes-repo.service';
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
  imgUrl: string[]=[];

  constructor(
    private userSvc: UserService,
    private clothesRepoSvc: ClothesRepoService,
    private storage: AngularFireStorage,

  ) { }

  ngOnInit(): void {
    this.userId = this.userSvc.user$.getValue().id
    this.getClothes();

    this.storage.ref('images/test.jpg').getDownloadURL().subscribe(image => {
      this.imgUrl = image;
    });
  }

  public getClothes() {
    const query = {}
    this.clothesRepoSvc.getClothes(query).subscribe((res: any) => {
      if(res.clothes){
        this.clothes = res.clothes;
        this.clothes.forEach((cloth, index) => {
          cloth.weather = cloth.weather.slice(1,-1)
          this.clothes[index].weathers = cloth.weather.split(',')
        })
      }
    })
  }

}
