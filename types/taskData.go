package types

// Тип для хранения исходных данных задачи
// Data - Исходные данные
// Offset - Смещение для задачи "Циклическая ротация"
type TaskData struct {
	Data   []int
	Offset int
}
