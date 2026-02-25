package main

// Distro представляет информацию о дистрибутиве Linux для WSL
type Distro struct {
	Name                 string   `json:"name"`
	FriendlyName         string   `json:"friendly_name"`
	Description          string   `json:"description"`
	GoSupport            string   `json:"go_support"`
	AIToolsCompatibility string   `json:"ai_tools_compatibility"`
	WSLInstallCmd        string   `json:"wsl_install_cmd"`
	Notes                []string `json:"notes"`
	Exp                  int      `json:"exp"`
}

// Полный список дистрибутивов (из вывода wsl --list --online)
var distros = []Distro{
	{
		Name:                 "Ubuntu",
		FriendlyName:         "Ubuntu",
		Description:          "Базовая версия Ubuntu LTS (обычно последняя стабильная).",
		GoSupport:            "Отлично",
		AIToolsCompatibility: "Отлично",
		WSLInstallCmd:        "wsl --install -d Ubuntu",
		Notes:                []string{"Популярный", "Много документации", "Стабильный"},
		Exp:                  50,
	},
	{
		Name:                 "Ubuntu-24.04",
		FriendlyName:         "Ubuntu 24.04 LTS",
		Description:          "Свежая LTS-версия Ubuntu с длительной поддержкой.",
		GoSupport:            "Отлично",
		AIToolsCompatibility: "Отлично",
		WSLInstallCmd:        "wsl --install -d Ubuntu-24.04",
		Notes:                []string{"Новые пакеты", "Стабильность LTS"},
		Exp:                  50,
	},
	{
		Name:                 "Ubuntu-22.04",
		FriendlyName:         "Ubuntu 22.04 LTS",
		Description:          "Предыдущая LTS-версия, проверенная временем.",
		GoSupport:            "Отлично",
		AIToolsCompatibility: "Отлично",
		WSLInstallCmd:        "wsl --install -d Ubuntu-22.04",
		Notes:                []string{"Надёжная", "Консервативные версии"},
		Exp:                  45,
	},
	{
		Name:                 "Ubuntu-20.04",
		FriendlyName:         "Ubuntu 20.04 LTS",
		Description:          "Более старая LTS, но всё ещё поддерживается.",
		GoSupport:            "Хорошо",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d Ubuntu-20.04",
		Notes:                []string{"Для обратной совместимости"},
		Exp:                  40,
	},
	{
		Name:                 "Debian",
		FriendlyName:         "Debian GNU/Linux",
		Description:          "Эталон стабильности, родитель Ubuntu.",
		GoSupport:            "Отлично",
		AIToolsCompatibility: "Отлично",
		WSLInstallCmd:        "wsl --install -d Debian",
		Notes:                []string{"Максимально стабильный", "Консервативные пакеты"},
		Exp:                  55,
	},
	{
		Name:                 "openSUSE-Tumbleweed",
		FriendlyName:         "openSUSE Tumbleweed",
		Description:          "Rolling-релиз, всегда самые свежие версии ПО.",
		GoSupport:            "Отлично",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d openSUSE-Tumbleweed",
		Notes:                []string{"Свежайший софт", "Нестабильность возможна"},
		Exp:                  60,
	},
	{
		Name:                 "openSUSE-Leap-16.0",
		FriendlyName:         "openSUSE Leap 16.0",
		Description:          "Стабильная версия с фиксированными релизами.",
		GoSupport:            "Отлично",
		AIToolsCompatibility: "Отлично",
		WSLInstallCmd:        "wsl --install -d openSUSE-Leap-16.0",
		Notes:                []string{"Стабильный", "Хорошая интеграция с YaST"},
		Exp:                  55,
	},
	{
		Name:                 "openSUSE-Leap-15.6",
		FriendlyName:         "openSUSE Leap 15.6",
		Description:          "Предыдущая стабильная ветка Leap.",
		GoSupport:            "Хорошо",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d openSUSE-Leap-15.6",
		Notes:                []string{"Проверенная временем"},
		Exp:                  45,
	},
	{
		Name:                 "SUSE-Linux-Enterprise-15-SP7",
		FriendlyName:         "SUSE Linux Enterprise 15 SP7",
		Description:          "Корпоративная версия SUSE, максимальная стабильность.",
		GoSupport:            "Хорошо",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d SUSE-Linux-Enterprise-15-SP7",
		Notes:                []string{"Промышленная надёжность", "Требует подписки?"},
		Exp:                  50,
	},
	{
		Name:                 "SUSE-Linux-Enterprise-16.0",
		FriendlyName:         "SUSE Linux Enterprise 16.0",
		Description:          "Новая мажорная версия SLE.",
		GoSupport:            "Хорошо",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d SUSE-Linux-Enterprise-16.0",
		Notes:                []string{"Современное ядро", "Корпоративная поддержка"},
		Exp:                  50,
	},
	{
		Name:                 "SUSE-Linux-Enterprise-15-SP6",
		FriendlyName:         "SUSE Linux Enterprise 15 SP6",
		Description:          "Предыдущий Service Pack SLE15.",
		GoSupport:            "Хорошо",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d SUSE-Linux-Enterprise-15-SP6",
		Notes:                []string{"Стабильный", "Проверенный"},
		Exp:                  45,
	},
	{
		Name:                 "kali-linux",
		FriendlyName:         "Kali Linux Rolling",
		Description:          "Дистрибутив для пентеста, множество предустановленных инструментов.",
		GoSupport:            "Хорошо",
		AIToolsCompatibility: "Средне",
		WSLInstallCmd:        "wsl --install -d kali-linux",
		Notes:                []string{"Для безопасности", "Избыточен для обычной разработки"},
		Exp:                  60,
	},
	{
		Name:                 "AlmaLinux-9",
		FriendlyName:         "AlmaLinux OS 9",
		Description:          "RHEL-совместимый дистрибутив, стабильный и бесплатный.",
		GoSupport:            "Отлично",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d AlmaLinux-9",
		Notes:                []string{"Преемник CentOS", "Стабильный"},
		Exp:                  55,
	},
	{
		Name:                 "AlmaLinux-8",
		FriendlyName:         "AlmaLinux OS 8",
		Description:          "Предыдущая версия AlmaLinux.",
		GoSupport:            "Хорошо",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d AlmaLinux-8",
		Notes:                []string{"Старый, но надёжный"},
		Exp:                  45,
	},
	{
		Name:                 "AlmaLinux-Kitten-10",
		FriendlyName:         "AlmaLinux OS Kitten 10",
		Description:          "Тестовая версия (разработка).",
		GoSupport:            "Средне",
		AIToolsCompatibility: "Средне",
		WSLInstallCmd:        "wsl --install -d AlmaLinux-Kitten-10",
		Notes:                []string{"Не для продакшена"},
		Exp:                  30,
	},
	{
		Name:                 "AlmaLinux-10",
		FriendlyName:         "AlmaLinux OS 10",
		Description:          "Новая мажорная версия AlmaLinux.",
		GoSupport:            "Отлично",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d AlmaLinux-10",
		Notes:                []string{"Современный", "Активно развивается"},
		Exp:                  55,
	},
	{
		Name:                 "archlinux",
		FriendlyName:         "Arch Linux",
		Description:          "Rolling-релиз, всё самое свежее, требует ручной настройки.",
		GoSupport:            "Отлично",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d archlinux",
		Notes:                []string{"Гибкий", "Сложный для новичков"},
		Exp:                  65,
	},
	{
		Name:                 "FedoraLinux-43",
		FriendlyName:         "Fedora Linux 43",
		Description:          "Последняя версия Fedora, инновации, свежие пакеты.",
		GoSupport:            "Отлично",
		AIToolsCompatibility: "Отлично",
		WSLInstallCmd:        "wsl --install -d FedoraLinux-43",
		Notes:                []string{"Современный", "Быстро обновляется"},
		Exp:                  55,
	},
	{
		Name:                 "FedoraLinux-42",
		FriendlyName:         "Fedora Linux 42",
		Description:          "Предыдущая версия Fedora.",
		GoSupport:            "Хорошо",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d FedoraLinux-42",
		Notes:                []string{"Ещё поддерживается"},
		Exp:                  45,
	},
	{
		Name:                 "eLxr",
		FriendlyName:         "eLxr 12.12.0.0 GNU/Linux",
		Description:          "Редкий дистрибутив, возможно специализированный.",
		GoSupport:            "Средне",
		AIToolsCompatibility: "Средне",
		WSLInstallCmd:        "wsl --install -d eLxr",
		Notes:                []string{"Экзотика", "Мало информации"},
		Exp:                  30,
	},
	{
		Name:                 "OracleLinux_9_5",
		FriendlyName:         "Oracle Linux 9.5",
		Description:          "RHEL-совместимый дистрибутив от Oracle.",
		GoSupport:            "Хорошо",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d OracleLinux_9_5",
		Notes:                []string{"Корпоративный", "Интеграция с Oracle"},
		Exp:                  50,
	},
	{
		Name:                 "OracleLinux_8_10",
		FriendlyName:         "Oracle Linux 8.10",
		Description:          "Предыдущая версия Oracle Linux.",
		GoSupport:            "Хорошо",
		AIToolsCompatibility: "Хорошо",
		WSLInstallCmd:        "wsl --install -d OracleLinux_8_10",
		Notes:                []string{"Стабильный"},
		Exp:                  45,
	},
	{
		Name:                 "OracleLinux_7_9",
		FriendlyName:         "Oracle Linux 7.9",
		Description:          "Старая версия Oracle Linux.",
		GoSupport:            "Средне",
		AIToolsCompatibility: "Средне",
		WSLInstallCmd:        "wsl --install -d OracleLinux_7_9",
		Notes:                []string{"Устаревшая"},
		Exp:                  35,
	},
}

// scoreToInt преобразует строковую оценку в число
func scoreToInt(score string) int {
	switch score {
	case "Отлично":
		return 2
	case "Хорошо":
		return 1
	default:
		return 0
	}
}

// Filters содержит параметры фильтрации
type Filters struct {
	MinGoScore int
	MinAIScore int
	MinExp     int
	SortBy     string // "name", "exp", "go", "ai"
	SortOrder  string // "asc", "desc"
}

// FilterDistros применяет фильтры и сортировку
func FilterDistros(dists []Distro, f Filters) []Distro {
	var result []Distro
	for _, d := range dists {
		goScore := scoreToInt(d.GoSupport)
		aiScore := scoreToInt(d.AIToolsCompatibility)
		if goScore >= f.MinGoScore && aiScore >= f.MinAIScore && d.Exp >= f.MinExp {
			result = append(result, d)
		}
	}

	// Простая сортировка пузырьком (для наглядности, можно заменить на sort.Slice)
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			less := false
			switch f.SortBy {
			case "name":
				if f.SortOrder == "asc" {
					less = result[i].FriendlyName > result[j].FriendlyName
				} else {
					less = result[i].FriendlyName < result[j].FriendlyName
				}
			case "exp":
				if f.SortOrder == "asc" {
					less = result[i].Exp > result[j].Exp
				} else {
					less = result[i].Exp < result[j].Exp
				}
			case "go":
				if f.SortOrder == "asc" {
					less = scoreToInt(result[i].GoSupport) > scoreToInt(result[j].GoSupport)
				} else {
					less = scoreToInt(result[i].GoSupport) < scoreToInt(result[j].GoSupport)
				}
			case "ai":
				if f.SortOrder == "asc" {
					less = scoreToInt(result[i].AIToolsCompatibility) > scoreToInt(result[j].AIToolsCompatibility)
				} else {
					less = scoreToInt(result[i].AIToolsCompatibility) < scoreToInt(result[j].AIToolsCompatibility)
				}
			default:
				if f.SortOrder == "asc" {
					less = result[i].FriendlyName > result[j].FriendlyName
				} else {
					less = result[i].FriendlyName < result[j].FriendlyName
				}
			}
			if less {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}

// SumExp возвращает сумму EXP для списка дистрибутивов
func SumExp(dists []Distro) int {
	total := 0
	for _, d := range dists {
		total += d.Exp
	}
	return total
}
