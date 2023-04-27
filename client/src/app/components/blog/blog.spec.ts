import { Blog } from './blog.model';

describe('Blog', () => {
  it('should create an instance', () => {
    const blog: Blog = {
      id: 1,
      title: 'Test Blog',
      content: 'This is a test blog post.',
      date: new Date('2023-04-24')
    };
    expect(blog).toBeTruthy();
  });
});
