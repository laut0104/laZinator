import { Component, OnInit } from '@angular/core';
import { FormArray, FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ClothesRepoService } from 'src/app/repositories/clothes-repo.service';
import { UserService } from 'src/app/services/user.service';
import { AngularFireStorage } from '@angular/fire/compat/storage';
import { finalize } from 'rxjs/operators';

@Component({
  selector: 'app-clothes-add',
  templateUrl: './clothes-add.component.html',
  styleUrls: ['./clothes-add.component.scss']
})
export class ClothesAddComponent implements OnInit {

  public details = '';
  public temperature_arr = '{';
  public weather_arr = '{';
  public event_arr = '{'

  public clothForm = this.fb.group({
    imageUrl: ['',Validators.required],
    size: ['', Validators.required],
    tops: [''],
    bottoms: [''],
    weather: this.fb.array([
      this.fb.control('', [
        Validators.required,
      ])
    ]),
    temperature: this.fb.array([
      this.fb.control('', [
        Validators.required,
      ])
    ]),
    events: this.fb.array([
      this.fb.control('', [
        Validators.required
      ])
    ])
  });
  public recipe: string = '{'
  public userId: number = 0;

  constructor(
    private fb: FormBuilder,
    private userSvc: UserService,
    public router: Router,
    private clothesRepoSvc: ClothesRepoService,
    private storage:AngularFireStorage
  ) { }

  ngOnInit(){
    this.userId = this.userSvc.user$.getValue().id
  }

  get weather() {
    return this.clothForm.get('weather') as FormArray;
  }
  get temperature() {
    return this.clothForm.get('temperature') as FormArray;
  }
  get events() {
    return this.clothForm.get('events') as FormArray;
  }


  addWeather() {
    if(this.temperature.length<5){
      this.weather.push(this.fb.control('', [
        Validators.required,
      ]));
    }
  }
  addTemperature() {
    if(this.temperature.length<4){
      this.temperature.push(this.fb.control('', [
        Validators.required,
      ]));
    }
  }
  addEvents() {
    if(this.events.length<6){
      this.events.push(this.fb.control('', [
        Validators.required,
      ]));
    }
    this.events.push(this.fb.control('', [
      Validators.required,
    ]));
  }


  removeWeather() {
    if(this.weather.length-1 >0) this.weather.removeAt(this.weather.length-1);
  }
  removeTemperature() {
    if(this.temperature.length-1 >0) this.temperature.removeAt(this.temperature.length-1);
  }
  removeEvents() {
    if(this.events.length-1 >0) this.events.removeAt(this.events.length-1);
  }


  createCloth() {
    const tops = this.clothForm.value.tops != '' ? this.clothForm.value.tops : "null"
    const bottoms = this.clothForm.value.bottoms != '' ? this.clothForm.value.bottoms : "null"

    this.details = '{' + this.clothForm.value.size + '|,' + tops + '|,' + bottoms + '}'

    this.clothForm.value.temperature?.map((temperature) => {
      this.temperature_arr = this.temperature_arr + temperature + ','
    })
    this.temperature_arr = this.temperature_arr.slice(0,-1) + '}'

    this.clothForm.value.weather?.map((weather) => {
      this.weather_arr = this.weather_arr + weather + ','
    })
    this.weather_arr = this.weather_arr.slice(0,-1) + '}'

    this.clothForm.value.events?.map((event) => {
      this.event_arr = this.event_arr + event + ','
    })
    this.event_arr = this.event_arr.slice(0,-1) + '}'

    const body = {
      'userid': this.userId,
      "cloth": this.imgUrl,
      "details": this.details,
      "weather": this.weather_arr,
      "temperature": this.temperature_arr,
      "events": this.event_arr,
    }
    this.clothesRepoSvc.createCloth(body).subscribe(() => {
      // this.router.navigate([`/clothes-list`]);
    })
  }

  imgSrc : string='../assets/images/placeholder.jpg';
  imgUrl : string | undefined;
  selectedImage: any = null;
  isSubmitted:boolean=false;

  showPreview(event:any){
    if(event.target.files && event.target.files[0]){
      const reader = new FileReader();
      reader.onload = (e:any) => this.imgSrc = e.target.result;
      reader.readAsDataURL(event.target.files[0]);
      this.selectedImage = event.target.files[0];
    }else{
      this.imgSrc = '../assets/images/placeholder.jpg';
      this.selectedImage = null;
    }
  }

  onSubmit(fromValue:any){
    this.isSubmitted = true;
    if(this.clothForm.valid){
      var filePath = `images/${this.selectedImage.name}_${new Date().getTime()}`
      const fileRef = this.storage.ref(filePath);
      this.storage.upload(filePath,this.selectedImage).snapshotChanges().pipe(
        finalize(()=>{
          fileRef.getDownloadURL().subscribe((url)=>{
            this.imgUrl = url.substr(0, url.indexOf('&'));
            this.createCloth();
            fromValue['imageUrl']=url;
          })
        })
      ).subscribe();
    }
  }

  get formControls(){
    return this.clothForm['controls'];
  }
}
