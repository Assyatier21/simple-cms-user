package models

type Article struct {
	Id          int    `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Slug        string `json:"slug" form:"slug"`
	HtmlContent string `json:"html_content" form:"html_content"`
	CategoryID  int    `json:"category_id" form:"category_id"`
	CreatedAt   string `json:"created_at" form:"created_at"`
	UpdatedAt   string `json:"updated_at" form:"updated_at"`
}
