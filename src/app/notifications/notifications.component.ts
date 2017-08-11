import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-notifications',
  templateUrl: './notifications.component.html',
  styleUrls: ['./notifications.component.css']
})
export class NotificationsComponent implements OnInit {
    public showNotifications:boolean = false;
    public numberofNotifications = 0;

  constructor() { }

  ngOnInit() {
  }

}
