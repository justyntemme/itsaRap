import { Component, OnInit } from '@angular/core';

import {HttpClient} from '@angular/common/http'
import {IsLoggedInService} from '../is-logged-in.service'

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css'],
  providers: [IsLoggedInService]
})
export class PostComponent implements OnInit {

  constructor(private http: HttpClient, public IsLoggedInService: IsLoggedInService) { }

  ngOnInit() {
  }

 
}