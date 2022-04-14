package types

type TaskResult struct {
	Payload [][]interface{} `json:"payload"`
	Results []interface{}   `json:"results"`
}
