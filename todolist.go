package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io"
	"net/http"
)

var db,_ = gorm.Open("mysql","root:root@/todolist?charset=utf8&parseTime=True&loc=Local")

type ToDoItemModel struct{
	Id int `gorm:"primary_key"`
	Description string
	Completed bool
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Info("API health is okay")

	db.Debug().DropTableIfExists(&ToDoItemModel{})
	db.Debug().AutoMigrate(&ToDoItemModel{})

	w.Header().Set("Content-Type","application/json")
	io.WriteString(w,`{"alive":true}`)

}

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	defer db.Close()
	log.Info("Starting the TodoList API server")
	router := mux.NewRouter()
	router.HandleFunc("/healthz",Healthz).Methods("GET")
	http.ListenAndServe(":8000",router)

}

