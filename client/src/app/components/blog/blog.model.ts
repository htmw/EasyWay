import { Comment } from './comment.model';

export class Blog {
  id: number;
  title: string;
  content: string;
  created_at: Date;
  updated_at: Date;
  comments: Comment[];
  image_name: string;

  constructor(id: number, title: string, content: string, created_at: Date, updated_at: Date, comments: Comment[]) {
    this.id = id;
    this.title = title;
    this.content = content;
    this.created_at = created_at;
    this.updated_at = updated_at;
    this.comments = comments.filter(comment => comment.blogId === this.id);
    this.image_name = "";
  }
}
