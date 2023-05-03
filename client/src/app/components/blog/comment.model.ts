import { Blog } from './blog.model';

export class Comment {
  id: number;
  blogId: number;
  content: string;
  created_at: Date;
  updated_at: Date;

  constructor(id: number, blogId: number, content: string, created_at: Date, updated_at: Date = new Date()) {
      this.id = id;
      this.blogId = blogId;
      this.content = content;
      this.created_at = created_at;
      this.updated_at = updated_at;
  }

  static createDefault(): Comment {
    return new Comment(0, 0, '', new Date(), new Date());
  }
}
