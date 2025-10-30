package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

// Функция преобразования строки
func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	divStr := strings.Split(data, ",")
	if len(divStr) != 2 {
		return 0, 0, fmt.Errorf("длина слайса меньше 2")
	}
	colSteps, err := strconv.Atoi(divStr[0])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования шагов:%w", err)
	}
	if colSteps < 0 {
		return 0, 0, fmt.Errorf("отрицательное количество шагов")
	}
	duration, err := time.ParseDuration(divStr[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка преобразования продолжительности:%w", err)
	}
	return colSteps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	colSteps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if colSteps < 0 {
		return ""
	}
	distance := float64(colSteps) * stepLength
	DistKm := distance / float64(mInKm)
	colCalories, err := spentcalories.WalkingSpentCalories(colSteps, weight, height, duration)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		colSteps, DistKm, colCalories)
}
