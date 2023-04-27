-- This SQL code creates a table called ucc_blog with columns id, title, content, created_at, and updated_at.

CREATE TABLE ucc_blog (
  id INT NOT NULL AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  content VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

-- This SQL code inserts sample data into the ucc_blog table.

insert into ucc_blog values (0, 'My First Blog Post', 'This is my first blog post.', NOW(), NOW());
insert into ucc_blog values (1, 'My Second Blog Post', 'This is my second blog post.', NOW(), NOW());
insert into ucc_blog values (2, 'My Third Blog Post', 'This is my third blog post.', NOW(), NOW());
