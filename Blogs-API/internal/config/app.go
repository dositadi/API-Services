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

func (a *App) InitializeBlogRouter(router *mux.Router, handler b.BlogHandler) {
	// Blog handlers.
	router.HandleFunc("/", handler.ListHandler).Methods("GET").Name("list-all-blogs")
	router.HandleFunc("/{id}", handler.GetHandler).Methods("GET").Name("get-blog")
	router.HandleFunc("/{id}", handler.PatchHandler).Methods("PATCH").Name("update-blog-fields")
	router.HandleFunc("/{id}", handler.UpdateHandler).Methods("PUT").Name("update-full-blog")
	router.HandleFunc("/", handler.PostHandler).Methods("POST").Name("post-blog")
	router.HandleFunc("/delete", handler.DeleteHandler).Methods("DELETE").Name("delete-blog")
	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
}

func (a *App) InitializeAuthRouter(router *mux.Router, handler b.BlogHandler) {
	// Auth handlers
	router.HandleFunc("/register", handler.RegisterUserHandler).Methods("POST").Name("register")
	router.HandleFunc("/login", handler.LoginUserHandler).Methods("POST").Name("login")
}

func (a *App) InitializeUserRouter(router *mux.Router, handler b.BlogHandler) {
	// User handlers.
	// implement the handlers
	router.HandleFunc("/", handler.ListHandler).Methods("GET")
	router.HandleFunc("/{id}", handler.GetHandler).Methods("GET")
	router.HandleFunc("/{id}", handler.PatchHandler).Methods("PATCH")
	router.HandleFunc("/{id}", handler.UpdateHandler).Methods("PUT")
	router.HandleFunc("/", handler.PostHandler).Methods("POST")
	router.HandleFunc("/", handler.DeleteHandler).Methods("DELETE")
	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
}

func (a *App) InitializeCommentsRouter(router *mux.Router, handler b.BlogHandler) {
	// Comment handlers.
	// implement the handlers
	router.HandleFunc("/", handler.ListCommentsHandler).Methods("GET").Name("list-all-comments")
	router.HandleFunc("/{id}", handler.GetCommentHandler).Methods("GET").Name("get-comment")
	router.HandleFunc("/{id}", handler.PatchCommentHandler).Methods("PATCH").Name("update-comment-field")
	router.HandleFunc("/{id}", handler.UpdateCommentHandler).Methods("PUT").Name("update-full-comment")
	router.HandleFunc("/", handler.PostCommentHandler).Methods("POST").Name("post-comment")
	router.HandleFunc("/delete", handler.DeleteCommentHandler).Methods("DELETE").Name("delete-comment")
	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
}

func (a *App) InitializetagsRouter(router *mux.Router, handler b.BlogHandler) {
	// Tags handlers.
	// implement the handlers
	router.HandleFunc("/", handler.ListHandler).Methods("GET")
	router.HandleFunc("/{id}", handler.GetHandler).Methods("GET")
	router.HandleFunc("/{id}", handler.PatchHandler).Methods("PATCH")
	router.HandleFunc("/{id}", handler.UpdateHandler).Methods("PUT")
	router.HandleFunc("/", handler.PostHandler).Methods("POST")
	router.HandleFunc("/delete", handler.DeleteHandler).Methods("DELETE")
	router.NotFoundHandler = http.HandlerFunc(handler.NotFound)
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

	authRoute := r.PathPrefix("/auth").Subrouter()

	userRoute := r.PathPrefix("/users").Subrouter()
	userRoute.Use(m.AuthenticateJWT)

	blogRoute := r.PathPrefix("/blogs").Subrouter()
	blogRoute.Use(m.AuthenticateJWT)

	commentRoute := r.PathPrefix("/comments").Subrouter()
	commentRoute.Use(m.AuthenticateJWT)

	tagRoute := r.PathPrefix("/tags").Subrouter()
	tagRoute.Use(m.AuthenticateJWT)

	a.InitializeBlogRouter(blogRoute, *handler)
	a.InitializeCommentsRouter(commentRoute, *handler)
	a.InitializeUserRouter(userRoute, *handler)
	a.InitializetagsRouter(tagRoute, *handler)
	a.InitializeAuthRouter(authRoute, *handler)

	a.Router.HandleFunc("/health", handler.HealthCheckHandler).Methods("GET")
}

func (a *App) Run() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           a.Router,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
