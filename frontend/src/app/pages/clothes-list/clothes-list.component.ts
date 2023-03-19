import { Component, OnInit } from '@angular/core';
import { ClothesRepoService } from 'src/app/repositories/clothes-repo.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-clothes-list',
  templateUrl: './clothes-list.component.html',
  styleUrls: ['./clothes-list.component.scss']
})
export class ClothesListComponent implements OnInit {
  public userId: number = 0;

  constructor(
    private userSvc: UserService,
    private clothesRepoSvc: ClothesRepoService,

  ) { }

  ngOnInit(): void {
    this.userId = this.userSvc.user$.getValue().id
    this.getClothes();
  }

  public getClothes() {
    const query = {
      "id": this.userId
    }
    this.clothesRepoSvc.getClothes(query).subscribe((res: any) => {
      console.log(res)
    })
  }

  test() {
    console.log('test')
  }

}
