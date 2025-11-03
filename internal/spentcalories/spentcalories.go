package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

// Функция преобразования строки
func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	divStr := strings.Split(data, ",")
	if len(divStr) != 3 {
		return 0, "", 0, fmt.Errorf("длина строки меньше 3")
	}
	colSteps, err := strconv.Atoi(divStr[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преобразования шагов:%w", err)
	}
	if colSteps <= 0 {
		return 0, "", 0, fmt.Errorf("отрицательное количество шагов")
	}
	viewActiv := divStr[1]
	durationAct, err := time.ParseDuration(divStr[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("ошибка преоьразования продолжительности:%w", err)

	}
	if durationAct <= 0 {
		return 0, "", 0, fmt.Errorf("отрицательный ввод продолжительности")
	}
	return colSteps, viewActiv, durationAct, nil

}

// Функция подсчёта дистанции
func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lengthStep := height * stepLengthCoefficient
	dist := (float64(steps) * lengthStep) / mInKm
	return dist
}

// Функция расчёта средней скорости
func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	Distance := distance(steps, height)
	averageSpeed := Distance / float64(duration.Hours())
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var calories float64
	var dist float64
	var speed float64

	switch activity {
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)

	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, duration)

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		activity, duration.Hours(), dist, speed, calories)

	return info, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("отрицательное количество шагов")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("неверное значение веса")

	}
	if height <= 0 {
		return 0, fmt.Errorf("неверное значение роста")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("неверне значение продолжительности")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	numCaloriesConsumed := (weight * averageSpeed * durationMin) / minInH
	return numCaloriesConsumed, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("отрицательное количество шагов")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("неверное значение веса")

	}
	if height <= 0 {
		return 0, fmt.Errorf("неверное значение роста")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("неверное значение продолжительности")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	durationMin := duration.Minutes()
	numCaloriesConsumed := (weight * averageSpeed * durationMin) / minInH
	res := numCaloriesConsumed * walkingCaloriesCoefficient
	return res, nil

}
