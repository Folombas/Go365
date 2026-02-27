package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Post — структура записи в блоге
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Exp       int       `json:"exp"`
	Excerpt   string    `json:"-"` // не сохраняется, используется в шаблонах
}

// User — данные пользователя (Гоши) для геймификации
type User struct {
	Level        int      `json:"level"`
	Exp          int      `json:"exp"`
	NextLevelExp int      `json:"next_level_exp"`
	Achievements []string `json:"achievements"`
}

var (
	posts     []Post
	nextID    = 1
	user      User
	postsFile = "posts.json"
	userFile  = "user.json"
)

const expPerPost = 10

// truncate обрезает текст до указанной длины, добавляя многоточие
func truncate(s string, max int) string {
	runes := []rune(s)
	if len(runes) <= max {
		return s
	}
	return string(runes[:max]) + "..."
}

// Загрузка данных из файлов
func loadData() {
	// Загрузка постов
	if data, err := ioutil.ReadFile(postsFile); err == nil {
		json.Unmarshal(data, &posts)
		if len(posts) > 0 {
			nextID = posts[len(posts)-1].ID + 1
		}
	}
	// Заполняем Excerpt для всех постов
	for i := range posts {
		posts[i].Excerpt = truncate(posts[i].Content, 200)
	}

	// Загрузка пользователя
	if data, err := ioutil.ReadFile(userFile); err == nil {
		json.Unmarshal(data, &user)
	} else {
		// Если нет, создаём нового с нуля
		user = User{
			Level:        1,
			Exp:          0,
			NextLevelExp: 100,
			Achievements: []string{},
		}
	}
}

// Сохранение данных
func savePosts() {
	// Перед сохранением убираем поле Excerpt (чтобы не сохранять)
	for i := range posts {
		posts[i].Excerpt = ""
	}
	data, _ := json.MarshalIndent(posts, "", "  ")
	ioutil.WriteFile(postsFile, data, 0644)
	// Восстанавливаем Excerpt после сохранения
	for i := range posts {
		posts[i].Excerpt = truncate(posts[i].Content, 200)
	}
}

func saveUser() {
	data, _ := json.MarshalIndent(user, "", "  ")
	ioutil.WriteFile(userFile, data, 0644)
}

// Добавление опыта, проверка уровня
func addExp(amount int) {
	user.Exp += amount
	for user.Exp >= user.NextLevelExp {
		user.Level++
		user.Exp -= user.NextLevelExp
		user.NextLevelExp = int(math.Floor(float64(user.NextLevelExp) * 1.5))
		user.Achievements = append(user.Achievements, fmt.Sprintf("Достигнут уровень %d! (+%d EXP)", user.Level, amount))
	}
	checkAchievements()
	saveUser()
}

// Проверка особых достижений
func checkAchievements() {
	if len(posts) >= 10 && !contains(user.Achievements, "10 постов (+50 EXP)") {
		user.Achievements = append(user.Achievements, "10 постов (+50 EXP)")
		addExp(50)
	}
	if len(posts) >= 50 && !contains(user.Achievements, "50 постов (+200 EXP)") {
		user.Achievements = append(user.Achievements, "50 постов (+200 EXP)")
		addExp(200)
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Обработчик главной страницы
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Получаем последние 5 постов
	recentPosts := posts
	if len(recentPosts) > 5 {
		recentPosts = posts[len(posts)-5:]
	}

	// Вычисляем процент прогресса
	progress := float64(user.Exp) / float64(user.NextLevelExp) * 100
	progressWidth := fmt.Sprintf("%.0f", progress)

	data := struct {
		Posts         []Post
		User          User
		ProgressWidth string
	}{
		Posts:         recentPosts,
		User:          user,
		ProgressWidth: progressWidth,
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, data)
}

// Обработчик списка всех постов
func postsHandler(w http.ResponseWriter, r *http.Request) {
	// Сортируем по дате (новые сверху)
	sorted := make([]Post, len(posts))
	copy(sorted, posts)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].CreatedAt.After(sorted[j].CreatedAt)
	})
	// Обновляем Excerpt для сортированных копий
	for i := range sorted {
		sorted[i].Excerpt = truncate(sorted[i].Content, 200)
	}

	data := struct {
		Posts []Post
		User  User
	}{
		Posts: sorted,
		User:  user,
	}

	tmpl := template.Must(template.ParseFiles("templates/posts.html"))
	tmpl.Execute(w, data)
}

// Обработчик отдельного поста
func postHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		http.NotFound(w, r)
		return
	}
	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var post Post
	for _, p := range posts {
		if p.ID == id {
			post = p
			break
		}
	}
	if post.ID == 0 {
		http.NotFound(w, r)
		return
	}

	data := struct {
		Post Post
		User User
	}{
		Post: post,
		User: user,
	}

	tmpl := template.Must(template.ParseFiles("templates/post.html"))
	tmpl.Execute(w, data)
}

// Обработчик добавления нового поста (POST)
func newPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")
	if title == "" || content == "" {
		http.Error(w, "Title and content required", http.StatusBadRequest)
		return
	}

	post := Post{
		ID:        nextID,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		Exp:       expPerPost,
		Excerpt:   truncate(content, 200),
	}
	nextID++
	posts = append(posts, post)
	addExp(expPerPost)
	savePosts()

	http.Redirect(w, r, fmt.Sprintf("/post/%d", post.ID), http.StatusSeeOther)
}

func main() {
	loadData()

	// Статические файлы (CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Маршруты
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/post/", postHandler)
	http.HandleFunc("/new", newPostHandler)

	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
