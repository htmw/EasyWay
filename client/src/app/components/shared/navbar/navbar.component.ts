import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { GlobalConstants } from 'src/app/common/global-constants';
import { NavigationEnd, NavigationStart, Router } from '@angular/router';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  isLoggedIn:boolean = false;
  currentRoute:string = '';
  searchResults: any[] = [];
  showDropdown: boolean = false;

  constructor(private http: HttpClient, public router: Router) { }

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
    if (query && query.length >= 3) {
      const searchUrl = `http://localhost:3000/api/searchServiceByName?name=${query}`;
      this.http.get(searchUrl).subscribe((data: any) => {
        console.log(data);
        const datalist = document.getElementById('search-results');
        if (datalist) {
          datalist.innerHTML = '';
          data.forEach((item: any) => {
            const option = document.createElement('option');
            option.value = item.name;
            option.setAttribute('data-id', item.id);
            datalist.appendChild(option);
          });
        }
        if (data.length > 0) {
          this.showDropdown = true;
        } else {
          this.showDropdown = false;
        }
      });
    } else {
      this.showDropdown = false;
    }
  }
  selectService(item: any) {
    this.showDropdown = false;
    console.log('Selected service:', item);
    // Do something with the selected service
  }
}
