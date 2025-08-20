// Data Validation

package main

import (
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strings"
)

func execsRoute() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /execs/login", handlers.LoginHandler)
	mux.HandleFunc("POST /execs/logout", handlers.LogoutHandler)
	return mux
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.Exec
	// Data Validation

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Body payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and password is required", http.StatusBadRequest)
		return
	}
	// Search for user if user exists
	db, err := sqlconnect.ConnectDb()

	if err != nil {
		http.Error(w, "Error Connecting to database", http.StatusInternalServerError)
		return
	}

	defer db.Close()
	user := &models.Exec{}
	err = db.QueryRow(`Select id, first_name, last_name, email, username, password,inactive_status, role 
	FROM exec WHERE username = ?`, req.Username).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email,
		&user.Username, &user.Password, &user.InactiveStatus, &user.Role)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusBadRequest)
			return
		}
		fmt.Println(err)
		http.Error(w, "Database Query error", http.StatusBadRequest)
		return
	}

	// is user active

	if user.InactiveStatus {
		http.Error(w, "Account is Inactive", http.StatusForbidden)
		return
	}

	// verify the password
	parts := strings.Split(user.Password, ".")
	if len(parts) != 2 {
		http.Error(w, "invalid encoded hash format", http.StatusForbidden)
		return
	}

	saltBase64 := parts[0]
	hashedPasswordbase64 := parts[1]

	salt, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to decode salt", http.StatusForbidden)
		return
	}
	hashedPassword, err := base64.StdEncoding.DecodeString(hashedPasswordbase64)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to decode hashpassword", http.StatusForbidden)
		return
	}
	hash := argon2.IDKey([]byte(req.Password), salt, 1, 64*1024, 4, 32)

	if len(hash) != len(hashedPassword) {
		http.Error(w, "Incorrect Password", http.StatusForbidden)
		return
	}
	if subtle.ConstantTimeCompare(hash, hashedPassword) == 1 {
		// do nothing
	} else {
		http.Error(w, "Incorrect Password", http.StatusForbidden)
		return
	}

}
