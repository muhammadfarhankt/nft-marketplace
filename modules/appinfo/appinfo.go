package appinfo

type Category struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}

type CategoryFilter struct {
	Title string `query:"title"`
}
