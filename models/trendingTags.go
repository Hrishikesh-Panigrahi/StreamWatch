package models

type TrendingTags struct {
	Id          int    `json:"id"`
	Tag         string `json:"tag"`
	Usage_count int    `json:"usage_count"`
}
