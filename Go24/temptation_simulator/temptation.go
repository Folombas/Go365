package main

import (
	"math/rand"
)

type Temptation struct {
	Name        string
	Power       int    // Сила искушения (0-100)
	Description string
}

type Motivation struct {
	Text   string
	Effect string // "focus+", "willpower+", "knowledge+"
}

var temptations = []Temptation{
	{"CapCut Видеомонтаж", 85, "Установить программу и монтировать видео отпуска"},
	{"Видеоигры", 70, "Поиграть в новую RPG-игру"},
	{"Соцсети", 60, "Прокрутить ленту Instagram 2 часа"},
	{"Бары/Клубы", 75, "Сходить в бар с друзьями"},
	{"Фильмы/Сериалы", 65, "Посмотреть новый сезон сериала"},
	{"Покупки онлайн", 55, "Купить ненужные вещи на маркетплейсе"},
	{"Еда", 50, "Съесть пиццу вместо ужина"},
	{"Прокрастинация", 80, "Отложить изучение Go на завтра"},
}

var motivations = []Motivation{
	{"Каждая строка кода на Go — кирпичик в фундаменте карьеры", "knowledge+"},
	{"Сегодняшний дискомфорт — завтрашний комфорт зарплаты", "focus+"},
	{"Распыляться — значит стоять на месте. Фокус — значит расти", "willpower+"},
	{"Хобби подождут, когда будет стабильный доход", "focus+"},
	{"Гофер — символ эффективности и простоты", "knowledge+"},
	{"Каждый коммит приближает к офису с видом на город", "knowledge+"},
	{"Ошибки — не провалы, а инструкции к улучшению", "focus+"},
	{"Go учит не только программировать, но и мыслить системно", "knowledge+"},
	{"Сила воли в программировании важнее, чем в спортзале", "willpower+"},
	{"Экосистема Go — твой новый город возможностей", "knowledge+"},
}

func GenerateTemptation() Temptation {
	index := rand.Intn(len(temptations))
	return temptations[index]
}

func GetRandomMotivation() Motivation {
	index := rand.Intn(len(motivations))
	return motivations[index]
}
