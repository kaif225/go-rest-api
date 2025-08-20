package models

import "database/sql"

type Exec struct {
	ID                   int            `json:"id,omitempty" db:"id,omitempty"`
	FirstName            string         `json:"first_name,omitempty" db:"first_name,omitempty"`
	LastName             string         `json:"last_name,omitempty" db:"last_name,omitempty"`
	Email                string         `json:"email,omitempty" db:"email,omitempty"`
	Username             string         `json:"username,omitempty" db:"username,omitempty"`
	Password             string         `json:"password,omitempty" db:"password,omitempty"`
	PasswordChangedAt    sql.NullString `json:"password_changed_at" db:"password_changed_at,omitempty"`   // it will be null until a password is changed
	UserCreatedAt        sql.NullString `json:"user_created_at,omitempty" db:"user_created_at,omitempty"` // Null until a user is created
	PasswordResetToken   sql.NullString `json:"password_reset_token,omitempty"`                           // short live then again become null
	PasswordTokenExpired sql.NullString `json:""`                                                         // short lived
	InactiveStatus       bool           `json:""`
	Role                 string         `json:""`
}

/*
CREATE TABLE IF NOT EXISTS exec(
 id INtPRIMARY KEY AUTO_INCREMENT,
 first_name varchar(255) NOT NULL ,
 last_name varchar(255) NOT NULL,
 email varchar(255) NOT NULL UNIQUE,
 username varchar(255) NOT NULL UNIQUE,
 password varchar(255) NOT NULL,
 password_changed_at varchar(255),
 user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 password_reset_token VARCHAR(255),
 password_token_expired VARCHAR(255),
 inactive_status BOOLEAN NOT NULL,
 role VARCHAR(50) NOT NULL,
 INDEX idx_email (email),
 INDEX idx_username (username)
);

*/
