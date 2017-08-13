import { Component, OnInit } from '@angular/core';
import { IsLoggedInService } from '../is-logged-in.service'

@Component({
  selector: 'app-logged-in',
  templateUrl: './logged-in.component.html',
  styleUrls: ['./logged-in.component.css'],
  providers: [IsLoggedInService]
})
export class LoggedInComponent implements OnInit {

constructor(private IsLoggedInService: IsLoggedInService) { }

  logout(){
    this.IsLoggedInService.setSessionID("")
    console.log("loggedOut");
    window.location.reload()
  }

  ngOnInit() {
  }

}
