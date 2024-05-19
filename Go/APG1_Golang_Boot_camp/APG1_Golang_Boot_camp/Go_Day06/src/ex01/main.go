package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/russross/blackfriday/v2"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "main/docs" // Swagger docs
)

// @title My Blog API
// @version 1.0
// @description This is a simple blog API with admin panel, markdown support, and pagination.
// @host localhost:8888
// @BasePath /

var (
	templates     *template.Template
	db            *sql.DB
	adminUsername string
	adminPassword string
)

type Article struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
}

func init() {
	funcMap := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
		"add": func(a, b int) int {
			return a + b
		},
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	templates = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))

	file, err := os.Open("admin_credentials.txt")
	if err != nil {
		log.Fatal("Error opening admin credentials file:", err)
	}
	defer file.Close()

	var dbUser, dbPassword, dbName, db_scripts string
	_, err = fmt.Fscanf(file, "admin_username: %s\nadmin_password: %s\ndb_user: %s\ndb_password: %s\ndb_name: %s\n",
		&adminUsername, &adminPassword, &dbUser, &dbPassword, &dbName)
	if err != nil {
		log.Fatal("Error reading admin credentials:", err)
	}

	psqlInfo := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	f, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading admin credentials:", err)
	}
	db_scripts = string(f)

	db_scripts = strings.SplitAfter(db_scripts, "-- SQL scripts")[1]

	// Create tables if they don't exist
	_, err = db.Exec(db_scripts)
	if err != nil {
		log.Fatal("Error creating articles table:", err)
	}

}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/post/:id", ViewPost)
	router.GET("/admin", Admin)
	router.POST("/admin", AdminPost)
	router.GET("/login", Login)
	router.POST("/login", LoginPost)
	router.Handler("GET", "/swagger/*any", httpSwagger.WrapHandler)
	router.ServeFiles("/css/*filepath", http.Dir("css"))
	router.ServeFiles("/images/*filepath", http.Dir("images"))
	router.ServeFiles("/js/*filepath", http.Dir("js"))

	log.Println("Server started on port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}

// Index godoc
// @Summary Show the homepage
// @Description Get the homepage with a list of articles
// @Tags public
// @Accept  json
// @Produce  html
// @Param page query int false "Page number"
// @Success 200 {string} string "Homepage HTML"
// @Router / [get]
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	page := r.URL.Query().Get("page")
	pageNum, err := strconv.Atoi(page)
	if err != nil || pageNum < 1 {
		pageNum = 1
	}
	offset := (pageNum - 1) * 3

	rows, err := db.Query("SELECT id, title, content, created_at FROM articles ORDER BY created_at DESC LIMIT 3 OFFSET $1", offset)
	if err != nil {
		fmt.Println("AAA")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var articles []Article

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Title, &article.Content, &article.CreatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if len(article.Content) > 200 {
			article.Content = article.Content[:200] + "..."
		}
		articles = append(articles, article)
	}

	templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
		"Articles": articles,
		"Page":     pageNum,
	})
}

// ViewPost godoc
// @Summary Show a single post
// @Description Get a single post by ID
// @Tags public
// @Accept  json
// @Produce  html
// @Param id path int true "Post ID"
// @Success 200 {string} string "Post HTML"
// @Failure 404 {string} string "Not Found"
// @Router /post/{id} [get]
func ViewPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	var article struct {
		Title   string
		Content string
	}

	err := db.QueryRow("SELECT title, content FROM articles WHERE id = $1", id).Scan(&article.Title, &article.Content)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	article.Content = string(blackfriday.Run([]byte(article.Content)))

	templates.ExecuteTemplate(w, "post.html", article)
}

// Admin godoc
// @Summary Show the admin panel
// @Description Get the admin panel for posting new articles
// @Tags admin
// @Accept  json
// @Produce  html
// @Success 200 {string} string "Admin Panel HTML"
// @Router /admin [get]
func Admin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	templates.ExecuteTemplate(w, "admin.html", nil)
}

// AdminPost godoc
// @Summary Post a new article
// @Description Post a new article from the admin panel
// @Tags admin
// @Accept  application/x-www-form-urlencoded
// @Produce  html
// @Param title formData string true "Article Title"
// @Param content formData string true "Article Content"
// @Success 302 {string} string "Redirect to homepage"
// @Router /admin [post]
func AdminPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")

	_, err := db.Exec("INSERT INTO articles (title, content) VALUES ($1, $2)", title, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// Login godoc
// @Summary Show the login page
// @Description Get the login page
// @Tags admin
// @Accept  json
// @Produce  html
// @Success 200 {string} string "Login Page HTML"
// @Router /login [get]
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

// LoginPost godoc
// @Summary Login to admin panel
// @Description Login to the admin panel
// @Tags admin
// @Accept  json
// @Produce  html
// @Param username formData string true "Admin Username"
// @Param password formData string true "Admin Password"
// @Success 302 {string} string "Redirect to admin panel"
// @Router /login [post]
func LoginPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == adminUsername && password == adminPassword {
		http.Redirect(w, r, "/admin", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
