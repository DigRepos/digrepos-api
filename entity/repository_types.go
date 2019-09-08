package entity

type Owner struct {
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
}

type RepositorySummary struct {
	Id            int64    `json:"id"`
	HTMLURL       string   `json:"url"`
	FullName      string   `json:"fullName"`
	Star          int      `json:"star"`
	ForksCount    int      `json:"forksCount"`
	WatchersCount int      `json:"watchersCount"`
	Owner         Owner    `json:"owner"`
	Description   string   `json:"description"`
	Homepage      string   `json:"homepage"`
	Topics        []string `json:"topics"`
	UpdatedAt     string   `json:"updatedAt"`
	Language      string   `json:"language"`
	Size          int      `json:"size"`
	License       string   `json:"license"`
}
