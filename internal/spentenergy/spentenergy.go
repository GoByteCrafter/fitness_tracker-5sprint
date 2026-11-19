package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

/*
Функция принимает:

	steps int — количество шагов.
	weight, height float64 — вес(кг.) и рост(м.) пользователя.
	duration time.Duration — продолжительность ходьбы.

И возвращает два значения:

	float64 — количество калорий, потраченных при ходьбе.
	error — ошибку, если входные параметры некорректны (подумайте, какие значения параметров имеют смысл).
*/
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры")
	}
	result, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return result * walkingCaloriesCoefficient, nil
}

/*
Функция принимает:

	steps int — количество шагов.
	weight, height float64 — вес(кг.) и рост(м.) пользователя.
	duration time.Duration — продолжительность бега.

И возвращает два значения:

	float64 — количество калорий, потраченных при беге.
	error — ошибку, если входные параметры некорректны
*/
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры")
	}

	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

// Функция принимает количество шагов steps, рост пользователя height и продолжительность активности duration и возвращает среднюю скорость.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

// Функция принимает количество шагов и рост пользователя в метрах, а возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := stepLengthCoefficient * height
	return (float64(steps) * stepLength) / mInKm
}
