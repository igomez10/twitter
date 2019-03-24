package main

import "fmt"
import "net"
import "net/http"

type client struct {
	db *net.Conn
}

var connection *net.Conn

func main() {
	c := client{}
	http.HandleFunc("/health", c.serveHealth)
	http.HandleFunc("/user/", c.RegisterUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Closed server")
	}
}

func openConnection() *net.Conn {
	var pClient client
	if connection != nil {
		return connection
	} else {
		// TODO open connection
		// TODO assign connection to pClient
		// TODO return pointer to pClient
		pClient = client{db: &connection}
	}
	return &pClient
}

func 

func (c *client) RegisterUser(w http.ResponseWriter, r *http.Request) {
	//sql.Open("mysql", dataSourceName string)
}

func (c *client) serveHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
