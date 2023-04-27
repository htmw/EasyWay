// Import required modules
import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router'
import { GlobalConstants } from 'src/app/common/global-constants';
import { flatMap } from 'rxjs';

// Define component decorator
@Component({
  selector: 'app-bookings',
  templateUrl: './bookings.component.html',
  styleUrls: ['./bookings.component.css']
})

// Define BookingsComponent class
export class BookingsComponent implements OnInit {
// Initialize arrays to hold booking data
  myBookingsData = [];
  myBookingsName:string[] = [];
  myBookingsDesc:string[] = [];

  cancelledBookingsData = [];
  cancelledBookingsName:string[] = [];
  cancelledBookingsDesc:string[] = [];
  // Set image URL using global constant
  imageURL = GlobalConstants.imageURL;
  // Constructor function
  constructor(private http: HttpClient, public router:Router) { }
  // Function to get service name using API call
  getServiceName(id:number, cancel:boolean) {
    this.http.get<any>(GlobalConstants.apiURL+'getServiceInfo?serviceId='+id)
      .subscribe(data => {
        if (cancel) {
          this.cancelledBookingsName.push( data.name );
        } else {
          this.myBookingsName.push( data.name );
        }
      }
    )
  }
  // Function to get service description using API call
  getServiceDescription(id:number, cancel:boolean) {
    this.http.get<any>(GlobalConstants.apiURL+'getServiceInfo?serviceId='+id)
      .subscribe(data => {
        if (cancel) {
          this.cancelledBookingsDesc.push( data.description );
        } else {
          this.myBookingsDesc.push( data.description );
        }
      }
    )
  }
  // Function to cancel a booking using API call
  cancelBooking(bookingId:number) {
    this.http.get<any>(GlobalConstants.apiURL+'cancelBooking?id='+bookingId)
      .subscribe(data => {
        console.log(data)
        confirm('Service cancelled! :-(\nHope to see you again.')
        this.router.navigate(['/bookings']).then(() => {
          window.location.reload();
        });
      },
      err => {
        alert(err)
      }
    )
  }
  // Initialize component on page load
  ngOnInit(): void {
    this.http.get<any>(GlobalConstants.apiURL+'getBookings?userId='+localStorage.getItem('id'))
      .subscribe(data => {
        console.log("-----------------/ Current Bookings /-----------------")
        console.log(data);
        // Loop through booking data and get service name and description
        for (let i=0; i<data.length; i++) {
          this.getServiceName(data[i]['service_id'], false);
          this.getServiceDescription(data[i]['service_id'], false);
        }
        // Log service name and description arrays for current bookings
        console.log(this.myBookingsName);
        console.log(this.myBookingsDesc);
        // Set myBookingsData array
        this.myBookingsData = data;

      }
    )
    // Get cancelled bookings
    this.http.get<any>(GlobalConstants.apiURL+'getCancelledBookings?userId='+localStorage.getItem('id'))
      .subscribe(data => {
        console.log("-----------------/ Cancelled Bookings /-----------------")
        console.log(data);
        // Loop through booking data and get service name and description
        for (let i=0; i<data.length; i++) {
          this.getServiceName(data[i]['service_id'], true);
          this.getServiceDescription(data[i]['service_id'], true);
        }
        // Log service name and description arrays for cancelled bookings
        console.log(this.cancelledBookingsName);
        console.log(this.cancelledBookingsDesc);

        this.cancelledBookingsData = data;

      }
    )


  }

}
