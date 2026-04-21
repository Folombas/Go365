package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UTC()

	// Фиксированные даты
	go111Day := time.Date(2026, time.April, 21, 0, 0, 0, 0, time.UTC)
	anniversary := time.Date(2026, time.April, 22, 0, 0, 0, 0, time.UTC)
	firstHelloDate := time.Date(2025, time.April, 22, 0, 0, 0, 0, time.UTC)

	fmt.Println("🐹 День Go111 — 21 апреля 2026 года 🐹")
	fmt.Println("=====================================")

	// Сегодня 21 апреля 2026?
	if now.Year() == go111Day.Year() && now.YearDay() == go111Day.YearDay() {
		fmt.Println("✅ СЕГОДНЯ официальный День Go111!")
		fmt.Println("🎉 Поздравляем всех gophers, включая вас!")

		fmt.Println("📅 А ЗАВТРА (22 апреля 2026) — важная дата:")
		fmt.Println("   Ровно ГОД, как вы написали свою первую программу на Go:")
		fmt.Println("   package main\n   import \"fmt\"\n   func main() { fmt.Println(\"Hello, World!\") }")
		fmt.Printf("   Это было %s.\n\n", firstHelloDate.Format("2 January 2006"))

		fmt.Println("❌ Но работа программистом Go до сих пор не найдена.")
		fmt.Println("🤔 Почему? Возможные причины:")
		fmt.Println("   - Теории много, реальных проектов мало")
		fmt.Println("   - Нет портфолио на GitHub")
		fmt.Println("   - Боитесь откликаться на вакансии")
		fmt.Println("   - Не показываете свои pet-проекты в резюме")

	} else if now.Year() == anniversary.Year() && now.YearDay() == anniversary.YearDay() {
		// Это 22 апреля 2026 (завтрашний день, но если вдруг запустить завтра)
		fmt.Println("📅 СЕГОДНЯ ровно год вашему первому 'Hello, World!' на Go!")
		fmt.Println("❌ Но вы всё ещё без работы.")
		// далее такой же совет
	} else {
		fmt.Println("⏳ Эта программа точно для 21 апреля 2026. Запустите её в этот день!")
		return
	}

	// Главный вопрос
	fmt.Println("\n💡 Стоит ли дальше обучаться Go?")
	fmt.Println("   ✅ Да, обязательно! Go — отличный язык с хорошим рынком.")
	fmt.Println("   ❗ Но учиться дальше нужно уже через практику, а не просто читать.")

	fmt.Println("\n📌 Что делать прямо сейчас (не откладывая на завтра):")
	fmt.Println("   1️⃣ За сегодня (21 апреля) соберите портфолио из 2-3 готовых проектов на Go.")
	fmt.Println("   2️⃣ Завтра, в годовщину, отправьте резюме минимум в 5 компаний.")
	fmt.Println("   3️⃣ Напишите в своём айти-блоге пост о годе с Go — это привлекает рекрутеров.")
	fmt.Println("   4️⃣ Найдите Go-сообщество в Сети и попросите код-ревью.")
	fmt.Println("   5️⃣ Перестаньте учить 'ещё один фреймворк' — начните решать реальные задачи (LeetCode, codewars, microservices).")

	fmt.Println("\n🎯 Запомните: год обучения — норма. Год без практики и поиска работы — потеря времени.")
	fmt.Println("🚀 У вас всё получится! Завтра начните действовать.")
}