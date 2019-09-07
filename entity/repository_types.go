package entity

type RepositorySummary struct {
	Id          int64  `json:"id"`
	FullName    string `json:"fullName"`
	Star        int    `json:"star"`
	Owner       string `json:"ownerName"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updatedAt"`
	Language    string `json:"language"`
	Size        int    `json:"size"`
	License     string `json:"license"`
}
