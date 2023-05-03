import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';

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
  uploadResponse: any;
  imageUrl: SafeUrl | undefined;
  showDetectionImage: boolean = false;
  // Define the lists of furniture types and plumbing types
  furnitureTypes = ['bench', 'chair', 'couch', 'bed', 'dining table'];
  plumbingTypes = ['sink', 'toilet'];
  // Define the service information for furniture repair and plumbing
  furnitureService = { service_id: 4, service_name: 'Furniture Repair', service_price: 70 };
  plumbingService = { service_id: 2, service_name: 'Plumbing', service_price: 100 };


  constructor(private http: HttpClient, private sanitizer: DomSanitizer) {}

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
        this.uploadResponse = response; // Store response in a class variable
        const fileUrl = response.response[0].fileUrl; // Access the fileUrl property
        this.showDetectionImage = true;
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

  fetchDetectionImage() {
    const folderName = 'detections';
    const url = `http://localhost:3000/api/detection?folderName=${folderName}`;
    this.http.get(url, { responseType: 'blob' }).subscribe(
      (response: Blob) => {
        const imageUrl = URL.createObjectURL(response);
        this.imageUrl = this.sanitizer.bypassSecurityTrustUrl(imageUrl);
      },
      (error) => {
        console.error('Error fetching image:', error);
        this.errorMessage = 'Error fetching image';
      }
    );
  }
}
