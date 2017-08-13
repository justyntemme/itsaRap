import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';

@Injectable()
export  class IsLoggedInService {

  private loggedIn: boolean = false;
  private userID: string = "";


  isLoggedIn():boolean{
    return this.loggedIn;
  }

  getSessionID():string{
    return this.userID;
  }

  setLoggedIn(isLoggedIn: boolean){
    this.loggedIn = isLoggedIn;

  }

  setSessionID(userID: string){
    this.userID = userID;
    if ( this.userID == ""){
        this.setLoggedIn(false)
    }
    this.CookieService.set('session-id', userID)
  }

  constructor(private CookieService:CookieService) { }

}
