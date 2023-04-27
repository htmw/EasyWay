-- Create a table to store comments on UCC blog posts
CREATE TABLE ucc_comment (
  id INT PRIMARY KEY,
  user_id INT UNSIGNED,
  blog_id INT UNSIGNED,
  content VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES ucc_users(id),
  FOREIGN KEY (blog_id) REFERENCES ucc_blog(id),
  );

-- Insert some sample comments for testing purposes
INSERT INTO ucc_comment VALUES (1, 1, 1, 'Great article, very informative!', '2023-04-25 09:30', '2023-04-25 09:30');
INSERT INTO ucc_comment VALUES (2, 2, 1, 'I totally agree, on-demand services are the way to go!', '2023-04-25 10:15', '2023-04-25 10:15');
INSERT INTO ucc_comment VALUES (3, 3, 2, 'I had a bad experience with an on-demand cleaning service, they didn\'t do a good job', '2023-04-25 11:20', '2023-04-25 11:20');
