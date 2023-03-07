import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { NavigationEnd, NavigationStart, Router } from '@angular/router';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  isLoggedIn:boolean = false;
  currentRoute:string = '';

  constructor(private http: HttpClient,public router:Router) { }

  ngOnInit(): void {
    if (localStorage.getItem('isLoggedIn') == 'true') {
      this.isLoggedIn = true;
    } else {
      this.isLoggedIn = false;
    }

    this.router.events.subscribe((event) => {
      if (event instanceof NavigationStart) {
        this.currentRoute = event.url;
      }
    });
  }

  searchServices(event: Event) {
    const query = (event.target as HTMLInputElement)?.value;
    if (query && query.length >= 3) { // only search if the query is at least 3 characters long
      this.http.get('/api/searchServices?q=' + query).subscribe((data: any) => {
        // display the search results in the console for now
        console.log(data);
      });
    }
  }
}
