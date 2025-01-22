package models

type ListTask struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Created_at  string `json:"createdAt"`
	Updated_at  string `json:"updatedAt"`
}
