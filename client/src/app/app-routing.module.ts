import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AppComponent } from './app.component';
import { RegisterComponent } from './components/register/register.component';
import { ServicesComponent } from './components/services/services.component';
import { LandingComponent } from './components/landing/landing.component';
import { HttpClientModule } from '@angular/common/http';
import { LoginComponent } from './components/login/login.component';
import { BookingsComponent } from './components/bookings/bookings.component';
import { ServiceBookingComponent } from './components/service-booking/service-booking.component';
import { LogoutComponent } from './components/logout/logout.component';
import { UserProfileComponent } from './components/user-profile/user-profile.component';
import { NavbarComponent } from './components/shared/navbar/navbar.component';
import { ForgotPasswordComponent } from './components/forgot-password/forgot-password.component';
import { ForgotUsernameComponent } from './components/forgot-username/forgot-username.component';
import { BlogComponent } from './components/blog/blog.component';
import { UploadComponent } from './components/upload/upload.component';

// Define the routes for the application
const routes: Routes = [
  {
    path: 'register',  // URL path for the register component
    component: RegisterComponent,  // Register component to display
  },
  {
    path: 'login',  // URL path for the login component
    component: LoginComponent,  // Login component to display
  },
  {
    path: 'home',  // URL path for the landing component
    component: LandingComponent  // Landing component to display
  },
  {
    path: '',  // Default URL path
    redirectTo: '/home',  // Redirect to the landing component by default
    pathMatch: 'full'  // Ensure the entire URL matches the default path
  },
  {
    path: 'blog',  // URL path for the blog component
    component: BlogComponent  // Blog component to display
  },
  {
    path: 'services',  // URL path for the services component
    component: ServicesComponent  // Services component to display
  },
  {
    path: 'bookings',  // URL path for the bookings component
    component: BookingsComponent  // Bookings component to display
  },
  {
    path: 'bookService',  // URL path for the service booking component
    component: ServiceBookingComponent  // Service booking component to display
  },
  {
    path: 'logout',  // URL path for the logout component
    component: LogoutComponent  // Logout component to display
  },
  {
    path: 'profile',  // URL path for the user profile component
    component: UserProfileComponent  // User profile component to display
  },
  {
    path: 'forgot-password',  // URL path for the forgot password component
    component: ForgotPasswordComponent  // Forgot password component to display
  },
  {
    path: 'forgot-username',  // URL path for the forgot username component
    component: ForgotUsernameComponent  // Forgot username component to display
  },
  {
    path: 'upload',  // URL path for the upload component
    component: UploadComponent  // Upload component to display
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes), HttpClientModule],  // Import the router module and HTTP client module
  exports: [RouterModule, HttpClientModule]  // Export the router module and HTTP client module
})

// Define the AppRoutingModule class
export class AppRoutingModule { }
