package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

func runConsole() {
	fmt.Println("================================================")
	fmt.Println("   WSL Go Environments Analyzer (консоль)      ")
	fmt.Println("================================================")
	fmt.Println("Привет! Сегодня мы исследуем все доступные дистрибутивы WSL")
	fmt.Println("и выберем лучшие для Go-разработки и ИИ-помощников.")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Дистрибутив (Friendly Name)\tИмя для установки\tПоддержка Go\tИИ-совместимость\tКоманда установки\tEXP")
	fmt.Fprintln(w, "---------------------------\t-----------------\t------------\t-----------------\t-----------------\t---")

	totalExp := 0
	for _, d := range distros {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%d\n",
			d.FriendlyName, d.Name, d.GoSupport, d.AIToolsCompatibility, d.WSLInstallCmd, d.Exp)
		totalExp += d.Exp
	}
	w.Flush()

	fmt.Printf("\n📊 Итого базовый EXP за изучение всех дистрибутивов: %d\n", totalExp)
	fmt.Println("🏆 Ты получаешь ачивку: \"Знаток WSL\" (+200 EXP)")
	totalExp += 200
	fmt.Printf("🎮 Общий EXP за модуль: %d\n", totalExp)

	fmt.Println("\n🔍 Детальная информация (JSON):")
	jsonData, _ := json.MarshalIndent(distros, "", "  ")
	fmt.Println(string(jsonData))

	fmt.Println("\n💡 Совет: Для Go-разработки лучше всего подходят Ubuntu, Debian, Fedora, AlmaLinux.")
	fmt.Println("   Если любишь всё свежее — Arch или Tumbleweed. Для максимальной стабильности — Debian или SLE.")
	fmt.Println("👟 И помни: как правильная обувь защищает от влаги, так и правильный дистрибутив защитит от багов!")
}

func main() {
	webMode := flag.Bool("web", false, "запустить веб-интерфейс")
	flag.Parse()

	if *webMode {
		StartWebServer() // обрати внимание: большая буква S
	} else {
		runConsole()
	}
}
