import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { AngularFireStorage } from '@angular/fire/compat/storage';
import { finalize } from 'rxjs/operators';

@Component({
  selector: 'app-image',
  templateUrl: './image.component.html',
  styleUrls: ['./image.component.scss']
})
export class ImageComponent implements OnInit {

  imgSrc : string | undefined;
  selectedImage: any = null;
  isSubmitted:boolean=false;

  formTemplate = new FormGroup({
    imageUrl : new FormControl('',Validators.required),
  })

  constructor(private storage:AngularFireStorage ) { 

  }

  ngOnInit(){
    this.resetFrom();
  }

  showPreview(event:any){
    if(event.target.files && event.target.files[0]){
      const reader = new FileReader();
      reader.onload = (e:any) => this.imgSrc = e.target.result;
      reader.readAsDataURL(event.target.files[0]);
      this.selectedImage = event.target.files[0];
    }else{
      this.imgSrc = '../../../assets/placeholder.jpg';
      this.selectedImage = null;
    }
  }

  onSubmit(fromValue:any){
    this.isSubmitted = true;
    if(this.formTemplate.valid){
      var filePath = `images/${this.selectedImage.name}_${new Date().getTime()}`
      const fileRef = this.storage.ref(filePath);
      this.storage.upload(filePath,this.selectedImage).snapshotChanges().pipe(
        finalize(()=>{
          fileRef.getDownloadURL().subscribe((url)=>{
            console.log(url);
            fromValue['imageUrl']=url;
            this.resetFrom();
          })
        })
      ).subscribe();
    }
  }

  get formControls(){
    return this.formTemplate['controls'];
  }

  resetFrom(){
    this.formTemplate.reset();
    this.formTemplate.setValue({
      imageUrl:'',
    });
    this.imgSrc = '../../../assets/placeholder.jpg';
    this.selectedImage = null;
    this.isSubmitted=false;
  }
}
