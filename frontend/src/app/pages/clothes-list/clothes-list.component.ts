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
  public temperatures: string[] = [];
  imgUrl: string[]=[];

  constructor(
    private userSvc: UserService,
    private clothesRepoSvc: ClothesRepoService,
    private storage: AngularFireStorage,

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

}
