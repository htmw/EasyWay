import { Component, OnInit } from '@angular/core';
import { FormGroup,FormControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router'
import { GlobalConstants } from '../../common/global-constants';
import { AppComponent } from 'src/app/app.component';
import { Validators } from '@angular/forms';


@Component({
  selector: 'app-forgot-password',
  templateUrl: './forgot-password.component.html',
  styleUrls: ['./forgot-password.component.css']
})

export class ForgotPasswordComponent implements OnInit {
  forgotPasswordForm = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email])
  });

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
  }

  onSubmit() {
  if (this.forgotPasswordForm.invalid) {
    return;
  }

  const email = this.forgotPasswordForm.value.email;

  this.http.post(GlobalConstants.apiURL + '/forgot-password', { email }).subscribe(
    (response) => {
      console.log('Password reset email sent.');
      // TODO: show success message to user
    },
    (error) => {
      console.log('Error resetting password:', error);
      // TODO: show error message to user
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
