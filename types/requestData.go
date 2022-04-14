package types

type RequestData struct {
	UserName string     `json:"user_name"`
	TaskName string     `json:"task"`
	Result   TaskResult `json:"results"`
}
