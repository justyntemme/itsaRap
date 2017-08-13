import { Component, OnInit } from '@angular/core';

import {HttpClient} from '@angular/common/http'
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-makea-post',
  templateUrl: './makea-post.component.html',
  styleUrls: ['./makea-post.component.css'],
  providers: [CookieService, HttpClient]
})
export class MakeaPostComponent implements OnInit {

  constructor(private http: HttpClient, public CookieService: CookieService) { }

  ngOnInit() {
  }


   postIpost(soundcloudURL, Title) {

     var sessionId = this.CookieService.get('session-id');

    var body  = `{"user_Id": "` + this.CookieService.get('session-id') + `", "post_title":  "` + Title + `", "soundcloud_URL": "` + soundcloudURL + `"}`;


     this.http.post('http://127.0.0.1:8080/posts',body).subscribe(
      (response) => {
        console.log("VALUE RECEIVED: " +JSON.stringify(response));
 
      },
      (err) =>{
        console.log("ERROR: " +err);
 
      },
      () => {
        console.log("Completed");
      }
    );

  }

}

