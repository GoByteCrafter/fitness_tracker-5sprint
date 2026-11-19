package personaldata

import "fmt"

// Структура для хранения данных пользователей
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Метод для вывода на экран данных пользователей
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
