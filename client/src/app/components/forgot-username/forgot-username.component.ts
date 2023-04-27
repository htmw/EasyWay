import { Component, OnInit } from '@angular/core';
import { FormGroup,FormControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router'
import { GlobalConstants } from '../../common/global-constants';
import { AppComponent } from 'src/app/app.component';
import { Validators } from '@angular/forms';

@Component({
  selector: 'app-login',
  templateUrl: './forgot-username.component.html',
  styleUrls: ['./forgot-username.component.css']
})

export class ForgotUsernameComponent implements OnInit {
  forgotUsernameForm = new FormGroup({
    email1: new FormControl('', [Validators.required, Validators.email])
  });

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
  }

  onSubmit() {
    if (this.forgotUsernameForm.invalid) {
      return;
    }

    const email = this.forgotUsernameForm.value.email1;

    this.http.post(GlobalConstants.apiURL + 'forgotUsername', { email }).subscribe(
      (response) => {
        console.log('Username sent on email.');
        alert('Username sent on email.')
      },
      (error) => {
        console.log('Invalid email!', error);
        alert('Invalid email!')
      }
    );
  }
  highlight(event: any): void {
    event.target.style['border-bottom'] = "1px solid rgba(30, 40, 51, 0.9)";
    event.target.style['opacity'] = "0.9";
  }

  dampen(event: any): void {
    event.target.style['border-bottom'] = "1px solid rgba(30, 40, 51, 0.6)";
    event.target.style['opacity'] = "0.6";
  }
}
