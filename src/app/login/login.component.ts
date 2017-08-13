import { Component, OnInit } from '@angular/core';

import {Md5} from 'ts-md5/dist/md5'

import { IsLoggedInService } from '../is-logged-in.service';

import {HttpClient} from '@angular/common/http'



@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css'],
  providers: [Md5],
})
export class LoginComponent implements OnInit {

  constructor(private http: HttpClient, private IsLoggedInService: IsLoggedInService) { }

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
        this.IsLoggedInService.setLoggedIn(true)
      },
      (err) =>{
        console.log("ERROR: " +err);
        alert("incorrect username:Password")
        this.IsLoggedInService.setLoggedIn(false)
      },
      () => {
        console.log("Completed");
      }
    );
  }

}
