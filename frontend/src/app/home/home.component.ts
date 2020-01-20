import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  URL = 'http://localhost:8080/api/get/allbook'
  constructor(private httpClient: HttpClient) {}
  ngOnInit() {
    this.httpClient.get(this.URL).subscribe(data => {
      console.log(data);
    })
  }

}
