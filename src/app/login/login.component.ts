import { Component, OnInit } from '@angular/core';

import {Md5} from 'ts-md5/dist/md5'


import {HttpClient} from '@angular/common/http'



@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  providers: [Md5],
})
export class LoginComponent implements OnInit {

  constructor(private http: HttpClient) { }

  ngOnInit() {
  }

  Login(username, password) {
     var headers = new Headers();
     var hashedpassword = Md5.hashStr(password)

     var body = `{"user_Name": "` + username + `", "user_Pass": "` + hashedpassword + `"}`;

    headers.append('Content-Type', 'application/json');

    this.http.post('http://127.0.0.1:8080/user',body).subscribe(
      (response) => {
        console.log("VALUE RECEIVED: " +response);
      },
      (err) =>{
        console.log("ERROR: " +err);
        alert("incorrect username:Password")
      },
      () => {
        console.log("Completed");
      }
    );
  }

}
