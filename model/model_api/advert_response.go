package model_api

type AdvertResponse struct {
	Id          int64                  `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Category    AdvertCategoryResponse `json:"category"`
}

type AdvertCategoryResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
