CREATE DATABASE IF NOT EXISTS kpt;
use kpt;
CREATE TABLE IF NOT EXISTS user(
  user_id INT PRIMARY KEY AUTO_INCREMENT,
  user_name VARCHAR(100),
  slack_id VARCHAR(100) UNIQUE,
  created_at TIMESTAMP NOT NULL default current_timestamp,
  updated_at TIMESTAMP NOT NULL default current_timestamp on update current_timestamp
) DEFAULT CHARACTER SET utf8mb4;

