package middleware

import (
	"database/sql"
	"encoding/json" // Package to encode and decode the json into struct and vice versa
	"fmt"
	"go-post/models" // Models package where User schema is defined
	"log"
	"net/http" // Used to access the request and response object of the api
	"os"       // Used to read the environtment variable
	"strconv"  // Package used to convert string into int type

	"github.com/gorilla/mux"   // Used to get the params from the route
	"github.com/joho/godotenv" // Package used to read the .env file
	_ "github.com/lib/pq"      // Postgres golang driver
)

// Response format
type response struct {
	ID      int64  `json:"id, omitempty"`
	Message string `json:"message, omitempty"`
}

// Create connection with postgres db
func createConnection() *sql.DB {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// Return the connection
	return db

}

// CreateUser create a user in the postgres db
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	// Create an empty user of type models.User
	var user models.User

	// Decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)

	// Decode the json request to user
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	// Call insert user function and pass the user
	insertID := insertUser(user)

	// Format a response object
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	// Send the response
	json.NewEncoder(w).Encode(res)
}

// GetUser will return a single user by its id
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// Convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	// Call the GetUser function with user id to retreive a single user
	user, err := getUser(int64(id))

	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	// Send the response
	json.NewEncoder(w).Encode(user)
}

// GetAllUser will return all the users
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get all the users in the db
	users, err := getAllUser()

	if err != nil {
		log.Fatalf("Unable to get all user. %v", err)
	}

	// Send all the users as response
	json.NewEncoder(w).Encode(users)
}

// UpdateUser update user's detail in the postgres db
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Origin", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "Content-Type")

	// Get the userid from the request from the params, key is "id"
	params := mux.Vars(r)

	// Convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to conver the string into int. %v", err)
	}

	// Create an empty user of type models.User
	var user models.User

	// Decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request Body. %v", err)
	}

	// Call update user to update the user
	updatedRows := updateUser(int64(id), user)

	// Format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected", updatedRows)

	// Format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// Send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteUser delete user's detail in the postgres db
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// Convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	// Call the deleteUser, convert the int to int64
	deletedRows := deleteUser(int64(id))

	// Format the message string
	msg := fmt.Sprintf("User deleted successfully. Total rows/record affected %v", deletedRows)

	// Format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// Send the response
	json.NewEncoder(w).Encode(res)
}

// ------------------------------------------ Handler Functions ------------------------//
// Insert one user in the DB
func insertUser(user models.User) int64 {

	// Create the postgres db connection
	db := createConnection()

	// Close the db connection
	defer db.Close()

	// Create the insert sql query
	// Returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`

	// The inserted id will store in this id
	var id int64

	// Execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// Return the inserted id
	return id
}

// Get one user from the DB by it's userid
func getUser(id int64) (models.User, error) {

	// Create the postgres db connection
	db := createConnection()

	// Close the db connection
	defer db.Close()

	// Create a username of models.User type
	var user models.User

	// Create the select sql query
	sqlStatemnt := `SELECT * FROM users WHERE userid=$1`

	// Execute the sql statement
	row := db.QueryRow(sqlStatemnt, id)

	// Unmarshal the row object to user
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned")
		return user, nil
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// Return empty user on error
	return user, err
}

// Get onse user from the DB by it's userid
func getAllUser() ([]models.User, error) {

	// Create the postgres db connection
	db := createConnection()

	// Close the db connection
	defer db.Close()

	var users []models.User

	// Create the select sql query
	sqlStatement := `SELECT * FROM users`

	// Execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// Close the statement
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var user models.User

		// unmarshal the row object the user
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// Append the user in the users slice
		users = append(users, user)
	}

	// Return empty user on error
	return users, err
}

// Update user in the DB
func updateUser(id int64, user models.User) int64 {

	// Create the postgres DB connection
	db := createConnection()

	// Close the db connection
	defer db.Close()

	// Create the update sql query
	sqlStatement := `UPDATE users SET name=$2, location=$3, age=$4 WHERE userid=$1`

	// Execute the sql statement
	res, err := db.Exec(sqlStatement, id, user.Name, user.Location, user.Age)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// Check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// Delete user in the DB
func deleteUser(id int64) int64 {

	// Create the postgres db connection
	db := createConnection()

	// Close the db connection
	defer db.Close()

	// Create the delete sql query
	sqlStatement := `DELETE FROM users WHERE userid=$1`

	// Execute the sql statemnt
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// Check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}
