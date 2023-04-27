import { Blog } from './blog.model';
import { User } from './user.model';

export class Comment {
  id : number;
  userId: number;
  blogId: number;
  content: string;
  created_at: Date;
  updated_at: Date;

  constructor(id: number, userId: number, blogId: number, content: string, created_at: Date, updated_at: Date) {
    this.id = id;
    this.userId = userId;
    this.blogId = blogId;
    this.content = content;
    this.created_at = created_at;
    this.updated_at = updated_at;
  }

  static createDefault(): Comment {
    return new Comment(0, 0, 0, '', new Date(), new Date());
  }
  
}
