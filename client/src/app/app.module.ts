import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { HttpClientModule } from '@angular/common/http';
import { ServicesComponent } from './components/services/services.component';
import { LandingComponent } from './components/landing/landing.component';
import { ServiceBookingComponent } from './components/service-booking/service-booking.component';
import { NavbarComponent } from './components/shared/navbar/navbar.component';
import { FooterComponent } from './components/shared/footer/footer.component';
import { BookingsComponent } from './components/bookings/bookings.component';
import { ForgotPasswordComponent } from './components/forgot-password/forgot-password.component';
import { ForgotUsernameComponent } from './components/forgot-username/forgot-username.component';
import { BlogComponent } from './components/blog/blog.component';
import { UploadComponent } from './components/upload/upload.component';
// import * as $ from 'jquery';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegisterComponent,
    ServicesComponent,
    LandingComponent,
    ServiceBookingComponent,
    NavbarComponent,
    FooterComponent,
    BookingsComponent,
    ForgotPasswordComponent,
    ForgotUsernameComponent,
    BlogComponent,
    UploadComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
