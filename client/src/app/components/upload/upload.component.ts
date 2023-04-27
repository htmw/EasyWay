import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.css']
})
export class UploadComponent {
  previewUrl: any;
  uploadSuccess: boolean = false;
  errorMessage: string = '';
  selectedFile: File | null = null;
  selectedFileUrl: string | null = null;

  constructor(private http: HttpClient) {}

  onSubmit(event: any) {
    event.preventDefault();
    const files = event.target.file.files;
    const formData = new FormData();
    for (let i = 0; i < files.length; i++) {
      formData.append('file', files[i]);
    }

    // Make HTTP POST request to server
    this.http.post<any>('http://localhost:3000/api/createUploadedFile', formData).subscribe(
      (response) => {
        console.log('Upload successful:', response);
        this.uploadSuccess = true;
        this.errorMessage = '';
      },
      (error) => {
        console.error('Upload error:', error);
        this.uploadSuccess = false;
        this.errorMessage = 'Upload failed, please try again.';
      }
    );
  }

  onFileSelected(event: any) {
    const file: File | null = event.target.files.item(0);
    if (file) {
      this.selectedFile = file;
      const reader = new FileReader();
      reader.onload = () => {
        this.selectedFileUrl = reader.result as string;
      };
      if (this.selectedFile) {
        reader.readAsDataURL(this.selectedFile);
      }
    }
  }
}
