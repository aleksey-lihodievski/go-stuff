package orm

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

type Todo struct {
	gorm.Model
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	log.Println("Get todos")

	if r.Method != http.MethodGet {
		return
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("error opening db")
	}

	todos := []Todo{}

	db.Find(&todos)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(todos)

	byteArray, err := json.MarshalIndent(todos, "", "  ")

	if err != nil {
		fmt.Println("Error marshalling", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteArray)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Create todos")

	if r.Method != http.MethodPost {
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type is not supported", http.StatusUnsupportedMediaType)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var todo Todo

	err := decoder.Decode(&todo)

	if err != nil {
		fmt.Println(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db.Create(&todo)

	// fmt.Println(body)

	fmt.Fprintf(w, "Method: %s", r.Method)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Update todo")

	if r.Method != http.MethodPut {
		return
	}

	fmt.Fprintf(w, "Method: %s", r.Method)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete todo")

	if r.Method != http.MethodDelete {
		return
	}

	// r.URL.Query().Get("id")
	param := strings.TrimPrefix(r.URL.Path, "/todo/delete/")

	if param == "" {
		fmt.Println("Error")
		http.Error(w, "Bad request: no param provided", http.StatusBadRequest)
		return
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	var todo Todo

	db.First(&todo, "id = ?", param)
	db.Delete(&todo)
}

func migrate() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("error opening db")
	}

	db.AutoMigrate(&Todo{})
}

func serve() {
	migrate()

	mux := http.NewServeMux()

	mux.HandleFunc("/todo/create", createTodo)
	mux.HandleFunc("/todo/update", updateTodo)
	mux.HandleFunc("/todo/delete/", deleteTodo)
	mux.HandleFunc("/todo", getTodos)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}

func Run() {
	//serve()
}
