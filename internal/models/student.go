package models

type Student struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Class     string `json:"class"`
}

/*
1/ first in teachers table run this
CREATE INDEX idx_class ON teachers(class)


2. CREATE TABLE IF NOT EXISTS students(
  id INT PRIMARY KEY AUTO_INCREMENT,
  first_name varchar(255) NOT NULL,
  last_name varchar(255) NOT NULL,
  email varchar(255) NOT NULL UNIQUE,
  class varchar(255) NOT NULL,
  INDEX (email),
  FOREIGN KEY (class) REFERENCES teachers(class)
)AUTO_INCREMENT=100;

*/
