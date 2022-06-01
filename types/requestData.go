package types

// Тип для отправки решения задачи на проверку
// UserName - Никнейм
// TaskName - Название задачи
// Result - Результат задачи
type RequestData struct {
	UserName string     `json:"user_name"`
	TaskName string     `json:"task"`
	Result   TaskResult `json:"results"`
}
