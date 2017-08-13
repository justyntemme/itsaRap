import { Component, OnInit,Input } from '@angular/core';

import {HttpClient} from '@angular/common/http'

import {Md5} from 'ts-md5/dist/md5'





@Component({
  selector: 'app-create-user',
  templateUrl: './create-user.component.html',
  styleUrls: ['./create-user.component.css'],
  providers: [Md5]
})
export class CreateUserComponent implements OnInit {
  public showCA:boolean = false;



 

  constructor(private http: HttpClient) { }



  ngOnInit() {
 
  }

  CreateUser(username, password){

    var hashedPass = Md5.hashStr(password);

    var body = `{"user_Name": "` + username + `", "user_Pass": "` + hashedPass + `"}`;
    console.log(body)
    var headers = new Headers();
    

    headers.append('Content-Type', 'application/json');

    this.http.post('http://127.0.0.1:8080/user',body).subscribe(
      (response) => {
        console.log("VALUE RECEIVED: " +response);
      },
      (err) =>{
        console.log("ERROR: " +err);
        alert("User not created. Username already taken")
      },
      () => {
        console.log("Completed");
      }
    );
    




   

}
}