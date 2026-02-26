package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	b "blog/internal/handlers"
	r "blog/internal/repository"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Config Configuration
}

func (a *App) InitializeDB() {
	a.Config = GetConfiguration()

	configString := mysql.NewConfig()
	configString.Net = "tcp"
	configString.DBName = a.Config.DBName
	configString.Passwd = a.Config.DBPassword
	configString.Addr = a.Config.DBHost
	configString.User = a.Config.DBUser

	var err error

	a.DB, err = sql.Open("mysql", configString.FormatDSN())
	fmt.Println(configString.FormatDSN())
	if err != nil {
		panic(err)
	}

	a.DB.SetMaxOpenConns(10)
	a.DB.SetMaxIdleConns(5)
	a.DB.SetConnMaxLifetime(10 * time.Minute)
	a.DB.SetConnMaxIdleTime(5 * time.Minute)

	a.Router = mux.NewRouter()
	a.InitializeRouter()
}

func (a *App) InitializeRouter() {
	database := r.NewBlogStore(a.DB)
	handler := b.NewBlogHandler(database)

	a.Router.HandleFunc("/Health", handler.HealthCheckHandler).Methods("GET")
	a.Router.HandleFunc("/Blogs", handler.ListHandler).Methods("GET")
	a.Router.HandleFunc("/Blogs/{id}", handler.GetHandler).Methods("GET")
	a.Router.HandleFunc("/Blogs/{id}", handler.PatchHandler).Methods("PATCH")
	a.Router.HandleFunc("/Blogs/{id}", handler.UpdateHandler).Methods("PUT")
	a.Router.HandleFunc("/Blogs", handler.PostHandler).Methods("POST")
	a.Router.HandleFunc("/Blogs/Delete", handler.DeleteHandler).Methods("DELETE")
	a.Router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
}

func (a *App) Run() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           a.Router,
		ReadTimeout:       500 * time.Millisecond,
		WriteTimeout:      500 * time.Millisecond,
		ReadHeaderTimeout: 500 * time.Millisecond,
		IdleTimeout:       1000 * time.Millisecond,
	}

	log.Fatal(server.ListenAndServe())
}
