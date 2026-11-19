package trainings

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

/*
Структура содержит все необходимые данные о тренировке:
  - количество шагов,
  - тип тренировки,
  - длительность тренировки,

а также данные из структуры personaldata.Personal, то есть:
  - имя,
  - вес
  - рост пользователя.
*/
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

/*
Метод парсит строку с данными формата "3456,Ходьба,3h00m"
и записывает данные в соответствующие поля структуры Training.
*/
func (t *Training) Parse(datastring string) (err error) {

	parseData := strings.Split(datastring, ",")
	if len(parseData) != 3 {
		return errors.New("неверный формат данных")
	}

	steps, err := strconv.Atoi(parseData[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}
	t.Steps = steps

	t.TrainingType = parseData[1]

	duration, err := time.ParseDuration(parseData[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше 0")
	}
	t.Duration = duration
	return nil
}

/*
Метод формирует и возвращает строку с данными о тренировке,
исходя из того, какой тип тренировки был передан.
*/
func (t Training) ActionInfo() (string, error) {

	var (
		calories float64
		err      error
	)
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			log.Println(err)
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, calories)
	return info, nil

}
