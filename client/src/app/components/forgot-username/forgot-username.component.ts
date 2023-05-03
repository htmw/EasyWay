import { Component, OnInit } from '@angular/core';
import { UntypedFormGroup, UntypedFormControl, Validators, ValidationErrors, AbstractControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router'
import { GlobalConstants } from '../../common/global-constants';
import { AppComponent } from 'src/app/app.component';

@Component({
  selector: 'app-forgot-username',
  templateUrl: './forgot-username.component.html',
  styleUrls: ['./forgot-username.component.css']
})
export class ForgotUsernameComponent implements OnInit {
  forgotUsernameForm = new UntypedFormGroup({
    email1: new UntypedFormControl('', [Validators.required, emailValidator])
  });

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
  }

  onSubmit() {
    const headers = {
      "Content-Type": "application/x-www-form-urlencoded"
    }
    if (this.forgotUsernameForm.invalid) {
      return;
    }

    const email = this.forgotUsernameForm.value.email1;
    const message = JSON.stringify({
      email: email,
      message: "This is a test email message."
    });

    this.http.post(GlobalConstants.apiURL + 'forgotUsername', message).subscribe(
      (response) => {
        console.log('Username sent on email.');
        alert('Username sent on email.');
      },
      (error) => {
        console.log('Error while sending username!', error);
        alert('Error while sending username!');
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

function emailValidator(control: AbstractControl): ValidationErrors | null {
  const emailRegex = /[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}/;

  if (!control.value) {
    return null;
  }

  const isValid = emailRegex.test(control.value);
  return isValid ? null : { invalidEmail: true };
}
