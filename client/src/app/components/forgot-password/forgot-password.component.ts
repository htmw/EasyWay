import { Component, OnInit } from '@angular/core';
import { UntypedFormGroup, UntypedFormControl, Validators, ValidationErrors, AbstractControl } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router'
import { GlobalConstants } from '../../common/global-constants';
import { AppComponent } from 'src/app/app.component';

@Component({
  selector: 'app-forgot-password',
  templateUrl: './forgot-password.component.html',
  styleUrls: ['./forgot-password.component.css']
})
export class ForgotPasswordComponent implements OnInit {
  forgotPasswordForm = new UntypedFormGroup({
    email: new UntypedFormControl('', [Validators.required, emailValidator])
  });

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
  }

  onSubmit() {
    const headers = {
      "Content-Type": "application/json"
    }
    if (this.forgotPasswordForm.invalid) {
      return;
    }

    const email = this.forgotPasswordForm.value.email;
    const message = JSON.stringify({
      email: email,
      message: "This is a test email message."
    });

    this.http.post(GlobalConstants.apiURL + 'forgotPassword', message).subscribe(
      (response) => {
        console.log('Email sent successfully.');
        alert('Email sent successfully.');
      },
      (error) => {
        console.log('Error while sending email!', error);
        alert('Error while sending email!');
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
