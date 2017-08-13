import { Component } from '@angular/core';
import { IsLoggedInService } from './is-logged-in.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],


  
})




export class AppComponent {
  constructor(private IsLoggedInService:IsLoggedInService){}


  title = 'app';
  
  public loggedIn:boolean = false;




    ngOnInit() {
      this.loggedIn = this.IsLoggedInService.isLoggedIn()
      
      
  }
}
