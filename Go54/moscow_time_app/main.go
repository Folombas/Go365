package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/api/time", handleTime)
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Московское время</title>
    <script>
        async function updateTime() {
            const response = await fetch('/api/time');
            const data = await response.json();
            document.getElementById('time').innerText = data.time;
        }
        setInterval(updateTime, 1000);
        window.onload = updateTime;
    </script>
</head>
<body>
    <h1>Текущее московское время:</h1>
    <div id="time" style="font-size: 48px; font-family: monospace;"></div>
</body>
</html>`
	t, _ := template.New("index").Parse(tmpl)
	t.Execute(w, nil)
}

func handleTime(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("Europe/Moscow")
	now := time.Now().In(loc).Format("15:04:05")
	json.NewEncoder(w).Encode(map[string]string{"time": now})
}
