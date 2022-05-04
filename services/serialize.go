package services

type Response struct {
	Status  bool        `json:"status"`
	Content interface{} `json:"content"`
}
