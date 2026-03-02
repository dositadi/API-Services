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
	m "blog/internal/middleware"
	r "blog/internal/repository"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	Config Configuration
}

func (a *App) InitializeBlogRouter(handler b.BlogHandler) {
	// Blog handlers.
	a.Router.HandleFunc("/", handler.ListHandler).Methods("GET")
	a.Router.HandleFunc("/s{id}", handler.GetHandler).Methods("GET")
	a.Router.HandleFunc("/{id}", handler.PatchHandler).Methods("PATCH")
	a.Router.HandleFunc("/{id}", handler.UpdateHandler).Methods("PUT")
	a.Router.HandleFunc("/", handler.PostHandler).Methods("POST")
	a.Router.HandleFunc("/delete", handler.DeleteHandler).Methods("DELETE")
	a.Router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
}

func (a *App) InitializeUserRouter(handler b.BlogHandler) {
	// Blog handlers.
	// implement the handlers
	a.Router.HandleFunc("/", handler.ListHandler).Methods("GET")
	a.Router.HandleFunc("/{id}", handler.GetHandler).Methods("GET")
	a.Router.HandleFunc("/{id}", handler.PatchHandler).Methods("PATCH")
	a.Router.HandleFunc("/{id}", handler.UpdateHandler).Methods("PUT")
	a.Router.HandleFunc("/", handler.PostHandler).Methods("POST")
	a.Router.HandleFunc("/", handler.DeleteHandler).Methods("DELETE")
	a.Router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
}

func (a *App) InitializeCommentsRouter(handler b.BlogHandler) {
	// Blog handlers.
	// implement the handlers
	a.Router.HandleFunc("/", handler.ListHandler).Methods("GET")
	a.Router.HandleFunc("/{id}", handler.GetHandler).Methods("GET")
	a.Router.HandleFunc("/{id}", handler.PatchHandler).Methods("PATCH")
	a.Router.HandleFunc("/{id}", handler.UpdateHandler).Methods("PUT")
	a.Router.HandleFunc("/", handler.PostHandler).Methods("POST")
	a.Router.HandleFunc("/delete", handler.DeleteHandler).Methods("DELETE")
	a.Router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
}

func (a *App) InitializetagsRouter(handler b.BlogHandler) {
	// Blog handlers.
	// implement the handlers
	a.Router.HandleFunc("/", handler.ListHandler).Methods("GET")
	a.Router.HandleFunc("/{id}", handler.GetHandler).Methods("GET")
	a.Router.HandleFunc("/{id}", handler.PatchHandler).Methods("PATCH")
	a.Router.HandleFunc("/{id}", handler.UpdateHandler).Methods("PUT")
	a.Router.HandleFunc("/", handler.PostHandler).Methods("POST")
	a.Router.HandleFunc("/delete", handler.DeleteHandler).Methods("DELETE")
	a.Router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
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

	r := a.Router.NewRoute().Subrouter()
	userRoute := r.PathPrefix("/users")
	blogRoutes := r.PathPrefix("/blogs")
	commentRoute := r.PathPrefix("/comments")
	tagRoute := r.PathPrefix("/tags")

	userRoute.Subrouter().Use(m.AuthenticateJWT)
	blogRoutes.Subrouter().Use(m.AuthenticateJWT)
	commentRoute.Subrouter().Use(m.AuthenticateJWT)
	tagRoute.Subrouter().Use(m.AuthenticateJWT)

	a.InitializeBlogRouter(*handler)
	a.InitializeCommentsRouter(*handler)
	a.InitializeUserRouter(*handler)
	a.InitializetagsRouter(*handler)

	a.Router.HandleFunc("/health", handler.HealthCheckHandler).Methods("GET")
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
