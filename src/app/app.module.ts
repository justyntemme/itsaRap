import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import {HttpClientModule} from '@angular/common/http';
import { IsLoggedInService} from './is-logged-in.service'
import { CookieService} from 'ngx-cookie-service'


import { AppComponent } from './app.component';
import { PostsComponent } from './posts/posts.component';
import { LoggedInComponent } from './logged-in/logged-in.component';
import { NotificationsComponent } from './notifications/notifications.component';
import { LoginComponent } from './login/login.component';
import { PostComponent } from './post/post.component';
import { CommentsComponent } from './comments/comments.component';
import { CreateUserComponent } from './create-user/create-user.component';
import { MakeaPostComponent } from './makea-post/makea-post.component';

@NgModule({
  declarations: [
    AppComponent,
    PostsComponent,
    LoggedInComponent,
    NotificationsComponent,
    LoginComponent,
    PostComponent,
    CommentsComponent,
    CreateUserComponent,
    MakeaPostComponent,


  ],
  imports: [
    BrowserModule,
    HttpClientModule,
  

 
  ],
  exports: [
  
  ],
  providers: [IsLoggedInService, CookieService],
  bootstrap: [AppComponent]
})
export class AppModule { }
