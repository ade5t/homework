package types

// Тип для хранения результатов задачи
// Payload - Исходные данные
// Results - Результат решения
type TaskResult struct {
	Payload [][]interface{} `json:"payload"`
	Results []interface{}   `json:"results"`
}
