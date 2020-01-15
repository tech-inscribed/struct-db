package main

import (
	"fmt"
	"net/http"

	"github.com/techinscribed/struct-db/controllers"
	"github.com/techinscribed/struct-db/sqldb"
)

func main() {
	db := sqldb.ConnectDB()

	h := controllers.NewBaseHandler(db)

	http.HandleFunc("/", h.HelloWorld)

	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	}

	s.ListenAndServe()

}
