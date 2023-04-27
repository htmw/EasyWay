// Import necessary modules from Angular
import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

// Declare component metadata
@Component({
  selector: 'app-root', // Selector for the component
  templateUrl: './app.component.html', // Location of component's HTML template
  styleUrls: ['./app.component.css'] // Location of component's CSS styles
})

// Declare component class
export class AppComponent implements OnInit {
  title = 'EasyWay'; // Component title
  data: any; // Placeholder for retrieved data

  // Inject HttpClient module into component
  constructor(private http: HttpClient) { }

  // Execute component logic on component initialization
  ngOnInit() {
    // Retrieve data from API endpoint and assign to 'data' variable
    this.http.get('https://jsonplaceholder.typicode.com/posts').subscribe((response: any) => {
      this.data = response;
    });
  }
}
