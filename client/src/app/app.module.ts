// Importing the required modules from Angular and our own components
import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { ReactiveFormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BlogComponent } from './components/blog/blog.component';
import { BookingsComponent } from './components/bookings/bookings.component';
import { FooterComponent } from './components/shared/footer/footer.component';
import { LandingComponent } from './components/landing/landing.component';
import { LoginComponent } from './components/login/login.component';
import { NavbarComponent } from './components/shared/navbar/navbar.component';
import { RegisterComponent } from './components/register/register.component';
import { ServiceBookingComponent } from './components/service-booking/service-booking.component';
import { ServicesComponent } from './components/services/services.component';
import { UploadComponent } from './components/upload/upload.component';
import { ForgotPasswordComponent } from './components/forgot-password/forgot-password.component';
import { ForgotUsernameComponent } from './components/forgot-username/forgot-username.component';

@NgModule({
  // Declaring all the components used in the module
  declarations: [
    AppComponent,
    BlogComponent,
    BookingsComponent,
    FooterComponent,
    LandingComponent,
    LoginComponent,
    NavbarComponent,
    RegisterComponent,
    ServiceBookingComponent,
    ServicesComponent,
    UploadComponent,
    ForgotPasswordComponent,
    ForgotUsernameComponent,
  ],
  // Importing required modules
  imports: [
    BrowserModule, // Required module for browser rendering
    AppRoutingModule, // Required module for app routing
    FormsModule, // Required module for template-driven forms
    HttpClientModule, // Required module for HTTP requests
    ReactiveFormsModule, // Required module for reactive forms
  ],
  // Bootstrapping the app component
  bootstrap: [AppComponent],
})
// Exporting the app module
export class AppModule {}
