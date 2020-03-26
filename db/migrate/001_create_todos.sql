-- Create Todos

CREATE DATABASE IF NOT EXISTS todos_development;

CREATE TABLE IF NOT EXISTS todos_development.todos (
  id INT PRIMARY KEY AUTO_INCREMENT,
  content VARCHAR(255),
  completed BOOLEAN
);
