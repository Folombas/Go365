package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	//"os"
	"sort"
	"time"
)

type MoodEntry struct {
	Mood      string    `json:"mood"`
	Note      string    `json:"note"`
	Timestamp time.Time `json:"timestamp"`
	Exp       int       `json:"exp"`
}

var entries []MoodEntry
var expTotal int
var achievements []string

const expPerEntry = 10
const dataFile = "data.json"

func loadEntries() {
	file, err := ioutil.ReadFile(dataFile)
	if err == nil {
		json.Unmarshal(file, &entries)
		expTotal = 0
		for _, e := range entries {
			expTotal += e.Exp
		}
	}
}

func saveEntries() {
	data, _ := json.MarshalIndent(entries, "", "  ")
	ioutil.WriteFile(dataFile, data, 0644)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		mood := r.FormValue("mood")
		note := r.FormValue("note")
		entry := MoodEntry{
			Mood:      mood,
			Note:      note,
			Timestamp: time.Now(),
			Exp:       expPerEntry,
		}
		entries = append(entries, entry)
		expTotal += expPerEntry
		checkAchievements()
		saveEntries()
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	moodCount := make(map[string]int)
	for _, e := range entries {
		moodCount[e.Mood]++
	}

	type moodStat struct {
		Mood  string
		Count int
	}
	var stats []moodStat
	for m, c := range moodCount {
		stats = append(stats, moodStat{m, c})
	}
	sort.Slice(stats, func(i, j int) bool { return stats[i].Count > stats[j].Count })

	var recent []MoodEntry
	if len(entries) > 10 {
		recent = entries[len(entries)-10:]
	} else {
		recent = entries
	}

	data := struct {
		ExpTotal     int
		Achievements []string
		Stats        []moodStat
		Recent       []MoodEntry
	}{
		ExpTotal:     expTotal,
		Achievements: achievements,
		Stats:        stats,
		Recent:       recent,
	}

	tmpl := template.Must(template.ParseFiles("templates/stats.html"))
	tmpl.Execute(w, data)
}

func checkAchievements() {
	if len(entries) >= 10 && !contains(achievements, "10 записей (+50 EXP)") {
		achievements = append(achievements, "10 записей (+50 EXP)")
		expTotal += 50
	}
	if len(entries) >= 50 && !contains(achievements, "50 записей (+100 EXP)") {
		achievements = append(achievements, "50 записей (+100 EXP)")
		expTotal += 100
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

func main() {
	loadEntries()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/stats", statsHandler)
	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
