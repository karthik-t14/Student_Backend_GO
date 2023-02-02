package middleware

import (
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"github.com/gorilla/mux" // used to get the params from the route
	_ "github.com/lib/pq"    // postgres golang driver
	"net/http"               // used to access the request and response object of the api
	"strconv"                // package used to covert string into int type
	"student_crudapp/models" // models package where User schema is defined
)

//type JsonResponse struct {
//	Type    string `json:"type"`
//	Message string `json:"message"`
//}
type response struct {
	ID      int    `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func setupDB() *sql.DB {
	//dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", "dbname=studentdb user=postgres password=admin host=localhost sslmode=disable") //sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	checkErr(err)
	return db
}

//Get all Students
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi,")
	db := setupDB()

	fmt.Println("Getting students...")

	rows, err := db.Query("SELECT * FROM student order by id")

	// check errors
	checkErr(err)

	// var response []JsonResponse
	var students []models.Student

	// Foreach student
	for rows.Next() {
		var id int
		var sname string
		var age int
		var branch string

		err = rows.Scan(&id, &sname, &age, &branch)

		// check errors
		checkErr(err)
		currentStudent := models.Student{Id: id, Name: sname, Age: age, Branch: branch}
		students = append(students, currentStudent)
	}
	//msg := fmt.Sprintf(books)
	//res := response{
	//	ID:      int64(id),
	//	Message: msg,
	//}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	json.NewEncoder(w).Encode(students)
}
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	// create an empty student
	db := setupDB()
	var student models.Student

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&student)
	checkErr(err)
	var studentID int
	err = db.QueryRow(`INSERT INTO student(id, sname, age, branch) VALUES($1, $2, $3, $4) RETURNING id`, student.Id, student.Name, student.Age, student.Branch).Scan(&studentID)

	//if err != nil {
	//	return 0, err
	//}
	checkErr((err))
	fmt.Printf("Last inserted ID: %v\n", studentID)
	res := response{
		ID:      int(studentID),
		Message: "added student",
	}
	json.NewEncoder(w).Encode(res)
}
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	db := setupDB()
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	checkErr(err)
	var student models.Student
	// decode the json request to book
	err = json.NewDecoder(r.Body).Decode(&student)
	checkErr(err)
	//var updatedIds int
	res, err := db.Exec(`UPDATE student set sname=$1, age=$2, branch=$3 where id=$4 RETURNING id`, student.Name, student.Age, student.Branch, id) //.Scan(&updatedIds)
	//if err != nil {
	//	return 0, err
	//}

	rowsUpdated, err := res.RowsAffected()
	fmt.Println(rowsUpdated)
	//if err != nil {
	//	return 0, err
	//}
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", rowsUpdated)
	checkErr(err)
	//res = JsonResponse{Type: "success", Message: "updated"}
	res1 := response{
		ID:      int(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res1)
}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	db := setupDB()
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])
	checkErr(err)
	//var del int
	res, err := db.Exec(`delete from student where id = $1`, id) //.Scanf(&del)
	checkErr(err)
	deletedRows, err := res.RowsAffected()
	fmt.Println(deletedRows)
	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	// format the response message
	//res = JsonResponse{Type: "success", Message: "Deleted"}
	res1 := response{
		ID:      int(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res1)
}

//---------------------------------------------------------------------------------------------
//GetBook

func GetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi, I'm error")
	params := mux.Vars(r)

	// convert the id type from string to int
	studentId, err := strconv.Atoi(params["id"])
	fmt.Println(studentId)
	checkErr(err)
	db := setupDB()
	// create an empty user of type models.User
	//var book models.Book

	// decode the json request to user
	//err = json.NewDecoder(r.Body).Decode(&book)

	checkErr(err)
	res := models.Student{}

	var id int
	var sname string
	var branch string
	var age int

	err = db.QueryRow(`SELECT id, sname, age, branch FROM student where id = $1`, studentId).Scan(&id, &sname, &age, &branch)
	if err == nil {
		res = models.Student{Id: id, Name: sname, Age: age, Branch: branch}
	}
	fmt.Println(res)
	// send the response
	json.NewEncoder(w).Encode(res)
}
