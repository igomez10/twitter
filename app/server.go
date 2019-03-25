package main

import "fmt"
import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"
import "log"

import "net/http"

type client struct {
	DB *gorm.DB
}

var connection *gorm.DB

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	c := client{}
	conn := c.OpenConnection("postgres", "password", "postgres", "localhost", 5432)
	c.DB = conn
	defer c.DB.Close()

	http.HandleFunc("/health", c.serveHealth)
	http.HandleFunc("/user/", c.RegisterUser).Put()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Closed server")
	}
}

func (c *client) OpenConnection(user, password, database, host string, port int) *gorm.DB {
	if connection != nil {
		return connection
	}
	baseURL := "postgres://%s:%s@%s:%d/%s?sslmode=disable"
	URL := fmt.Sprintf(baseURL, user, password, host, port, database)
	db, err := gorm.Open("postgres", URL)
	if err != nil {
		log.Println("Could not connect to database", err)
		return nil
	}
	connection = db
	fmt.Println("Connected to DB ✅")
	return db

}

func (c *client) RegisterUser(w http.ResponseWriter, r *http.Request) {
	num := 5
	if c.UserExists(num) {
		fmt.Fprint(w, "CONFLICT")
		return
	}

}

func (c client) UserExists(ID int) bool {
	var wantedUser User
	c.DB.First(&wantedUser, "ID = ?", ID)
	if wantedUser.ID != 0 {
		return true
	}
	return false
}

func (c *client) serveHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
