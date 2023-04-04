import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.css']
})
export class UploadComponent {
  constructor(private http: HttpClient) {}
  onSubmit(event: any) {
  event.preventDefault();
  const files = event.target.file.files;
  const formData = new FormData();
  for (let i = 0; i < files.length; i++) {
    formData.append('file', files[i]);
  }

  // Make HTTP POST request to server
  this.http.post('/api/upload', formData).subscribe(
    (response) => {
      console.log('Upload successful:', response);
    },
    (error) => {
      console.error('Upload error:', error);
    }
  );
 }
}
