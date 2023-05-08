// Import required modules
import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router'
import { GlobalConstants } from 'src/app/common/global-constants';
import { HttpHeaders } from '@angular/common/http';
import { flatMap } from 'rxjs';
import { HttpErrorResponse } from '@angular/common/http';

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

  // Declare selectedBooking and updateBookingDialogVisible
  selectedBooking: any;
  updateBookingDialogVisible = false;

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

  // Function to edit a booking using API call
  showUpdateBookingDialog(service: any) {
    console.log('Updating booking:', service);
    // Set the booking data in the component's state
    this.selectedBooking = {
      id: service.id,
      userId: service.user_id,
      serviceId: service.service_id,
      date: service.date,
      startTime: service.start_time,
      endTime: service.end_time,
      note: service.note
    };

    // Show the update booking dialog
    this.updateBookingDialogVisible = true;
  }

  closeUpdateBookingDialog() {
  this.updateBookingDialogVisible = false;
}


  // Initialize component on page load
  ngOnInit(): void {
    this.loadMyBookings();
  }

  loadMyBookings(): void {
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

updateBooking() {
  console.log('Updating booking with the following details:', this.selectedBooking);

  // Create a new Date object with the updated date and time
  const newDate = new Date(this.selectedBooking.date + 'T' + this.selectedBooking.startTime);
  const newEndDate = new Date(this.selectedBooking.date + 'T' + this.selectedBooking.endTime);

  // Update the selected booking object with the new date and time
  this.selectedBooking.date = newDate.toISOString().substring(0, 10); // keep only the date part
  this.selectedBooking.startTime = newDate.toLocaleTimeString('en-US', {hour12: false, hourCycle: 'h23'}).substring(0, 5); // keep only the hour and minute part
  this.selectedBooking.endTime = newEndDate.toLocaleTimeString('en-US', {hour12: false, hourCycle: 'h23'}).substring(0, 5); // keep only the hour and minute part


  // Send the updated booking to the server
  const requestBody = {
  id: this.selectedBooking.id,
  user_id: this.selectedBooking.userId,
  service_id: this.selectedBooking.serviceId,
  date: new Date(this.selectedBooking.date).toISOString().substring(0, 10),
  start_time: this.selectedBooking.startTime,
  end_time: this.selectedBooking.endTime,
  note: this.selectedBooking.note
};

  console.log('Request body:', requestBody);
  this.http.put<any>(GlobalConstants.apiURL+'updateBooking', requestBody).subscribe(
      (data: any) => {
        console.log('Booking updated successfully:', data);
        // Hide the update booking dialog
        this.updateBookingDialogVisible = false;
        // Reload the page to reflect the updated booking
        window.location.reload();
      },
      (error: any) => {
        console.error('Error updating booking:', error);
        if (error instanceof HttpErrorResponse) {
          console.error(`Status: ${error.status}, Error: ${error.error}`);
        } else {
          console.error('An error occurred:', error.message);
        }
      }
    );


}

}
