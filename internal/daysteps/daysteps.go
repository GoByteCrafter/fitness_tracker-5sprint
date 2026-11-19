package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

/*
Структура содержит все необходимые данные о дневных прогулках:
  - количество шагов
  - длительность

а также данные из структуры personaldata.Personal:
  - вес
  - рост
*/
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

/*
Метод парсит строку с данными формата "678,0h50m"
и записывает данные в соответствующие поля структуры DaySteps.
*/
func (ds *DaySteps) Parse(datastring string) (err error) {

	parseData := strings.Split(datastring, ",")
	if len(parseData) != 2 {
		return fmt.Errorf("ошибка: некорректный формат данных %s", datastring)
	}

	steps, err := strconv.Atoi(parseData[0])
	switch {
	case err != nil:
		return err
	case steps <= 0:
		return errors.New("ошибка: количество шагов должно быть больше 0")
	default:
		ds.Steps = steps
	}

	duration, err := time.ParseDuration(parseData[1])
	switch {
	case err != nil:
		return err
	case duration <= 0:
		return errors.New("ошибка: продолжительность должна быть больше 0")
	default:
		ds.Duration = duration
	}

	return nil
}

/*
Метод формирует и возвращает строку с данными о прогулке.
*/
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)
	return result, nil
}
