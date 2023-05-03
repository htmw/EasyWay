import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { GlobalConstants } from 'src/app/common/global-constants';
import { Router } from '@angular/router';
import { Blog } from './blog.model';
import { Comment } from './comment.model';
import { NgForm } from '@angular/forms';

@Component({
  selector: 'app-blog',
  templateUrl: './blog.component.html',
  styleUrls: ['./blog.component.css']
})
export class BlogComponent implements OnInit {
  blogs: Blog[] = [];
  comments: Comment[] = [];
  imageURL = GlobalConstants.imageURL;
  selectedBlogId: number | null = null;
  commentContent: string = '';

  constructor(private http: HttpClient) { }

  ngOnInit() {
    this.loadBlogs();
  }

  loadBlogs() {
    this.http.get<Blog[]>(GlobalConstants.apiURL + 'getAllBlogs').subscribe(blogs => {
      this.blogs = blogs;
    });
  }

  loadComments(blogId: number) {
    this.selectedBlogId = blogId;
    this.http.get<Comment[]>(GlobalConstants.apiURL + `getAllComments/${blogId}`).subscribe(comments => {
      this.comments = comments;
    });
  }

  addComment(form: NgForm) {
    const blogId = this.selectedBlogId;
    const content = form.value.content;

    const commentInput = { blogId, content };
    console.log('commentInput:', commentInput);

    this.http.post<Comment>(GlobalConstants.apiURL + `blogs/${blogId}/comments`, JSON.stringify(commentInput)).subscribe(
      comment => {
        console.log('comment added:', comment);
        this.comments.push(comment);
        form.reset();
      },
      error => {
        console.error('error adding comment:', error);
      }
    );
  }
}
