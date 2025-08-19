package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restapi/internal/models"
	"restapi/internal/repository/sqlconnect"
	"strconv"
	"strings"
)

// var (
// 	teachers = make(map[int]models.Teacher)
// 	//mutex    = &sync.Mutex{}
// 	//nextID = 1
// )

// func TeachersHandler(w http.ResponseWriter, r *http.Request) {

// 	switch r.Method {
// 	case http.MethodGet:
// 		getTeacherHandler(w, r)
// 	case http.MethodPut:
// 		// PUT METHOD - lecture 199
// 		updateTeacherhandler(w, r)
// 	case http.MethodPost:
// 		addTeacherHandler(w, r)

// 	case http.MethodPatch:
// 		patchTeacherHandler(w, r)
// 	case http.MethodDelete:
// 		deleteTeacherHandler(w, r)
// 	}
// }

func isValidSortOrder(order string) bool {
	return order == "asc" || order == "desc"
}

func isValidSortField(field string) bool {
	validField := map[string]bool{
		"first_name": true,
		"last_name":  true,
		"email":      true,
		"class":      true,
		"subject":    true,
	}
	return validField[field]
}

func GetTeachersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDb()

	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// To handle query params

	query := "Select ID, first_name, last_name, email, class, subject FROM teachers WHERE 1=1"

	var args []interface{}

	query = addSorting(r, query)

	query, args = addFilter(r, query, args)

	rows, err := db.Query(query, args...)

	if err != nil {
		fmt.Println(err)
		http.Error(w, fmt.Sprintf("Database query error %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	teacherList := make([]models.Teacher, 0)

	for rows.Next() {
		var teacher models.Teacher
		err := rows.Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error Scanning database results %v", err), http.StatusInternalServerError)
			return
		}
		teacherList = append(teacherList, teacher)
	}

	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "Success",
		Count:  len(teacherList),
		Data:   teacherList,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// Handle Path Parameters

}

func GetOneTeacherHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDb()

	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	idStr := r.PathValue("id")
	fmt.Println(idStr)
	////////////////
	// To handle query params

	// Handle Path Parameters
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	var teacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject from teachers WHERE id = ?", id).Scan(&teacher.ID, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Class, &teacher.Subject)
	if err == sql.ErrNoRows {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(teacher)
}

func addSorting(r *http.Request, query string) string {
	sortParams := r.URL.Query()["sortby"]

	if len(sortParams) > 0 {
		query += " ORDER BY "
		for i, param := range sortParams {
			parts := strings.Split(param, ":")
			if len(parts) != 2 {
				continue
			}
			field, order := parts[0], parts[1]

			if !isValidSortField(field) || !isValidSortOrder(order) {
				continue
			}

			if i > 0 {
				query += ","
			}
			query += " " + field + " " + order
		}
	}
	return query
}

func addFilter(r *http.Request, query string, args []interface{}) (string, []interface{}) {
	params := map[string]string{
		"first_name": "first_name",
		"last_name":  "last_name",
		"email":      "email",
		"class":      "class",
		"subject":    "subject",
	}

	for param, dbField := range params {
		value := r.URL.Query().Get(param)

		if value != "" {
			query += " AND " + dbField + " = ?"
			args = append(args, value)
		}
	}
	return query, args
}

// Func for Post teacher Handler

func AddTeacherHandler(w http.ResponseWriter, r *http.Request) {

	db, err := sqlconnect.ConnectDb()

	if err != nil {
		http.Error(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var newTeachers []models.Teacher
	err = json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO teachers (first_name, last_name, email, class, subject) VALUES(?,?,?,?,?)")

	if err != nil {
		http.Error(w, fmt.Sprintf("Error in creating sql query: %v", err), http.StatusInternalServerError)

		return
	}
	defer stmt.Close()

	addedTeachers := make([]models.Teacher, len(newTeachers))

	for i, newTeacher := range newTeachers {
		res, err := stmt.Exec(newTeacher.FirstName, newTeacher.LastName, newTeacher.Email, newTeacher.Class, newTeacher.Subject)
		if err != nil {
			http.Error(w, "Error in inserting values", http.StatusInternalServerError)
			return
		}

		lastID, err := res.LastInsertId()
		if err != nil {
			http.Error(w, "Error Getting last inserted ID", http.StatusInternalServerError)
			return
		}
		newTeacher.ID = int(lastID)
		addedTeachers[i] = newTeacher
		//addedTeachers = append(addedTeachers, newTeacher)

	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := struct {
		Status string           `json:"status"`
		Count  int              `json:"count"`
		Data   []models.Teacher `json:"data"`
	}{
		Status: "success",
		Count:  len(addedTeachers),
		Data:   addedTeachers,
	}
	json.NewEncoder(w).Encode(response)

}

// Lecture 199
func UpdateTeacherhandler(w http.ResponseWriter, r *http.Request) {
	//idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	fmt.Println("The number is : ", id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	var updatedTeacher models.Teacher

	err = json.NewDecoder(r.Body).Decode(&updatedTeacher)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// To extract the existing information
	var existingTeacher models.Teacher

	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName, &existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject) // to print single row

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		fmt.Println(err)
		http.Error(w, "Unable to retrive data", http.StatusInternalServerError)
		return
	}

	updatedTeacher.ID = existingTeacher.ID

	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", updatedTeacher.FirstName,
		updatedTeacher.LastName, updatedTeacher.Email, updatedTeacher.Class, updatedTeacher.Subject, updatedTeacher.ID)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error in updating database", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(updatedTeacher)
}

func PatchTeachersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDb()

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	/*
	  the type of updates is slice of map because we are patching key value of multiple teachers at the same time like
	  id = 101
	  email = ""
	  id = 102
	  first_name = ""
	*/
	var updates []map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusBadGateway)
		return
	}

	tx, err := db.Begin() // it starts a transaction, when you need to run a series of sql statements
	// that should either all pass or all fail then we start a transaction

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error Starting Transaction", http.StatusInternalServerError)
		return
	}

	for _, update := range updates {
		idStr, ok := update["id"].(string)
		if !ok {
			tx.Rollback()
			http.Error(w, "Teacher ID is invalid", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Error Converting ID to int", http.StatusBadRequest)
		}

		var teacherFromDb models.Teacher
		err = db.QueryRow("SELECT id , first_name, last_name, email,class, subject FROM teachers WHERE id = ?", id).Scan(
			&teacherFromDb.ID, &teacherFromDb.FirstName, &teacherFromDb.LastName, &teacherFromDb.Email, &teacherFromDb.Class, &teacherFromDb.Subject)
		if err != nil {
			tx.Rollback()
			if err == sql.ErrNoRows {
				http.Error(w, "Teacher Not Found", http.StatusNotFound)
				return
			}
			fmt.Println("Kch to hai :", err)
			http.Error(w, "Error Retriving teacher", http.StatusInternalServerError)
			return
		}
		// Apply updates using reflection
		teacherVal := reflect.ValueOf(&teacherFromDb).Elem()
		teacherType := teacherVal.Type()

		for k, v := range update {
			if k == "id" {
				continue
			}
			for i := 0; i < teacherVal.NumField(); i++ {
				field := teacherType.Field(i)
				if field.Tag.Get("json") == k+",omitempty" {
					fieldVal := teacherVal.Field(i)
					if fieldVal.CanSet() {
						val := reflect.ValueOf(v)
						if val.Type().ConvertibleTo(fieldVal.Type()) {
							fieldVal.Set(val.Convert(fieldVal.Type()))
						} else {
							tx.Rollback()
							log.Printf("connot convert %v to %v", val.Type(), fieldVal.Type())
							return
						}
					}
					break
				}
			}
		}
		_, err = tx.Exec("UPDATE teachers SET first_name = ?, last_name = ?,email = ?, class = ?, subject = ? WHERE id = ?", teacherFromDb.FirstName,
			teacherFromDb.LastName, teacherFromDb.Email, teacherFromDb.Class, teacherFromDb.Subject, teacherFromDb.ID)
		if err != nil {
			tx.Rollback()
			fmt.Println("Naya error:", err)
			http.Error(w, "Error Updating teacher", http.StatusInternalServerError)
			return
		}
	}
	// COmmit the transaction
	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error commiting transaction", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Lecture 200 - Patch function

func PatchOneTeacherHandler(w http.ResponseWriter, r *http.Request) {
	//idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
		return
	}

	var updates map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Request Payload", http.StatusInternalServerError)
		return
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error in connected to database", http.StatusInternalServerError)
		return
	}

	defer db.Close()
	var existingTeacher models.Teacher
	err = db.QueryRow("SELECT id, first_name, last_name, email, class, subject FROM teachers WHERE id = ?", id).Scan(&existingTeacher.ID, &existingTeacher.FirstName, &existingTeacher.LastName,
		&existingTeacher.Email, &existingTeacher.Class, &existingTeacher.Subject)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			http.Error(w, "Teacher not found", http.StatusNotFound)
			return
		}
		fmt.Println(err)
		http.Error(w, "Unable to retrive data", http.StatusInternalServerError)
		return
	}

	// for k, v := range updates {

	// 	switch k {
	// 	case "first_name":
	// 		existingTeacher.FirstName = v.(string)
	// 	case "last_name":
	// 		existingTeacher.LastName = v.(string)
	// 	case "email":
	// 		existingTeacher.Email = v.(string)
	// 	case "class":
	// 		existingTeacher.Class = v.(string)
	// 	case "subject":
	// 		existingTeacher.Subject = v.(string)
	// 	}
	// }

	// Using reflect package in place of swicth it is because we can have multiple keys for swicth will be haktic
	// Lecture 201
	teacherVal := reflect.ValueOf(&existingTeacher).Elem()
	// fmt.Println(teacherVal) using this if we do https://localhost:3000/teachers/100 it will show all the data for id 100
	teacherType := teacherVal.Type()
	for k, v := range updates {
		for i := 0; i < teacherVal.NumField(); i++ {
			field := teacherType.Field(i)
			fmt.Println(field.Tag.Get("json")) // it will print the tag of json that is in teachers struct like
			//"id,omitempty"
			if field.Tag.Get("json") == k+",omitempty" {
				if teacherVal.Field(i).CanSet() {
					fmt.Println("Old Value: ", teacherVal.Field(i))
					fmt.Println("New Value: ", reflect.ValueOf(v))
					teacherVal.Field(i).Set(reflect.ValueOf(v).Convert(teacherVal.Field(i).Type()))
					/*
					   teacherVal.Field(i) it will give the Old value that we are patching
					   reflect.ValueOf(v) it gives the new value that is getting replaced
					   teacherVal.Field(i).Type() it prints the type
					*/
				}
			}
		}
	}
	//////////////////////////////////

	_, err = db.Exec("UPDATE teachers SET first_name = ?, last_name = ?, email = ?, class = ?, subject = ? WHERE id = ?", existingTeacher.FirstName,
		existingTeacher.LastName, existingTeacher.Email, existingTeacher.Class, existingTeacher.Subject, existingTeacher.ID)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error in updating database", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(existingTeacher)
}

// Lecture 202 Delete method

func DeleteOneTeacherHandler(w http.ResponseWriter, r *http.Request) {
	//idStr := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Teacher ID", http.StatusBadRequest)
		return
	}

	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error in COnnecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM teachers WHERE id = ?", id)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error deleting teacher", http.StatusInternalServerError)
		return
	}
	// result is the successfull response

	rowsEffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error retriving delte result", http.StatusInternalServerError)
		return
	}

	if rowsEffected == 0 {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	//w.WriteHeader(http.StatusNoContent)
	// Response Body
	w.Header().Set("Content-Type", "application/json")

	response := struct {
		Status string `json:"status"`
		ID     int    `json:"id"`
	}{
		Status: "Teacher Successfuly deleted",
		ID:     id,
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteTeachersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sqlconnect.ConnectDb()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error COnnecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Getting multiple ids
	var ids []int

	err = json.NewDecoder(r.Body).Decode(&ids)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid Payload", http.StatusBadRequest)
		return
	}
	///////////////

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error Creating Transaction", http.StatusBadRequest)
		return
	}
	////////////////////////
	stmt, err := tx.Prepare("DELETE from teachers WHERE id = ?")
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		http.Error(w, "Error Preparing statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	deletedIds := []int{}

	for _, id := range ids {
		result, err := stmt.Exec(id)
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			http.Error(w, "Error executing command", http.StatusInternalServerError)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			fmt.Println(err)
			http.Error(w, "Error retriving deleted teachers", http.StatusInternalServerError)
			return
		}
		if rowsAffected > 0 {
			deletedIds = append(deletedIds, id)
		}
		if rowsAffected < 1 {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("ID %v does not exist", id), http.StatusBadRequest)
			return
		}
	}
	// COmmit
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		http.Error(w, "Error comminting transaction", http.StatusInternalServerError)
		return
	}

	if len(deletedIds) < 1 {
		http.Error(w, "Ids do not exists", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := struct {
		Status     string `json:"status"`
		DeletedIDs []int  `json:"deleted_ids"`
	}{
		Status:     "teachers Deleted Sucessfully",
		DeletedIDs: deletedIds,
	}

	json.NewEncoder(w).Encode(response)
}
