package models

type Post struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Body        string `json:"body"`
	DateCreated string `json:"date_created"`
	ImageUrl    string `json:"image_url"`
}
