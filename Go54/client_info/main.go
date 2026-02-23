package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/info", infoHandler)
	http.ListenAndServe(":8081", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Информация о пользователе</title>
    <script>
        async function loadInfo() {
            const response = await fetch('/api/info');
            const data = await response.json();
            document.getElementById('os').innerText = data.os;
            document.getElementById('browser').innerText = data.browser;
            document.getElementById('time').innerText = data.time;
        }
        window.onload = loadInfo;
    </script>
</head>
<body>
    <h1>Информация о вашем устройстве</h1>
    <p><strong>Операционная система:</strong> <span id="os">загрузка...</span></p>
    <p><strong>Браузер:</strong> <span id="browser">загрузка...</span></p>
    <p><strong>Время входа:</strong> <span id="time">загрузка...</span></p>
</body>
</html>`
	t, _ := template.New("index").Parse(tmpl)
	t.Execute(w, nil)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.UserAgent()
	os := detectOS(userAgent)
	browser := detectBrowser(userAgent)
	now := time.Now().Format("2006-01-02 15:04:05")
	info := map[string]string{
		"os":      os,
		"browser": browser,
		"time":    now,
	}
	json.NewEncoder(w).Encode(info)
}

func detectOS(ua string) string {
	ua = strings.ToLower(ua)
	switch {
	case strings.Contains(ua, "windows nt"):
		return "Windows"
	case strings.Contains(ua, "mac os"):
		return "macOS"
	case strings.Contains(ua, "linux"):
		return "Linux"
	case strings.Contains(ua, "android"):
		return "Android"
	case strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad"):
		return "iOS"
	default:
		return "Неизвестная ОС"
	}
}

func detectBrowser(ua string) string {
	ua = strings.ToLower(ua)
	switch {
	case strings.Contains(ua, "edg/"):
		return "Microsoft Edge"
	case strings.Contains(ua, "opr/") || strings.Contains(ua, "opera"):
		return "Opera"
	case strings.Contains(ua, "chrome"):
		return "Google Chrome"
	case strings.Contains(ua, "safari") && !strings.Contains(ua, "chrome"):
		return "Safari"
	case strings.Contains(ua, "firefox"):
		return "Mozilla Firefox"
	case strings.Contains(ua, "msie") || strings.Contains(ua, "trident"):
		return "Internet Explorer"
	default:
		return "Неизвестный браузер"
	}
}
