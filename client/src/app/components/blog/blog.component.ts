import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router'
import { GlobalConstants } from 'src/app/common/global-constants';
import { flatMap } from 'rxjs';
import { map } from 'rxjs/operators';
import { Blog } from './blog.model';
import { Comment } from './comment.model';

@Component({
  selector: 'app-blog',
  templateUrl: './blog.component.html',
  styleUrls: ['./blog.component.css']
})
export class BlogComponent implements OnInit {

  posts: any[] = [];
  comments: any[] = [];
  imageURL = GlobalConstants.imageURL;
  newComment: Comment = new Comment(0, 0, 0, '', new Date(), new Date());

  constructor(private http: HttpClient, private router: Router) { }

  readBlogs(service: any) {
    if (localStorage.getItem('isLoggedIn') == 'false') {
      alert('Please log-in to book to a service.')
    } else {
      this.router.navigate(['/blogs']);
    }
  }

  ngOnInit(): void {
    this.http.get<any>(GlobalConstants.apiURL + 'getAllBlogs')
      .subscribe(data => {
        console.log(data);
        this.posts = data;
      }, err => {
        console.log(err);
      });
  }

  getAllComments(blogId: number) {
    this.http.get<any>(GlobalConstants.apiURL + 'getAllComments?blog_id=' + blogId)
      .subscribe(data => {
        console.log(data);
        if (data && data.length > 0) {
          this.comments = data;
        } else {
          this.comments = [];
          console.log("No comments available for this blog post.");
        }
      }, err => {
        console.log(err);
      });
  }

  addComment(post: Blog) {
    this.newComment.blogId = post.id;
    this.newComment.created_at = new Date();
    this.http.post<any>(GlobalConstants.apiURL + 'addComment', this.newComment)
      .subscribe(data => {
        console.log(data);
        this.comments.push(data);
        this.newComment = new Comment(0, 0, 0, '', new Date(), new Date());
      }, err => {
        console.log(err);
      });
  }

  goToBlogDetails(id: number): void {
    // Call getAllComments() with the selected blog post's ID
    this.getAllComments(id);
    this.router.navigate(['/blog', id]);
  }

  onSubmit(blogId: number): void {
      this.newComment.blogId = blogId;
      this.newComment.created_at = new Date();
      this.http.post<Comment>(GlobalConstants.apiURL + 'addComment', this.newComment)
        .subscribe(data => {
          console.log(data);
          this.getAllComments(blogId);
          this.newComment = new Comment(0, 0, 0, '', new Date(), new Date());
        }, err => {
          console.log(err);
        });
    }
}
