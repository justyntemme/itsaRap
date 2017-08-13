import { Component } from '@angular/core';
import { IsLoggedInService } from './is-logged-in.service';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],


  
})


//https://stackoverflow.com/questions/12576798/angularjs-how-to-watch-service-variables

export class AppComponent {
  constructor(public IsLoggedInService:IsLoggedInService, private CookieService: CookieService){}


  title = 'app';
  
  public loggedIn:boolean = false;
  public sessionID: string = "";



    ngOnInit() {
      this.sessionID = this.CookieService.get('session-id');
      console.log(this.sessionID)
      if (this.sessionID != "" ){
        this.IsLoggedInService.setLoggedIn(true);

      }

      this.loggedIn = this.IsLoggedInService.isLoggedIn()
  }
}
