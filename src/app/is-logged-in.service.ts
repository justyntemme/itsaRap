import { Injectable } from '@angular/core';

@Injectable()
export default class IsLoggedInService {

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
  }

  constructor() { }

}
